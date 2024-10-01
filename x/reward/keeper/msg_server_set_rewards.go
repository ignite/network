package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	launchtypes "github.com/ignite/network/x/launch/types"
	"github.com/ignite/network/x/reward/types"
)

func (k msgServer) SetRewards(ctx context.Context, msg *types.MsgSetRewards) (*types.MsgSetRewardsResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	provider, err := k.addressCodec.StringToBytes(msg.Provider)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid authority address")
	}

	// determine if the chain exists
	chain, err := k.launchKeeper.GetChain(ctx, msg.LaunchID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchID)
	}

	// check coordinator
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, provider)
	if err != nil {
		return nil, err
	}

	if chain.CoordinatorID != coordID {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCoordinatorID, "%d", coordID)
	}
	// reward can't be changed once launch is triggered
	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(launchtypes.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	var (
		previousCoins            sdk.Coins
		previousLastRewardHeight int64
		poolFound                bool
	)
	rewardPool, err := k.RewardPool.Get(ctx, msg.LaunchID)
	if errors.Is(err, collections.ErrNotFound) {
		poolFound = false
		// create the reward pool and transfer tokens if not created yet
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, provider, types.ModuleName, msg.Coins); err != nil {
			return nil, sdkerrors.Wrap(types.ErrInsufficientFunds, err.Error())
		}
		rewardPool = types.NewRewardPool(msg.LaunchID, 0)
	} else if err != nil {
		return nil, err
	} else {
		poolFound = true
		previousCoins = rewardPool.RemainingCoins
		previousLastRewardHeight = rewardPool.LastRewardHeight
		if err := SetBalance(ctx, k.bankKeeper, provider, msg.Coins, rewardPool.RemainingCoins); err != nil {
			return nil, err
		}
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if msg.Coins.Empty() || msg.LastRewardHeight == 0 {
		rewardPool.InitialCoins = sdk.NewCoins()
		rewardPool.RemainingCoins = sdk.NewCoins()
		rewardPool.LastRewardHeight = 0
		if err := k.RewardPool.Remove(ctx, msg.LaunchID); err != nil {
			return nil, err
		}
		if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventRewardPoolRemoved{LaunchID: msg.LaunchID}); err != nil {
			return nil, err
		}
	} else {
		rewardPool.InitialCoins = msg.Coins
		rewardPool.RemainingCoins = msg.Coins
		rewardPool.Provider = msg.Provider
		rewardPool.LastRewardHeight = msg.LastRewardHeight
		if err := k.RewardPool.Set(ctx, rewardPool.LaunchID, rewardPool); err != nil {
			return nil, err
		}
		if !poolFound {
			if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventRewardPoolCreated{
				LaunchID: rewardPool.LaunchID,
				Provider: rewardPool.Provider,
			}); err != nil {
				return nil, err
			}
		}
	}

	return &types.MsgSetRewardsResponse{
		PreviousCoins:            previousCoins,
		PreviousLastRewardHeight: previousLastRewardHeight,
		NewCoins:                 rewardPool.InitialCoins,
		NewLastRewardHeight:      rewardPool.LastRewardHeight,
	}, nil
}

// SetBalance set balance to Coins on the module account
// calling the transfer depending on the balance difference
func SetBalance(
	ctx context.Context,
	bankKeeper types.BankKeeper,
	provider sdk.AccAddress,
	coins sdk.Coins,
	poolCoins sdk.Coins,
) error {
	if coins.DenomsSubsetOf(poolCoins) && coins.Equal(poolCoins) {
		return nil
	}
	if poolCoins != nil && !poolCoins.IsZero() {
		if err := bankKeeper.SendCoinsFromModuleToAccount(
			ctx,
			types.ModuleName,
			provider,
			poolCoins,
		); err != nil {
			return ignterrors.Critical(err.Error())
		}
	}
	if coins != nil && !coins.IsZero() {
		if err := bankKeeper.SendCoinsFromAccountToModule(
			ctx,
			provider,
			types.ModuleName,
			coins,
		); err != nil {
			return sdkerrors.Wrap(types.ErrInsufficientFunds, err.Error())
		}
	}
	return nil
}
