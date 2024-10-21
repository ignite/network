package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
	projecttypes "github.com/ignite/network/x/project/types"
)

func initCreationFeeAndFundCoordAccounts(
	t *testing.T,
	keeper *keeper.Keeper,
	bk bankkeeper.Keeper,
	ctx context.Context,
	fee sdk.Coins,
	numCreations int64,
	addrs ...string,
) {
	// set fee param to `coins`
	params, err := keeper.Params.Get(ctx)
	require.NoError(t, err)
	params.ChainCreationFee = fee
	err = keeper.Params.Set(ctx, params)
	require.NoError(t, err)

	coins := sdk.NewCoins()
	for _, coin := range fee {
		coin.Amount = coin.Amount.MulRaw(numCreations)
		coins = coins.Add(coin)
	}

	// add `coins` to balance of each coordinator address
	// using `project` module account for minting as `launch` does not have one
	for _, addr := range addrs {
		accAddr, err := keeper.AddressCodec().StringToBytes(addr)
		require.NoError(t, err)
		err = bk.MintCoins(ctx, projecttypes.ModuleName, coins)
		require.NoError(t, err)
		err = bk.SendCoinsFromModuleToAccount(ctx, projecttypes.ModuleName, accAddr, coins)
		require.NoError(t, err)
	}
}

func TestMsgCreateChain(t *testing.T) {
	var (
		coordAddrs       = make([]string, 5)
		coordMap         = make(map[string]uint64)
		prjtMap          = make(map[string]uint64)
		ctx, tk, ts      = testkeeper.NewTestSetup(t)
		chainCreationFee = sample.Coins(r)
	)

	// Create an invalid coordinator
	_, invalidCoordAddr := ts.CreateCoordinator(ctx, r)
	invalidCoordAddress := invalidCoordAddr.String()

	// Create coordinators
	for i := range coordAddrs {
		addr := sample.Address(r)
		coordAddrs[i] = addr
		coordMap[addr], _ = ts.CreateCoordinatorWithAddr(ctx, r, addr)
	}

	// Create a project for each valid coordinator
	for i := range coordAddrs {
		addr := coordAddrs[i]
		prjtMap[addr] = ts.CreateProject(ctx, r, addr)
	}

	// assign random sdk.Coins to `chainCreationFee` param and provide balance to coordinators
	// coordAddrs[4] is not funded
	initCreationFeeAndFundCoordAccounts(t, tk.LaunchKeeper, tk.BankKeeper, ctx, chainCreationFee, 1, coordAddrs[:4]...)

	// create message with an invalid metadata length
	msgCreateChainInvalidMetadata := sample.MsgCreateChain(
		r,
		coordAddrs[0],
		"",
		false,
		prjtMap[coordAddrs[0]],
	)
	launchParams, err := tk.LaunchKeeper.Params.Get(ctx)
	require.NoError(t, err)
	msgCreateChainInvalidMetadata.Metadata = sample.Metadata(r, launchParams.MaxMetadataLength+1)

	for _, tc := range []struct {
		name          string
		msg           types.MsgCreateChain
		wantedChainID uint64
		err           error
	}{
		{
			name:          "should create a chain",
			msg:           sample.MsgCreateChain(r, coordAddrs[0], "", false, prjtMap[coordAddrs[0]]),
			wantedChainID: 0,
		},
		{
			name:          "should allow creating a chain with a unique chain ID",
			msg:           sample.MsgCreateChain(r, coordAddrs[1], "", false, prjtMap[coordAddrs[1]]),
			wantedChainID: 1,
		},
		{
			name:          "should allow creating a chain with genesis url",
			msg:           sample.MsgCreateChain(r, coordAddrs[2], "foo.com", false, prjtMap[coordAddrs[2]]),
			wantedChainID: 2,
		},
		{
			name:          "should allow creating a chain with project",
			msg:           sample.MsgCreateChain(r, coordAddrs[3], "", true, prjtMap[coordAddrs[3]]),
			wantedChainID: 3,
		},
		{
			name: "should prevent creating a chain where coordinator doesn't exist for the chain",
			msg:  sample.MsgCreateChain(r, sample.Address(r), "", false, 0),
			err:  profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should prevent creating a chain with invalid project id",
			msg:  sample.MsgCreateChain(r, coordAddrs[0], "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
		{
			name: "should prevent creating a chain with invalid coordinator address",
			msg:  sample.MsgCreateChain(r, invalidCoordAddress, "", true, 1000),
			err:  types.ErrCreateChainFail,
		},
		{
			name: "should prevent creating a chain with insufficient balance to cover creation fee",
			msg:  sample.MsgCreateChain(r, coordAddrs[4], "", false, prjtMap[coordAddrs[4]]),
			err:  types.ErrFundCommunityPool,
		},
		{
			name: "should prevent a chain with invalid metadata length",
			msg:  msgCreateChainInvalidMetadata,
			err:  types.ErrInvalidMetadataLength,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// get account initial balance
			accAddr, err := tk.LaunchKeeper.AddressCodec().StringToBytes(tc.msg.Coordinator)
			require.NoError(t, err)
			preBalance := tk.BankKeeper.SpendableCoins(ctx, accAddr)

			got, err := ts.LaunchSrv.CreateChain(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.EqualValues(t, tc.wantedChainID, got.LaunchId)

			// The chain must exist in the store
			chain, err := tk.LaunchKeeper.GetChain(ctx, got.LaunchId)
			require.NoError(t, err)
			require.EqualValues(t, coordMap[tc.msg.Coordinator], chain.CoordinatorId)
			require.EqualValues(t, got.LaunchId, chain.LaunchId)
			require.EqualValues(t, tc.msg.GenesisChainId, chain.GenesisChainId)
			require.EqualValues(t, tc.msg.SourceUrl, chain.SourceUrl)
			require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)
			require.EqualValues(t, tc.msg.Metadata, chain.Metadata)
			require.EqualValues(t, tc.msg.InitialGenesis, chain.InitialGenesis)

			// Chain created from MsgCreateChain is never a mainnet
			require.False(t, chain.IsMainnet)

			require.Equal(t, tc.msg.HasProject, chain.HasProject)

			if tc.msg.HasProject {
				require.Equal(t, tc.msg.ProjectId, chain.ProjectId)
				projectChains, err := tk.ProjectKeeper.GetProjectChains(ctx, tc.msg.ProjectId)
				require.NoError(t, err)
				require.Contains(t, projectChains.Chains, chain.LaunchId)
			}

			// check fee deduction
			postBalance := tk.BankKeeper.SpendableCoins(ctx, accAddr)
			require.True(t, preBalance.Sub(chainCreationFee...).Equal(postBalance))
		})
	}
}
