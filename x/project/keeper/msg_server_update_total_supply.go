package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func (k msgServer) UpdateTotalSupply(ctx context.Context, msg *types.MsgUpdateTotalSupply) (*types.MsgUpdateTotalSupplyResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid authority address")
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get project params")
	}

	project, err := k.GetProject(ctx, msg.ProjectID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.ProjectID)
	}

	// Get the coordinator ID associated to the sender address
	coordinatorID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, coordinatorAddress)
	if err != nil {
		return nil, err
	}

	if project.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordinatorInvalid, fmt.Sprintf(
			"coordinator of the project is %d",
			project.CoordinatorID,
		))
	}

	if project.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%d", msg.ProjectID)
	}

	// Validate provided totalSupply
	if err := types.ValidateTotalSupply(msg.TotalSupplyUpdate, params.TotalSupplyRange); err != nil {
		if errors.Is(err, types.ErrInvalidSupplyRange) {
			return nil, ignterrors.Critical(err.Error())
		}
		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	project.TotalSupply = types.UpdateTotalSupply(project.TotalSupply, msg.TotalSupplyUpdate)
	if err := k.Project.Set(ctx, project.ProjectID, project); err != nil {
		return nil, ignterrors.Criticalf("project not set %s", err.Error())
	}

	return &types.MsgUpdateTotalSupplyResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventProjectTotalSupplyUpdated{
		ProjectID:          project.ProjectID,
		CoordinatorAddress: msg.Coordinator,
		TotalSupply:        project.TotalSupply,
	})
}
