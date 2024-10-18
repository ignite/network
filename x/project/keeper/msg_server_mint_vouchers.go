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

func (k msgServer) MintVouchers(ctx context.Context, msg *types.MsgMintVouchers) (*types.MsgMintVouchersResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	coordinatorAddress, err := k.addressCodec.StringToBytes(msg.Coordinator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid coordinator address %s", err.Error())
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

	// Increase the project shares
	totalShares, err := k.TotalShares.Get(ctx)
	if err != nil {
		return nil, ignterrors.Criticalf("total shares not found %s", err.Error())
	}

	project.AllocatedShares = types.IncreaseShares(project.AllocatedShares, msg.Shares)
	reached, err := types.IsTotalSharesReached(project.AllocatedShares, totalShares)
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if reached {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.ProjectID)
	}

	// Mint vouchers to the coordinator account
	vouchers, err := types.SharesToVouchers(msg.Shares, msg.ProjectID)
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, vouchers); err != nil {
		return nil, sdkerrors.Wrap(types.ErrVouchersMinting, err.Error())
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, coordinatorAddress, vouchers); err != nil {
		return nil, ignterrors.Criticalf("can't send minted coins %s", err.Error())
	}

	if err := k.Project.Set(ctx, project.ProjectID, project); err != nil {
		return nil, ignterrors.Criticalf("can't set project %s", err.Error())
	}

	return &types.MsgMintVouchersResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
		&types.EventProjectSharesUpdated{
			ProjectID:          project.ProjectID,
			CoordinatorAddress: msg.Coordinator,
			AllocatedShares:    project.AllocatedShares,
		})
}
