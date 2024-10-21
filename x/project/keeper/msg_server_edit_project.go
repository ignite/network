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

func (k msgServer) EditProject(ctx context.Context, msg *types.MsgEditProject) (*types.MsgEditProjectResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get project params")
	}

	// check if the metadata length is valid
	if uint64(len(msg.Metadata)) > params.MaxMetadataLength {
		return nil, sdkerrors.Wrapf(types.ErrInvalidMetadataLength,
			"metadata length %d is greater than maximum %d",
			len(msg.Metadata),
			params.MaxMetadataLength,
		)
	}

	project, err := k.GetProject(ctx, msg.ProjectId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.ProjectId)
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, coordinatorAddress)
	if err != nil {
		return nil, err
	}

	if project.CoordinatorId != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordinatorInvalid, fmt.Sprintf(
			"coordinator of the project is %d",
			project.CoordinatorId,
		))
	}

	if len(msg.Name) > 0 {
		project.ProjectName = msg.Name
	}

	if len(msg.Metadata) > 0 {
		project.Metadata = msg.Metadata
	}

	if err := k.Project.Set(ctx, project.ProjectId, project); err != nil {
		return nil, ignterrors.Criticalf("project not set %s", err.Error())
	}

	return &types.MsgEditProjectResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventProjectInfoUpdated{
		ProjectId:          project.ProjectId,
		CoordinatorAddress: msg.Coordinator,
		ProjectName:        project.ProjectName,
		Metadata:           project.Metadata,
	})
}
