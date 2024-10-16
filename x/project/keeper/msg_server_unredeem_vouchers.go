package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/project/types"
)

func (k msgServer) UnredeemVouchers(ctx context.Context, msg *types.MsgUnredeemVouchers) (*types.MsgUnredeemVouchersResponse, error) {
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

	mainnetLaunched, err := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
	if err != nil {
		return nil, ignterrors.Critical(err.Error())
	}
	if mainnetLaunched {
		return nil, sdkerrors.Wrap(types.ErrMainnetLaunchTriggered, fmt.Sprintf(
			"mainnet %d launch is already triggered",
			project.MainnetID,
		))
	}

	// Check if the account already exists
	account, err := k.MainnetAccount.Get(ctx, collections.Join(msg.ProjectID, sdk.AccAddress(sender)))
	if err != nil {
		return nil, ignterrors.Criticalf("can't get mainnet account %s", err.Error())
	}

	// Update the shares of the account
	account.Shares, err = types.DecreaseShares(account.Shares, msg.Shares)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrSharesDecrease, err.Error())
	}

	// If the account no longer has shares, it can be removed from the store
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if types.IsEqualShares(account.Shares, types.EmptyShares()) {
		if err := k.MainnetAccount.Remove(ctx, collections.Join(msg.ProjectID, sdk.AccAddress(sender))); err != nil {
			return nil, ignterrors.Criticalf("can't remove mainnet account %s", err.Error())
		}
		if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventMainnetAccountRemoved{
			ProjectID: project.ProjectID,
			Address:   account.Address,
		}); err != nil {
			return nil, err
		}
	} else {
		if err := k.MainnetAccount.Set(ctx, collections.Join(msg.ProjectID, sdk.AccAddress(sender)), account); err != nil {
			return nil, ignterrors.Criticalf("can't set mainnet account %s", err.Error())
		}
		if err := sdkCtx.EventManager().EmitTypedEvent(&types.EventMainnetAccountUpdated{
			ProjectID: account.ProjectID,
			Address:   account.Address,
			Shares:    account.Shares,
		}); err != nil {
			return nil, err
		}
	}

	// Mint vouchers from the removed shares and send them to sender balance
	vouchers, err := types.SharesToVouchers(msg.Shares, msg.ProjectID)
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, vouchers); err != nil {
		return nil, sdkerrors.Wrap(types.ErrVouchersMinting, err.Error())
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, vouchers); err != nil {
		return nil, ignterrors.Criticalf("can't send minted coins %s", err.Error())
	}

	return &types.MsgUnredeemVouchersResponse{}, nil
}
