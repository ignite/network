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

	chain, err := k.GetChain(ctx, msg.LaunchId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchId)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchId)
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

	// Modify from provided values
	if msg.GenesisChainId != "" {
		chain.GenesisChainId = msg.GenesisChainId
	}
	if msg.SourceUrl != "" {
		chain.SourceUrl = msg.SourceUrl
	}
	if msg.SourceHash != "" {
		chain.SourceHash = msg.SourceHash
	}
	if msg.InitialGenesis != nil {
		chain.InitialGenesis = *msg.InitialGenesis
	}

	if err := k.Chain.Set(ctx, chain.LaunchId, chain); err != nil {
		return nil, ignterrors.Criticalf("chain not set %s", err.Error())
	}

	return &types.MsgUpdateLaunchInformationResponse{}, nil
}
