package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/types"
)

func TestMsgBurnVouchers(t *testing.T) {
	var (
		ctx, tk, ts    = testkeeper.NewTestSetup(t)
		project        = sample.Project(r, 0)
		addr           = sample.AccAddress(r)
		vouchersTooBig = sdk.NewCoins(
			sdk.NewCoin("v/0/foo", sdkmath.NewInt(networktypes.TotalShareNumber+1)),
		)
	)
	// Create shares
	shares, err := types.NewShares("1000foo,500bar,300foobar")
	require.NoError(t, err)

	// Set project
	project.AllocatedShares = shares
	project.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, project)
	require.NoError(t, err)

	// Create vouchers
	vouchers, err := types.SharesToVouchers(shares, project.ProjectID)
	require.NoError(t, err)

	invalidProjectID := uint64(1000)
	vouchersErr, err := types.SharesToVouchers(shares, invalidProjectID)
	require.NoError(t, err)

	t.Run("should allow setting initial balances", func(t *testing.T) {
		err = tk.BankKeeper.MintCoins(ctx, types.ModuleName, vouchers)
		require.NoError(t, err)
		err = tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, vouchers)
		require.NoError(t, err)
	})

	for _, tc := range []struct {
		name string
		msg  types.MsgBurnVouchers
		err  error
	}{
		{
			name: "should allow burn voucher",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
		},
		{
			name: "should allow burn voucher two",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[1]),
			},
		},
		{
			name: "should allow burn voucher three",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[2]),
			},
		},
		{
			name: "should fail for non existing project",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: invalidProjectID,
				Vouchers:  sdk.NewCoins(vouchersErr[0]),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail for invalid sender address",
			msg: types.MsgBurnVouchers{
				Sender:    "invalid_address",
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[1]),
			},
			err: types.ErrInvalidSigner,
		},
		{
			name: "should not burn more than allocated shares",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  vouchersTooBig,
			},
			err: types.ErrInsufficientVouchers,
		},

		{
			name: "should fail for insufficient funds for voucher one",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[0]),
			},
			err: types.ErrInsufficientVouchers,
		},

		{
			name: "should fail for insufficient funds for voucher two",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[1]),
			},
			err: types.ErrInsufficientVouchers,
		},

		{
			name: "should fail for insufficient funds for voucher three",
			msg: types.MsgBurnVouchers{
				Sender:    addr.String(),
				ProjectID: project.ProjectID,
				Vouchers:  sdk.NewCoins(vouchers[2]),
			},
			err: types.ErrInsufficientVouchers,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousProject types.Project
			var previousBalance sdk.Coins
			var creatorAddr sdk.AccAddress

			// Get values before message execution
			if tc.err == nil {
				previousProject, err = tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectID)
				require.NoError(t, err)

				creatorAddr, err = tk.ProjectKeeper.AddressCodec().StringToBytes(tc.msg.Sender)
				require.NoError(t, err)
				previousBalance = tk.BankKeeper.GetAllBalances(ctx, creatorAddr)
			}

			// Execute message
			_, err = ts.ProjectSrv.BurnVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			project, err := tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectID)
			require.NoError(t, err)

			// Allocated shares of the project must be decreased
			burned, err := types.VouchersToShares(tc.msg.Vouchers, tc.msg.ProjectID)
			require.NoError(t, err)

			expectedShares, err := types.DecreaseShares(previousProject.AllocatedShares, burned)
			require.NoError(t, err)
			require.True(t, types.IsEqualShares(expectedShares, project.AllocatedShares))

			// Check coordinator balance
			balance := tk.BankKeeper.GetAllBalances(ctx, creatorAddr)
			expectedBalance := previousBalance.Sub(tc.msg.Vouchers...)
			require.True(t, balance.Equal(expectedBalance))
		})
	}
}
