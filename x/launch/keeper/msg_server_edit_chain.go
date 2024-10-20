package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func (k msgServer) EditChain(ctx context.Context, msg *types.MsgEditChain) (*types.MsgEditChainResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
	}

	chain, err := k.GetChain(ctx, msg.LaunchId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchId)
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get launch params")
	}

	// check if the metadata length is valid
	if uint64(len(msg.Metadata)) > params.MaxMetadataLength {
		return nil, sdkerrors.Wrapf(types.ErrInvalidMetadataLength,
			"metadata length %d is greater than maximum %d",
			len(msg.Metadata),
			params.MaxMetadataLength,
		)
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

	if len(msg.Metadata) > 0 {
		chain.Metadata = msg.Metadata
	}

	if msg.SetProjectId {
		// check if chain already has id associated
		if chain.HasProject {
			return nil, sdkerrors.Wrapf(types.ErrChainHasProject,
				"project with id %d already associated with chain %d",
				chain.ProjectId,
				chain.LaunchId,
			)
		}

		// check if chain coordinator is project coordinator
		project, err := k.projectKeeper.GetProject(ctx, msg.ProjectId)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "%d", msg.ProjectId)
		}

		if project.CoordinatorId != chain.CoordinatorId {
			return nil, sdkerrors.Wrapf(profiletypes.ErrCoordinatorInvalid,
				"coordinator of the project is %d, chain coordinator is %d",
				project.CoordinatorId,
				chain.CoordinatorId,
			)
		}

		chain.ProjectId = msg.ProjectId
		chain.HasProject = true

		err = k.projectKeeper.AddChainToProject(ctx, chain.ProjectId, chain.LaunchId)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrAddChainToProject, err.Error())
		}
	}

	if err := k.Chain.Set(ctx, chain.LaunchId, chain); err != nil {
		return nil, ignterrors.Criticalf("chain not set %s", err.Error())
	}

	return &types.MsgEditChainResponse{}, err
}
