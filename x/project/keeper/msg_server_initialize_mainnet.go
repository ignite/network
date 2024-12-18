package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	launchtypes "github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func (k msgServer) InitializeMainnet(ctx context.Context, msg *types.MsgInitializeMainnet) (*types.MsgInitializeMainnetResponse, error) {
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

	if project.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%d", msg.ProjectId)
	}

	if project.TotalSupply.Empty() {
		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, "total supply is empty")
	}

	// Get the coordinator ID associated to the sender address
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

	initialGenesis := launchtypes.NewDefaultInitialGenesis()

	// Create the mainnet chain for launch
	mainnetID, err := k.launchKeeper.CreateNewChain(
		ctx,
		coordinatorID,
		msg.MainnetChainId,
		msg.SourceUrl,
		msg.SourceHash,
		initialGenesis,
		true,
		msg.ProjectId,
		true,
		sdk.NewCoins(), // no enforced default for mainnet
		[]byte{},
	)
	if err != nil {
		return nil, ignterrors.Criticalf("cannot create the mainnet: %s", err.Error())
	}

	// Set mainnet as initialized and save the change
	project.MainnetId = mainnetID
	project.MainnetInitialized = true
	if err := k.Project.Set(ctx, project.ProjectId, project); err != nil {
		return nil, ignterrors.Criticalf("project not set %s", err.Error())
	}

	return &types.MsgInitializeMainnetResponse{
			MainnetId: mainnetID,
		}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventProjectMainnetInitialized{
			ProjectId:          project.ProjectId,
			CoordinatorAddress: msg.Coordinator,
			MainnetId:          project.MainnetId,
		})
}
