package keeper

import (
	"context"
	"errors"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	networktypes "github.com/ignite/network/pkg/types"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/reward/types"
)

// DistributeRewards distributes rewards based on the monitoring packet
// this `closeRewardPool` is a boolean that specifies if the reward pool
// must be closed after the reward distribution.
// In a model where rewards are distributed in a single round, this
// boolean is always `true`. In this way, if the monitoring packet is
// distributed earlier and `lastBlockHeight < rewardPool.LastRewardHeight`
// only a portion of the rewards will be distributed and the remaining is
// refunded to the reward's provider.
// When rewards are distributed periodically, this value is set to `false`
// so the reward pool is not closed as long as `lastBlockHeight` does not
// reach `rewardPool.LastRewardHeight`
func (k Keeper) DistributeRewards(
	ctx context.Context,
	launchID uint64,
	signatureCounts networktypes.SignatureCounts,
	lastBlockHeight int64,
	closeRewardPool bool,
) error {
	// get the reward pool related to the chain
	rewardPool, err := k.RewardPool.Get(ctx, launchID)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrRewardPoolNotFound, "%d", launchID)
	}

	if rewardPool.Closed {
		return sdkerrors.Wrapf(types.ErrRewardPoolClosed, "%d", launchID)
	}

	provider, err := k.addressCodec.StringToBytes(rewardPool.Provider)
	if err != nil {
		return ignterrors.Criticalf("can't parse the provider address %s", err.Error())
	}

	// lastBlockHeight must be strictly greater than the current reward height for the pool
	if lastBlockHeight <= rewardPool.CurrentRewardHeight {
		return sdkerrors.Wrapf(
			types.ErrInvalidLastBlockHeight,
			"last block height %d must be greater than current reward height for the reward pool %d",
			lastBlockHeight,
			rewardPool.CurrentRewardHeight,
		)
	}

	// only the monitored blocks relative to last reward height are rewarded
	blockRatioNumerator := sdkmath.LegacyNewDec(lastBlockHeight).Sub(sdkmath.LegacyNewDec(rewardPool.CurrentRewardHeight))
	blockRatioDenominator := sdkmath.LegacyNewDec(rewardPool.LastRewardHeight).Sub(sdkmath.LegacyNewDec(rewardPool.CurrentRewardHeight))
	blockRatio := blockRatioNumerator.Quo(blockRatioDenominator)
	if blockRatio.GT(sdkmath.LegacyOneDec()) {
		blockRatio = sdkmath.LegacyOneDec()
	}

	// store the total relative signature distributed to calculate the refund for the round
	totalRelativeSignaturesDistributed := sdkmath.LegacyZeroDec()

	// store rewards to distributes per address
	rewardsToDistribute := make(map[string]sdk.Coins)

	// calculate the total reward for all validators
	for _, signatureCount := range signatureCounts.Counts {

		// get the operator address of the signature counts with the chain prefix
		config := sdk.GetConfig()
		if config == nil {
			return ignterrors.Critical("SDK config not set")
		}
		opAddr, err := signatureCount.GetOperatorAddress(config.GetBech32AccountAddrPrefix())
		if err != nil {
			return sdkerrors.Wrapf(types.ErrInvalidSignatureCounts, "invalid operator address: %s", signatureCount.OpAddress)
		}

		// if the operator address is associated with a validator profile, this address is used to receive rewwards
		// otherwise rewards are distributed to the operator address account
		valAddr := opAddr
		validatorByOpAddr, err := k.profileKeeper.GetValidatorByOperatorAddress(ctx, opAddr)
		if errors.Is(err, profiletypes.ErrValidatorByOperatorAddressNotFound) {
			valAddr = validatorByOpAddr.ValidatorAddress
		} else if err != nil {
			return ignterrors.Criticalf("can't get validator by operator address: %s", err.Error())
		}

		// calculate the total relative signature distributed to calculate the refund for the round
		totalRelativeSignaturesDistributed = totalRelativeSignaturesDistributed.Add(signatureCount.RelativeSignatures)

		// compute reward relative to the signature and block count
		// and update reward pool
		signatureRatio := signatureCount.RelativeSignatures.Quo(
			sdkmath.LegacyNewDecFromInt(sdkmath.NewIntFromUint64(signatureCounts.BlockCount)),
		)
		rewards, err := CalculateRewards(blockRatio, signatureRatio, rewardPool.RemainingCoins)
		if err != nil {
			return ignterrors.Criticalf("invalid reward: %s", err.Error())
		}
		rewardsToDistribute[valAddr] = rewards

	}

	// distribute the rewards to validators
	for address, rewards := range rewardsToDistribute {
		coins, isNegative := rewardPool.RemainingCoins.SafeSub(rewards...)
		if isNegative {
			return ignterrors.Criticalf("negative reward pool: %s", rewardPool.RemainingCoins.String())
		}
		rewardPool.RemainingCoins = coins

		// send rewards to the address
		account, err := k.addressCodec.StringToBytes(address)
		if err != nil {
			return ignterrors.Criticalf("can't parse address %s", err.Error())
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, account, rewards); err != nil {
			return ignterrors.Criticalf("send rewards error: %s", err.Error())
		}
		if err := sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventRewardsDistributed{
			LaunchID: launchID,
			Receiver: address,
			Rewards:  rewards,
		}); err != nil {
			return ignterrors.Criticalf("error emitting event: %s", err.Error())
		}
	}

	// if the reward pool is closed or last reward height is reached
	// the remaining coins are refunded and reward pool is deleted
	if closeRewardPool || lastBlockHeight >= rewardPool.LastRewardHeight {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			rewardPool.RemainingCoins,
		); err != nil {
			return ignterrors.Criticalf("send rewards error: %s", err.Error())
		}

		// close the pool
		rewardPool.Closed = true
		rewardPool.RemainingCoins = rewardPool.RemainingCoins.Sub(rewardPool.RemainingCoins...) // sub coins transferred
		return k.RewardPool.Set(ctx, rewardPool.LaunchID, rewardPool)
	}

	// Otherwise, the refund is relative to the block ratio and the reward pool is updated
	// refundRatio is blockCount.
	// This is sum of signaturesRelative values from validator to compute refund
	blockCount := sdkmath.LegacyNewDecFromInt(sdkmath.NewIntFromUint64(signatureCounts.BlockCount))
	refundRatioNumerator := blockCount.Sub(totalRelativeSignaturesDistributed)
	refundRatio := refundRatioNumerator.Quo(blockCount)
	refund, err := CalculateRewards(blockRatio, refundRatio, rewardPool.RemainingCoins)
	if err != nil {
		return ignterrors.Criticalf("invalid reward: %s", err.Error())
	}

	// if refund is non-null, refund is sent to the provider
	if !refund.IsZero() {
		coins, isNegative := rewardPool.RemainingCoins.SafeSub(refund...)
		if isNegative {
			return ignterrors.Criticalf("negative reward pool: %s", rewardPool.RemainingCoins.String())
		}
		rewardPool.RemainingCoins = coins

		// send rewards to the address
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			rewardPool.RemainingCoins,
		); err != nil {
			return ignterrors.Criticalf("send rewards error: %s", err.Error())
		}
	}

	// update the current reward height for next reward
	rewardPool.CurrentRewardHeight = lastBlockHeight
	return k.RewardPool.Set(ctx, rewardPool.LaunchID, rewardPool)
}

// CalculateRewards calculates the reward relative to the signature and block ratio
func CalculateRewards(blockRatio, signatureRatio sdkmath.LegacyDec, coins sdk.Coins) (sdk.Coins, error) {
	// ratio can't be greater than one
	if blockRatio.GT(sdkmath.LegacyOneDec()) {
		return nil, fmt.Errorf("block ratio is greater than 1 %s", blockRatio.String())
	}
	if signatureRatio.GT(sdkmath.LegacyOneDec()) {
		return nil, fmt.Errorf("signature ratio is greater than 1 %s", signatureRatio.String())
	}

	// if one ratio is zero, rewards are null
	if blockRatio.IsZero() || signatureRatio.IsZero() {
		return sdk.NewCoins(), nil
	}

	// calculate rewards
	rewards := sdk.NewCoins()
	for _, coin := range coins {
		amount := blockRatio.Mul(signatureRatio).Mul(sdkmath.LegacyNewDecFromInt(coin.Amount))
		coin.Amount = amount.TruncateInt()
		rewards = rewards.Add(coin)
	}
	return rewards, nil
}
