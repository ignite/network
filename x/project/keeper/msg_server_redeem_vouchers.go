package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/project/types"
)

func (k msgServer) RedeemVouchers(ctx context.Context, msg *types.MsgRedeemVouchers) (*types.MsgRedeemVouchersResponse, error) {
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

	// Convert and validate vouchers first
	shares, err := types.VouchersToShares(msg.Vouchers, msg.ProjectID)
	if err != nil {
		return nil, ignterrors.Criticalf("verified voucher are invalid %s", err.Error())
	}

	// Send coins and burn them
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, msg.Vouchers); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientVouchers, "%s", sender)
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Vouchers); err != nil {
		return nil, ignterrors.Criticalf("can't burn coins %s", err.Error())
	}

	// Check if the account already exists

	accountAddress, err := k.addressCodec.StringToBytes(msg.Account)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid account address %s", err.Error())
	}

	found := true
	account, err := k.MainnetAccount.Get(ctx, collections.Join(msg.ProjectID, sdk.AccAddress(accountAddress)))
	if errors.Is(err, collections.ErrNotFound) {
		found = false
		// If not, create the account
		account = types.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:   msg.Account,
			Shares:    types.EmptyShares(),
		}
	} else if err != nil {
		return nil, ignterrors.Criticalf("can't get mainnet account %s", err.Error())
	}

	// Increase the account shares
	account.Shares = types.IncreaseShares(account.Shares, shares)
	if err := k.MainnetAccount.Set(ctx, collections.Join(msg.ProjectID, sdk.AccAddress(accountAddress)), account); err != nil {
		return nil, ignterrors.Criticalf("can't set mainnet account %s", err.Error())
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if !found {
		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventMainnetAccountCreated{
			ProjectID: account.ProjectID,
			Address:   account.Address,
			Shares:    account.Shares,
		})
	} else {
		err = sdkCtx.EventManager().EmitTypedEvent(&types.EventMainnetAccountUpdated{
			ProjectID: account.ProjectID,
			Address:   account.Address,
			Shares:    account.Shares,
		})
	}

	return &types.MsgRedeemVouchersResponse{}, err
}
