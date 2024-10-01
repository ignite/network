package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/project/types"
)

func (k msgServer) CreateProject(ctx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
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

	// Validate provided totalSupply
	if err := types.ValidateTotalSupply(msg.TotalSupply, params.TotalSupplyRange); err != nil {
		if errors.Is(err, types.ErrInvalidSupplyRange) {
			return nil, ignterrors.Critical(err.Error())
		}

		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	// Deduct project creation fee if set
	creationFee := params.ProjectCreationFee
	if !creationFee.Empty() {
		if err = k.distributionKeeper.FundCommunityPool(ctx, creationFee, coordinatorAddress); err != nil {
			return nil, sdkerrors.Wrap(types.ErrFundCommunityPool, err.Error())
		}
	}

	// Append the new project
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	project := types.NewProject(
		0,
		msg.ProjectName,
		coordinatorID,
		msg.TotalSupply,
		msg.Metadata,
		sdkCtx.BlockTime().Unix(),
	)

	projectID, err := k.AppendProject(ctx, project)
	if err != nil {
		return nil, err
	}

	// Initialize the list of project chains
	if err := k.ProjectChains.Set(ctx, projectID, types.ProjectChains{
		ProjectID: projectID,
		Chains:    []uint64{},
	}); err != nil {
		return nil, ignterrors.Criticalf("project chains not set %s", err.Error())
	}

	return &types.MsgCreateProjectResponse{ProjectID: projectID}, sdkCtx.EventManager().EmitTypedEvent(&types.EventProjectCreated{
		ProjectID:          projectID,
		CoordinatorAddress: msg.Coordinator,
		CoordinatorID:      coordinatorID,
	})
}
