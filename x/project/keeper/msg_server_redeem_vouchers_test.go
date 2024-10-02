package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	ignterrors "github.com/ignite/network/pkg/errors"
	networktypes "github.com/ignite/network/pkg/types"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/types"
)

func TestMsgRedeemVouchers(t *testing.T) {
	var (
		ctx, tk, ts = testkeeper.NewTestSetup(t)

		addr                   = sample.AccAddress(r)
		existAddr              = sample.AccAddress(r)
		project                = sample.Project(r, 0)
		projectMainnetLaunched = sample.Project(r, 1)
		shares                 types.Shares
		vouchers               sdk.Coins
		err                    error
		vouchersTooBig         = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdkmath.NewInt(networktypes.TotalShareNumber+1)),
		)
	)

	t.Run("should allow creation of valid shares", func(t *testing.T) {
		shares, err = types.NewShares("1000foo,500bar,300foobar")
		require.NoError(t, err)
	})

	// Set projects
	project.AllocatedShares = shares
	project.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, project)
	require.NoError(t, err)

	projectMainnetLaunched.MainnetInitialized = true
	projectMainnetLaunched.AllocatedShares = shares
	chainLaunched := sample.Chain(r, 0, 0)
	chainLaunched.LaunchTriggered = true
	chainLaunched.IsMainnet = true
	projectMainnetLaunched.MainnetID, err = tk.LaunchKeeper.AppendChain(ctx, chainLaunched)
	require.NoError(t, err)
	projectMainnetLaunched.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, projectMainnetLaunched)
	require.NoError(t, err)

	t.Run("should allow creation of valid vouchers", func(t *testing.T) {
		vouchers, err = types.SharesToVouchers(shares, project.ProjectID)
		require.NoError(t, err)
	})

	t.Run("should allow setting test balances", func(t *testing.T) {
		err = tk.BankKeeper.MintCoins(ctx, types.ModuleName, vouchers)
		require.NoError(t, err)
		err = tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, vouchers)
		require.NoError(t, err)

		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(project.ProjectID, existAddr), types.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:   existAddr.String(),
			Shares:    shares,
		})
		require.NoError(t, err)
		err = tk.BankKeeper.MintCoins(ctx, types.ModuleName, vouchers)
		require.NoError(t, err)
		err = tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, existAddr, vouchers)
		require.NoError(t, err)
	})

	for _, tc := range []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "should allow redeem voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "should allow redeem voucher two",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "should allow redeem voucher three",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "should allow redeem all",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   sample.Address(r),
				ProjectID: project.ProjectID,
				Vouchers:  vouchers,
			},
		},
		{
			name: "should fail with non existing project",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: 10000,
				Vouchers:  sample.Coins(r),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with invalid vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sample.Coins(r),
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should fail with invalid sender address",
			msg: types.MsgRedeemVouchers{
				Sender:    "invalid_address",
				Account:   addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  vouchers,
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should fail with insufficient funds",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  vouchersTooBig,
			},
			err: types.ErrInsufficientVouchers,
		},

		{
			name: "should fail with account without funds for vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  vouchers,
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "should fail with account without funds for voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "should fail with project with launched mainnet",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectID: projectMainnetLaunched.ProjectID,
				Vouchers:  sample.Coins(r),
			},
			err: types.ErrMainnetLaunchTriggered,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousAccount types.MainnetAccount
			var previousBalance sdk.Coins
			var foundAccount bool
			var accountAddr sdk.AccAddress

			// Get values before message execution
			if tc.err == nil {
				accountAddr, err = sdk.AccAddressFromBech32(tc.msg.Account)
				require.NoError(t, err)

				previousAccount, err = tk.ProjectKeeper.GetMainnetAccount(ctx, tc.msg.ProjectID, accountAddr)
				if err != nil {
					require.ErrorIs(t, err, types.ErrAccountNotFound)
					foundAccount = false
				} else {
					foundAccount = true
				}
				if foundAccount {
					previousBalance = tk.BankKeeper.GetAllBalances(ctx, accountAddr)
				}
			}

			// Execute message
			_, err = ts.ProjectSrv.RedeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			shares, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.ProjectID)
			require.NoError(t, err)

			account, err := tk.ProjectKeeper.GetMainnetAccount(ctx, tc.msg.ProjectID, accountAddr)
			require.NoError(t, err)

			// Check account shares
			expectedShares := shares
			if foundAccount {
				expectedShares = types.IncreaseShares(previousAccount.Shares, shares)
			}
			require.True(t, types.IsEqualShares(expectedShares, account.Shares))

			// Check account balance
			expectedVouchers := sdk.Coins{}
			if foundAccount {
				var negative bool
				expectedVouchers, negative = previousBalance.SafeSub(tc.msg.Vouchers...)
				require.False(t, negative)
			}
			balance := tk.BankKeeper.GetAllBalances(ctx, accountAddr)
			require.True(t, expectedVouchers.Equal(balance))
		})
	}
}
