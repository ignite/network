package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func (k msgServer) UpdateSpecialAllocations(ctx context.Context, msg *types.MsgUpdateSpecialAllocations) (*types.MsgUpdateSpecialAllocationsResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
	}

	project, err := k.GetProject(ctx, msg.ProjectId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.ProjectId)
	}

	// get the coordinator ID associated to the sender address
	coordinatorID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, coordinatorAddress)
	if err != nil {
		return nil, err
	}

	if project.CoordinatorId != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordinatorInvalid, fmt.Sprintf(
			"coordinator of the project is %d",
			project.CoordinatorId,
		))
	}

	// verify mainnet launch is not triggered
	mainnetLaunched, err := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectId)
	if err != nil {
		return nil, ignterrors.Critical(err.Error())
	}
	if mainnetLaunched {
		return nil, sdkerrors.Wrap(types.ErrMainnetLaunchTriggered, fmt.Sprintf(
			"mainnet %d launch is already triggered",
			project.MainnetId,
		))
	}

	// decrease allocated shares from current special allocations
	project.AllocatedShares, err = types.DecreaseShares(project.AllocatedShares, project.SpecialAllocations.TotalShares())
	if err != nil {
		return nil, ignterrors.Critical("project allocated shares should be bigger than current special allocations" + err.Error())
	}

	// increase with new special allocations
	project.AllocatedShares = types.IncreaseShares(project.AllocatedShares, msg.SpecialAllocations.TotalShares())

	// increase the project shares
	totalShares, err := k.TotalShares.Get(ctx)
	if err != nil {
		return nil, ignterrors.Criticalf("total shares not found %s", err.Error())
	}

	reached, err := types.IsTotalSharesReached(project.AllocatedShares, totalShares)
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if reached {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.ProjectId)
	}

	project.SpecialAllocations = msg.SpecialAllocations
	if err := k.Project.Set(ctx, project.ProjectId, project); err != nil {
		return nil, ignterrors.Criticalf("project not set %s", err.Error())
	}

	return &types.MsgUpdateSpecialAllocationsResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvents(
		&types.EventProjectSharesUpdated{
			ProjectId:          project.ProjectId,
			CoordinatorAddress: msg.Coordinator,
			AllocatedShares:    project.AllocatedShares,
		})
}
