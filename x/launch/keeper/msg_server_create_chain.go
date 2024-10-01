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
		return nil, sdkerrors.Wrap(err, "invalid coordinator address")
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
		msg.GenesisChainID,
		msg.SourceURL,
		msg.SourceHash,
		msg.InitialGenesis,
		msg.HasProject,
		msg.ProjectID,
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
		coordAddr, err := k.addressCodec.StringToBytes(msg.Coordinator)
		if err != nil {
			return nil, ignterrors.Criticalf("invalid coordinator bech32 address %s", err.Error())
		}
		if err = k.distributionKeeper.FundCommunityPool(ctx, creationFee, coordAddr); err != nil {
			return nil, sdkerrors.Wrap(types.ErrFundCommunityPool, err.Error())
		}
	}

	return &types.MsgCreateChainResponse{
		LaunchID: launchID,
	}, nil
}
