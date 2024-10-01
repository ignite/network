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
		return nil, sdkerrors.Wrap(err, "invalid coordinator address")
	}

	chain, err := k.GetChain(ctx, msg.LaunchID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchID)
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

	if chain.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordinatorInvalid, fmt.Sprintf(
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		))
	}

	if len(msg.Metadata) > 0 {
		chain.Metadata = msg.Metadata
	}

	if msg.SetProjectID {
		// check if chain already has id associated
		if chain.HasProject {
			return nil, sdkerrors.Wrapf(types.ErrChainHasProject,
				"project with id %d already associated with chain %d",
				chain.ProjectID,
				chain.LaunchID,
			)
		}

		// check if chain coordinator is project coordinator
		project, err := k.projectKeeper.GetProject(ctx, msg.ProjectID)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "%d", msg.ProjectID)
		}

		if project.CoordinatorID != chain.CoordinatorID {
			return nil, sdkerrors.Wrapf(profiletypes.ErrCoordinatorInvalid,
				"coordinator of the project is %d, chain coordinator is %d",
				project.CoordinatorID,
				chain.CoordinatorID,
			)
		}

		chain.ProjectID = msg.ProjectID
		chain.HasProject = true

		err = k.projectKeeper.AddChainToProject(ctx, chain.ProjectID, chain.LaunchID)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrAddChainToProject, err.Error())
		}
	}

	if err := k.Chain.Set(ctx, chain.LaunchID, chain); err != nil {
		return nil, ignterrors.Criticalf("chain not set %s", err.Error())
	}

	return &types.MsgEditChainResponse{}, err
}
