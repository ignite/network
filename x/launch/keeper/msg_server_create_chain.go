package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
)

func (k msgServer) CreateChain(ctx context.Context, msg *types.MsgCreateChain) (*types.MsgCreateChainResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
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

	launchID, err := k.CreateNewChain(
		ctx,
		coordinatorID,
		msg.GenesisChainId,
		msg.SourceUrl,
		msg.SourceHash,
		msg.InitialGenesis,
		msg.HasProject,
		msg.ProjectId,
		false,
		msg.AccountBalance,
		msg.Metadata,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreateChainFail, err.Error())
	}

	// Deduct chain creation fee if set
	creationFee := params.ChainCreationFee
	if !params.ChainCreationFee.Empty() {
		if err = k.distributionKeeper.FundCommunityPool(ctx, creationFee, coordinatorAddress); err != nil {
			return nil, sdkerrors.Wrap(types.ErrFundCommunityPool, err.Error())
		}
	}

	return &types.MsgCreateChainResponse{
		LaunchId: launchID,
	}, nil
}
