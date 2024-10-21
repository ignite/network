package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

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
		vouchersTooBig         = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdkmath.NewInt(networktypes.TotalShareNumber+1)),
		)
	)

	shares, err := types.NewShares("1000foo,500bar,300foobar")
	require.NoError(t, err)

	// Set projects
	project.AllocatedShares = shares
	project.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, project)
	require.NoError(t, err)

	projectMainnetLaunched.MainnetInitialized = true
	projectMainnetLaunched.AllocatedShares = shares
	chainLaunched := sample.Chain(r, 0, 0)
	chainLaunched.LaunchTriggered = true
	chainLaunched.IsMainnet = true
	projectMainnetLaunched.MainnetId, err = tk.LaunchKeeper.AppendChain(ctx, chainLaunched)
	require.NoError(t, err)
	projectMainnetLaunched.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, projectMainnetLaunched)
	require.NoError(t, err)

	vouchers, err := types.SharesToVouchers(shares, project.ProjectId)
	require.NoError(t, err)

	invalidProjectID := uint64(10000)
	vouchersErr, err := types.SharesToVouchers(shares, invalidProjectID)
	require.NoError(t, err)

	vouchersMainnet, err := types.SharesToVouchers(shares, projectMainnetLaunched.ProjectId)
	require.NoError(t, err)

	t.Run("should allow setting test balances", func(t *testing.T) {
		err = tk.BankKeeper.MintCoins(ctx, types.ModuleName, vouchers)
		require.NoError(t, err)
		err = tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, vouchers)
		require.NoError(t, err)

		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(project.ProjectId, existAddr), types.MainnetAccount{
			ProjectId: project.ProjectId,
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
				ProjectId: project.ProjectId,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "should allow redeem voucher two",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectId: project.ProjectId,
				Vouchers:  sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "should allow redeem voucher three",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectId: project.ProjectId,
				Vouchers:  sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "should allow redeem all",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   sample.Address(r),
				ProjectId: project.ProjectId,
				Vouchers:  vouchers,
			},
		},
		{
			name: "should fail with non existing project",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectId: invalidProjectID,
				Vouchers:  vouchersErr,
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with invalid vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectId: project.ProjectId,
				Vouchers:  vouchersErr,
			},
			err: types.ErrNoMatchVouchers,
		},
		{
			name: "should fail with invalid sender address",
			msg: types.MsgRedeemVouchers{
				Sender:    "invalid_address",
				Account:   addr.String(),
				ProjectId: project.ProjectId,
				Vouchers:  vouchers,
			},
			err: types.ErrInvalidSigner,
		},
		{
			name: "should fail with insufficient funds",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectId: project.ProjectId,
				Vouchers:  vouchersTooBig,
			},
			err: types.ErrInsufficientVouchers,
		},

		{
			name: "should fail with account without funds for vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectId: project.ProjectId,
				Vouchers:  vouchers,
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "should fail with account without funds for voucher one",
			msg: types.MsgRedeemVouchers{
				Sender:    existAddr.String(),
				Account:   existAddr.String(),
				ProjectId: project.ProjectId,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientVouchers,
		},
		{
			name: "should fail with project with launched mainnet",
			msg: types.MsgRedeemVouchers{
				Sender:    addr.String(),
				Account:   addr.String(),
				ProjectId: projectMainnetLaunched.ProjectId,
				Vouchers:  vouchersMainnet,
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
				accountAddr, err = tk.ProjectKeeper.AddressCodec().StringToBytes(tc.msg.Account)
				require.NoError(t, err)

				previousAccount, err = tk.ProjectKeeper.GetMainnetAccount(ctx, tc.msg.ProjectId, accountAddr)
				if err != nil {
					require.ErrorIs(t, err, types.ErrMainnetAccountNotFound)
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

			shares, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.ProjectId)
			require.NoError(t, err)

			account, err := tk.ProjectKeeper.GetMainnetAccount(ctx, tc.msg.ProjectId, accountAddr)
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
