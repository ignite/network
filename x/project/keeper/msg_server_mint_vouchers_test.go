package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func TestMsgMintVouchers(t *testing.T) {
	var (
		ctx, tk, ts    = testkeeper.NewTestSetup(t)
		coordID        uint64
		coord          = sample.Address(r)
		coordNoProject = sample.Address(r)

		shares, _    = types.NewShares("1000foo,500bar,300foobar")
		sharesTooBig = types.NewSharesFromCoins(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(networktypes.TotalShareNumber+1)),
		))
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coord,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorId
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordNoProject,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	// Set project
	project := sample.Project(r, 0)
	project.CoordinatorId = coordID
	projectID, err := tk.ProjectKeeper.AppendProject(ctx, project)
	require.NoError(t, err)
	project.ProjectId = projectID

	for _, tc := range []struct {
		name string
		msg  types.MsgMintVouchers
		err  error
	}{
		{
			name: "should allow minting  vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectId:   0,
				Shares:      shares,
			},
		},
		{
			name: "should allow minting same vouchers again",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectId:   0,
				Shares:      shares,
			},
		},
		{
			name: "should allow minting other vouchers",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectId:   0,
				Shares:      sample.Shares(r),
			},
		},
		{
			name: "should not mint more than total shares",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectId:   0,
				Shares:      sharesTooBig,
			},
			err: types.ErrTotalSharesLimit,
		},
		{
			name: "should fail with non existing project",
			msg: types.MsgMintVouchers{
				Coordinator: coord,
				ProjectId:   1000,
				Shares:      shares,
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with non existing coordinator",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectId:   0,
				Shares:      shares,
			},
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should fail with invalid coordinator",
			msg: types.MsgMintVouchers{
				Coordinator: coordNoProject,
				ProjectId:   0,
				Shares:      shares,
			},
			err: profiletypes.ErrCoordinatorInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousProject types.Project
			var previousBalance sdk.Coins

			coordAddr, err := tk.ProjectKeeper.AddressCodec().StringToBytes(tc.msg.Coordinator)
			require.NoError(t, err)

			// Get values before message execution
			if tc.err == nil {
				previousProject, err = tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectId)
				require.NoError(t, err)

				previousBalance = tk.BankKeeper.GetAllBalances(ctx, coordAddr)
			}

			// Execute message
			_, err = ts.ProjectSrv.MintVouchers(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			project, err := tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectId)
			require.NoError(t, err)

			// Allocated shares of the project must be increased
			expectedShares := types.IncreaseShares(previousProject.AllocatedShares, tc.msg.Shares)
			require.True(t, types.IsEqualShares(expectedShares, project.AllocatedShares))

			// Check coordinator balance
			minted, err := types.SharesToVouchers(tc.msg.Shares, tc.msg.ProjectId)
			require.NoError(t, err)
			balance := tk.BankKeeper.GetAllBalances(ctx, coordAddr)
			require.True(t, balance.Equal(previousBalance.Add(minted...)))
		})
	}
}
