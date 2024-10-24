package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func (k msgServer) TriggerLaunch(ctx context.Context, msg *types.MsgTriggerLaunch) (*types.MsgTriggerLaunchResponse, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get launch params")
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
	}

	chain, err := k.GetChain(ctx, msg.LaunchId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchId)
	}

	// Get the coordinator ID associated to the sender address
	coordinatorID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, coordinatorAddress)
	if err != nil {
		return nil, err
	}

	if chain.CoordinatorId != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordinatorInvalid, fmt.Sprintf(
			"coordinator of the chain is %d",
			chain.CoordinatorId,
		))
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchId)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()
	launchTime := blockTime.Add(msg.LaunchTime)
	if launchTime.Before(blockTime.Add(params.LaunchTimeRange.MinLaunchTime)) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooLow, "%s", msg.LaunchTime.String())
	}
	if launchTime.After(blockTime.Add(params.LaunchTimeRange.MaxLaunchTime)) {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooHigh, "%s", msg.LaunchTime.String())
	}

	// set launch timestamp
	chain.LaunchTriggered = true
	chain.LaunchTime = launchTime

	// set revision height for monitoring IBC client
	chain.ConsumerRevisionHeight = sdkCtx.BlockHeight()

	if err := k.Chain.Set(ctx, chain.LaunchId, chain); err != nil {
		return nil, ignterrors.Criticalf("chain not set %s", err.Error())
	}

	err = sdkCtx.EventManager().EmitTypedEvent(&types.EventLaunchTriggered{
		LaunchId:        msg.LaunchId,
		LaunchTimestamp: chain.LaunchTime,
	})

	return &types.MsgTriggerLaunchResponse{}, err
}
