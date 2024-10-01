package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

// Requests returns all request.
func (k Keeper) Requests(ctx context.Context) ([]types.Request, error) {
	requests := make([]types.Request, 0)
	err := k.Request.Walk(ctx, nil, func(_ collections.Pair[uint64, uint64], value types.Request) (bool, error) {
		requests = append(requests, value)
		return false, nil
	})
	return requests, err
}

// GetRequestCounter get request counter for a specific chain ID
func (k Keeper) GetRequestCounter(ctx context.Context, launchID uint64) (uint64, error) {
	seq, err := k.RequestSeq.Get(ctx, launchID)
	if errors.Is(err, collections.ErrNotFound) {
		return 0, nil
	}
	return seq, err
}

// GetNextRequestIDWithUpdate increments bid id by one and set it.
func (k Keeper) GetNextRequestIDWithUpdate(ctx context.Context, launchID uint64) (uint64, error) {
	seq, err := k.RequestSeq.Get(ctx, launchID)
	if errors.Is(err, collections.ErrNotFound) {
		seq = 0
	} else if err != nil {
		return 0, err
	}
	seq++
	return seq, k.RequestSeq.Set(ctx, launchID, seq)
}

// AppendRequest appends a Request in the store with a new launch id and update the count
func (k Keeper) AppendRequest(ctx context.Context, request types.Request) (uint64, error) {
	requestID, err := k.GetNextRequestIDWithUpdate(ctx, request.LaunchID)
	if err != nil {
		return 0, ignterrors.Criticalf("failed to get next request sequence %s", err.Error())
	}
	request.RequestID = requestID
	if err := k.Request.Set(ctx, collections.Join(request.LaunchID, requestID), request); err != nil {
		return 0, ignterrors.Criticalf("request not set %s", err.Error())
	}
	return requestID, nil
}

// CheckAccount check account inconsistency and return
// if an account exists for genesis or vesting accounts
func CheckAccount(ctx context.Context, k Keeper, launchID uint64, address string) (bool, error) {
	accAddress, err := k.addressCodec.StringToBytes(address)
	if err != nil {
		return false, ignterrors.Criticalf("invalid bech32 format %s", address)
	}

	foundGenesis, err := k.GenesisAccount.Has(ctx, collections.Join(launchID, sdk.AccAddress(accAddress)))
	if err != nil {
		return false, err
	}
	foundVesting, err := k.VestingAccount.Has(ctx, collections.Join(launchID, sdk.AccAddress(accAddress)))
	if err != nil {
		return false, err
	}
	if foundGenesis && foundVesting {
		return false, ignterrors.Criticalf("account %s for chain %d found in vesting and genesis accounts", address, launchID)
	}
	return foundGenesis || foundVesting, nil
}

// ApplyRequest approves the request and performs
// the launch information changes
func ApplyRequest(
	ctx context.Context,
	k Keeper,
	chain types.Chain,
	request types.Request,
	coord profiletypes.Coordinator,
) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	err := CheckRequest(ctx, k, chain.LaunchID, request)
	if err != nil {
		return err
	}

	switch requestContent := request.Content.Content.(type) {
	case *types.RequestContent_GenesisAccount:
		ga := requestContent.GenesisAccount
		if !chain.AccountBalance.Empty() {
			ga.Coins = chain.AccountBalance
		}

		address, err := k.addressCodec.StringToBytes(ga.Address)
		if err != nil {
			return err
		}
		if err := k.GenesisAccount.Set(ctx, collections.Join(ga.LaunchID, sdk.AccAddress(address)), *ga); err != nil {
			return err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventGenesisAccountAdded{
			Address:            ga.Address,
			Coins:              ga.Coins,
			LaunchID:           chain.LaunchID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_VestingAccount:
		va := requestContent.VestingAccount
		if !chain.AccountBalance.Empty() {
			switch opt := va.VestingOptions.Options.(type) { //nolint:gocritic
			case *types.VestingOptions_DelayedVesting:
				dv := opt.DelayedVesting
				va = &types.VestingAccount{
					Address:  va.Address,
					LaunchID: va.LaunchID,
					VestingOptions: *types.NewDelayedVesting(
						chain.AccountBalance,
						chain.AccountBalance,
						dv.EndTime,
					),
				}
			}
		}

		address, err := k.addressCodec.StringToBytes(va.Address)
		if err != nil {
			return err
		}
		if err := k.VestingAccount.Set(ctx, collections.Join(va.LaunchID, sdk.AccAddress(address)), *va); err != nil {
			return err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventVestingAccountAdded{
			Address:            va.Address,
			VestingOptions:     va.VestingOptions,
			LaunchID:           chain.LaunchID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_AccountRemoval:
		ar := requestContent.AccountRemoval

		address, err := k.addressCodec.StringToBytes(ar.Address)
		if err != nil {
			return err
		}
		if err := k.GenesisAccount.Remove(ctx, collections.Join(chain.LaunchID, sdk.AccAddress(address))); err != nil && !errors.Is(err, collections.ErrNotFound) {
			return err
		}
		if err := k.VestingAccount.Remove(ctx, collections.Join(chain.LaunchID, sdk.AccAddress(address))); err != nil && !errors.Is(err, collections.ErrNotFound) {
			return err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventAccountRemoved{
			Address:            ar.Address,
			LaunchID:           chain.LaunchID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_GenesisValidator:
		ga := requestContent.GenesisValidator

		address, err := k.addressCodec.StringToBytes(ga.Address)
		if err != nil {
			return err
		}
		if err := k.GenesisValidator.Set(ctx, collections.Join(chain.LaunchID, sdk.AccAddress(address)), *ga); err != nil {
			return err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventValidatorAdded{
			Address:            ga.Address,
			GenTx:              ga.GenTx,
			ConsPubKey:         ga.ConsPubKey,
			SelfDelegation:     ga.SelfDelegation,
			Peer:               ga.Peer,
			LaunchID:           chain.LaunchID,
			HasProject:         chain.HasProject,
			ProjectID:          chain.ProjectID,
			CoordinatorAddress: coord.Address,
		})

	case *types.RequestContent_ValidatorRemoval:
		vr := requestContent.ValidatorRemoval

		address, err := k.addressCodec.StringToBytes(vr.ValAddress)
		if err != nil {
			return err
		}
		if err := k.GenesisValidator.Remove(ctx, collections.Join(chain.LaunchID, sdk.AccAddress(address))); err != nil && !errors.Is(err, collections.ErrNotFound) {
			return err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventValidatorRemoved{
			GenesisValidatorAccount: vr.ValAddress,
			LaunchID:                chain.LaunchID,
			HasProject:              chain.HasProject,
			ProjectID:               chain.ProjectID,
			CoordinatorAddress:      coord.Address,
		})

	case *types.RequestContent_ParamChange:
		cp := requestContent.ParamChange

		if err := k.ParamChange.Set(ctx, collections.Join(chain.LaunchID, types.ParamChangeSubKey(cp.Module, cp.Param)), *cp); err != nil {
			return err
		}

		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventParamChanged{
			LaunchID: cp.LaunchID,
			Module:   cp.Module,
			Param:    cp.Param,
			Value:    cp.Value,
		})

	}
	return err
}

// CheckRequest verifies that a request can be applied
func CheckRequest(
	ctx context.Context,
	k Keeper,
	launchID uint64,
	request types.Request,
) error {
	if err := request.Content.Validate(launchID); err != nil {
		return ignterrors.Critical(err.Error())
	}

	switch requestContent := request.Content.Content.(type) {
	case *types.RequestContent_GenesisAccount:
		ga := requestContent.GenesisAccount
		found, err := CheckAccount(ctx, k, launchID, ga.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %d already exist",
				ga.Address, launchID,
			)
		}
	case *types.RequestContent_VestingAccount:
		va := requestContent.VestingAccount
		found, err := CheckAccount(ctx, k, launchID, va.Address)
		if err != nil {
			return err
		}
		if found {
			return sdkerrors.Wrapf(types.ErrAccountAlreadyExist,
				"account %s for chain %d already exist",
				va.Address, launchID,
			)
		}
	case *types.RequestContent_AccountRemoval:
		ar := requestContent.AccountRemoval
		found, err := CheckAccount(ctx, k, launchID, ar.Address)
		if err != nil {
			return err
		}
		if !found {
			return sdkerrors.Wrapf(types.ErrAccountNotFound,
				"account %s for chain %d not found",
				ar.Address, launchID,
			)
		}
	case *types.RequestContent_GenesisValidator:
		ga := requestContent.GenesisValidator

		address, err := k.addressCodec.StringToBytes(ga.Address)
		if err != nil {
			return ignterrors.Criticalf("invalid bech32 format %s", address)
		}
		found, err := k.GenesisValidator.Has(ctx, collections.Join(launchID, sdk.AccAddress(address)))
		if err != nil {
			return ignterrors.Criticalf("invalid bech32 format %s", address)
		}
		if found {
			return sdkerrors.Wrapf(types.ErrValidatorAlreadyExist,
				"genesis validator %s for chain %d already exist",
				ga.Address, launchID,
			)
		}
	case *types.RequestContent_ValidatorRemoval:
		vr := requestContent.ValidatorRemoval

		address, err := k.addressCodec.StringToBytes(vr.ValAddress)
		if err != nil {
			return ignterrors.Criticalf("invalid bech32 format %s", address)
		}
		found, err := k.GenesisValidator.Has(ctx, collections.Join(launchID, sdk.AccAddress(address)))
		if err != nil {
			return ignterrors.Criticalf("invalid bech32 format %s", address)
		}
		if !found {
			return sdkerrors.Wrapf(types.ErrValidatorNotFound,
				"genesis validator %s for chain %d not found",
				vr.ValAddress, launchID,
			)
		}
	case *types.RequestContent_ParamChange:
		// currently no on-chain checks can be performed on change param
	}

	return nil
}
