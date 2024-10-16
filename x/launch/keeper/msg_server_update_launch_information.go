package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func (k msgServer) UpdateLaunchInformation(ctx context.Context, msg *types.MsgUpdateLaunchInformation) (*types.MsgUpdateLaunchInformationResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
	}

	chain, err := k.GetChain(ctx, msg.LaunchID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
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

	// Modify from provided values
	if msg.GenesisChainID != "" {
		chain.GenesisChainID = msg.GenesisChainID
	}
	if msg.SourceURL != "" {
		chain.SourceURL = msg.SourceURL
	}
	if msg.SourceHash != "" {
		chain.SourceHash = msg.SourceHash
	}
	if msg.InitialGenesis != nil {
		chain.InitialGenesis = *msg.InitialGenesis
	}

	if err := k.Chain.Set(ctx, chain.LaunchID, chain); err != nil {
		return nil, ignterrors.Criticalf("chain not set %s", err.Error())
	}

	return &types.MsgUpdateLaunchInformationResponse{}, nil
}
