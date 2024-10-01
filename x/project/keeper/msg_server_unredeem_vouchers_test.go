package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/types"
)

func TestMsgUnredeemVouchers(t *testing.T) {
	var (
		ctx, tk, ts = testkeeper.NewTestSetup(t)

		accountAddr              = sample.Address(r)
		account                  = sample.MainnetAccount(r, 0, accountAddr)
		accountShare, _          = types.NewShares("30foo,30bar")
		accountFewSharesAddr     = sample.Address(r)
		accountFewShares         = sample.MainnetAccount(r, 0, accountFewSharesAddr)
		accountFewSharesShare, _ = types.NewShares("30foo,15bar")

		project                = sample.Project(r, 0)
		projectMainnetLaunched = sample.Project(r, 1)
		shares, _              = types.NewShares("10foo,10bar")
	)
	account.Shares = accountShare
	accountFewShares.Shares = accountFewSharesShare

	// Create projects
	projectID, err := tk.ProjectKeeper.AppendProject(ctx, project)
	require.NoError(t, err)

	projectMainnetLaunched.MainnetInitialized = true
	chainLaunched := sample.Chain(r, 0, 0)
	chainLaunched.LaunchTriggered = true
	chainLaunched.IsMainnet = true
	projectMainnetLaunched.MainnetID, err = tk.LaunchKeeper.AppendChain(ctx, chainLaunched)
	require.NoError(t, err)
	projectMainnetLaunched.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, projectMainnetLaunched)
	require.NoError(t, err)

	// Create accounts
	accountAddress, err := sdk.AccAddressFromBech32(account.Address)
	require.NoError(t, err)
	err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(projectID, accountAddress), account)
	require.NoError(t, err)

	accountFewSharesAddress, err := sdk.AccAddressFromBech32(accountFewShares.Address)
	require.NoError(t, err)
	err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(projectID, accountFewSharesAddress), accountFewShares)
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgUnredeemVouchers
		err  error
	}{
		{
			name: "should allow unredeem vouchers",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should allow unredeem vouchers a second time",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should allow unredeem vouchers to zero",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should allow unredeem vouchers from another account",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountFewSharesAddr,
				ProjectID: 0,
				Shares:    shares,
			},
		},
		{
			name: "should prevent if not enough shares in balance",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountFewSharesAddr,
				ProjectID: 0,
				Shares:    shares,
			},
			err: types.ErrSharesDecrease,
		},
		{
			name: "should prevent for non existent project",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: 1000,
				Shares:    shares,
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should prevent for non existent account",
			msg: types.MsgUnredeemVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Shares:    shares,
			},
			err: types.ErrAccountNotFound,
		},
		{
			name: "should prevent for project with launched mainnet",
			msg: types.MsgUnredeemVouchers{
				Sender:    accountAddr,
				ProjectID: projectMainnetLaunched.ProjectID,
				Shares:    sample.Shares(r),
			},
			err: types.ErrMainnetLaunchTriggered,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousAccount types.MainnetAccount
			var previousBalance sdk.Coins

			accountAddr, err := sdk.AccAddressFromBech32(tc.msg.Sender)
			require.NoError(t, err)

			// Get values before message execution
			if tc.err == nil {
				previousAccount, err = tk.ProjectKeeper.GetMainnetAccount(ctx, tc.msg.ProjectID, accountAddr)
				require.NoError(t, err)

				previousBalance = tk.BankKeeper.GetAllBalances(ctx, accountAddr)
			}

			// Execute message
			_, err = ts.ProjectSrv.UnredeemVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			if types.IsEqualShares(tc.msg.Shares, previousAccount.Shares) {
				// All unredeemed
				_, err := tk.ProjectKeeper.GetMainnetAccount(ctx, tc.msg.ProjectID, accountAddr)
				require.NoError(t, err)

			} else {
				account, err := tk.ProjectKeeper.GetMainnetAccount(ctx, tc.msg.ProjectID, accountAddr)
				require.NoError(t, err)

				expectedShares, err := types.DecreaseShares(previousAccount.Shares, tc.msg.Shares)
				require.NoError(t, err)
				require.True(t, types.IsEqualShares(expectedShares, account.Shares))
			}

			// Compare balance
			unredeemed, err := types.SharesToVouchers(tc.msg.Shares, tc.msg.ProjectID)
			require.NoError(t, err)
			balance := tk.BankKeeper.GetAllBalances(ctx, accountAddr)
			require.True(t, balance.Equal(previousBalance.Add(unredeemed...)))
		})
	}
}
