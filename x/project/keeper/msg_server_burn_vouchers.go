package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/project/types"
)

func (k msgServer) BurnVouchers(ctx context.Context, msg *types.MsgBurnVouchers) (*types.MsgBurnVouchersResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	sender, err := k.addressCodec.StringToBytes(msg.Sender)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid sender address %s", err.Error())
	}

	project, err := k.GetProject(ctx, msg.ProjectID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.ProjectID)
	}

	// Convert and validate vouchers first
	shares, err := types.VouchersToShares(msg.Vouchers, msg.ProjectID)
	if err != nil {
		return nil, ignterrors.Criticalf("verified voucher are invalid %s", err.Error())
	}

	// Send coins and burn them
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, msg.Vouchers); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientVouchers, "%s", err.Error())
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Vouchers); err != nil {
		return nil, ignterrors.Criticalf("can't burn coins %s", err.Error())
	}

	// Decrease the project shares
	project.AllocatedShares, err = types.DecreaseShares(project.AllocatedShares, shares)
	if err != nil {
		return nil, ignterrors.Criticalf("invalid allocated share amount %s", err.Error())
	}
	if err := k.Project.Set(ctx, project.ProjectID, project); err != nil {
		return nil, ignterrors.Criticalf("can't set project %s", err.Error())
	}

	return &types.MsgBurnVouchersResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventProjectSharesUpdated{
		ProjectID:       project.ProjectID,
		AllocatedShares: project.AllocatedShares,
	})
}
