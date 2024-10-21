package keeper

import (
	"context"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func (k msgServer) RevertLaunch(ctx context.Context, msg *types.MsgRevertLaunch) (*types.MsgRevertLaunchResponse, error) {
	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get launch params")
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
		return nil, sdkerrors.Wrapf(
			profiletypes.ErrCoordinatorInvalid,
			"coordinator of the chain is %d",
			chain.CoordinatorId,
		)
	}

	if !chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrNotTriggeredLaunch, "%d", msg.LaunchId)
	}

	if chain.MonitoringConnected {
		return nil, sdkerrors.Wrapf(types.ErrChainMonitoringConnected, "%d", msg.LaunchId)
	}

	// We must wait for a specific delay once the chain is launched before being able to revert it
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.BlockTime().Before(chain.LaunchTime.Add(params.RevertDelay)) {
		return nil, sdkerrors.Wrapf(types.ErrRevertDelayNotReached, "%d", msg.LaunchId)
	}

	chain.LaunchTriggered = false
	chain.LaunchTime = time.Unix(0, 0).UTC()
	if err := k.Chain.Set(ctx, chain.LaunchId, chain); err != nil {
		return nil, ignterrors.Criticalf("chain not set %s", err.Error())
	}

	// clear associated client IDs from monitoring
	if err := k.monitoringcKeeper.ClearVerifiedClientIdList(ctx, msg.LaunchId); err != nil {
		return nil, ignterrors.Criticalf("failed to clear monitoring client IDs %s", err.Error())
	}
	err = sdkCtx.EventManager().EmitTypedEvent(&types.EventLaunchReverted{
		LaunchId: msg.LaunchId,
	})

	return &types.MsgRevertLaunchResponse{}, err
}
