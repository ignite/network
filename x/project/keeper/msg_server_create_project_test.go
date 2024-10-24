package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func initCreationFeeAndFundCoordAccounts(
	t *testing.T,
	keeper keeper.Keeper,
	bk bankkeeper.Keeper,
	ctx sdk.Context,
	fee sdk.Coins,
	numCreations int64,
	addrs ...string,
) {
	// set fee param to `coins`
	params, err := keeper.Params.Get(ctx)
	require.NoError(t, err)
	params.ProjectCreationFee = fee
	err = keeper.Params.Set(ctx, params)
	require.NoError(t, err)

	coins := sdk.NewCoins()
	for _, coin := range fee {
		coin.Amount = coin.Amount.MulRaw(numCreations)
		coins = coins.Add(coin)
	}

	t.Run("should add coins to balance of each coordinator address", func(t *testing.T) {
		for _, addr := range addrs {
			accAddr, err := keeper.AddressCodec().StringToBytes(addr)
			require.NoError(t, err)
			err = bk.MintCoins(ctx, types.ModuleName, coins)
			require.NoError(t, err)
			err = bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddr, coins)
			require.NoError(t, err)
		}
	})
}

func TestMsgCreateProject(t *testing.T) {
	var (
		coordAddrs         = make([]string, 3)
		coordinatorMap     = make(map[string]uint64)
		ctx, tk, ts        = testkeeper.NewTestSetup(t)
		projectCreationFee = sample.Coins(r)
	)

	params, err := tk.ProjectKeeper.Params.Get(ctx)
	require.NoError(t, err)
	maxMetadataLength := params.MaxMetadataLength

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		for i := range coordAddrs {
			addr := sample.Address(r)
			coordAddrs[i] = addr
			coordinatorMap[addr], _ = ts.CreateCoordinatorWithAddr(ctx, r, addr)
		}
	})

	// assign random sdk.Coins to `projectCreationFee` param and provide balance to coordinators
	// coordAddrs[2] is not funded
	initCreationFeeAndFundCoordAccounts(t, tk.ProjectKeeper, tk.BankKeeper, ctx, projectCreationFee, 1, coordAddrs[:2]...)

	for _, tc := range []struct {
		name       string
		msg        types.MsgCreateProject
		expectedID uint64
		err        error
	}{
		{
			name: "should allow create a project 1",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[0],
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			expectedID: uint64(0),
		},
		{
			name: "should allow create a project from a different coordinator",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[1],
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			expectedID: uint64(1),
		},
		{
			name: "should allow create a project from a non existing coordinator",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: sample.Address(r),
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should allow create a project with an invalid token supply",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[0],
				TotalSupply: sample.CoinsWithRange(r, 10, 20),
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "should fail for insufficient balance to cover creation fee",
			msg: types.MsgCreateProject{
				ProjectName: sample.ProjectName(r),
				Coordinator: coordAddrs[2],
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrFundCommunityPool,
		},
		{
			name: "should fail when the change had too long metadata",
			msg: types.MsgCreateProject{
				Coordinator: sample.Address(r),
				ProjectName: sample.ProjectName(r),
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, maxMetadataLength+1),
			},
			err: types.ErrInvalidMetadataLength,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// get account initial balance
			accAddr, err := tk.ProjectKeeper.AddressCodec().StringToBytes(tc.msg.Coordinator)
			require.NoError(t, err)
			preBalance := tk.BankKeeper.SpendableCoins(ctx, accAddr)

			got, err := ts.ProjectSrv.CreateProject(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expectedID, got.ProjectId)
			project, err := tk.ProjectKeeper.GetProject(ctx, got.ProjectId)
			require.NoError(t, err)
			require.EqualValues(t, got.ProjectId, project.ProjectId)
			require.EqualValues(t, tc.msg.ProjectName, project.ProjectName)
			require.EqualValues(t, coordinatorMap[tc.msg.Coordinator], project.CoordinatorId)
			require.False(t, project.MainnetInitialized)
			require.True(t, tc.msg.TotalSupply.Equal(project.TotalSupply))
			require.EqualValues(t, types.Shares(nil), project.AllocatedShares)

			// Empty list of project chains
			projectChains, err := tk.ProjectKeeper.GetProjectChains(ctx, got.ProjectId)
			require.NoError(t, err)
			require.EqualValues(t, got.ProjectId, projectChains.ProjectId)
			require.Empty(t, projectChains.Chains)

			// check fee deduction
			postBalance := tk.BankKeeper.SpendableCoins(ctx, accAddr)
			require.True(t, preBalance.Sub(projectCreationFee...).Equal(postBalance))
		})
	}
}
