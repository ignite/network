package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func TestMsgInitializeMainnet(t *testing.T) {
	var (
		coordID                     uint64
		projectID                   uint64 = 0
		projectMainnetInitializedID uint64 = 1
		projectIncorrectCoordID     uint64 = 2
		projectEmptySupplyID        uint64 = 3
		coordAddr                          = sample.Address(r)
		coordAddrNoProject                 = sample.Address(r)

		ctx, tk, ts = testkeeper.NewTestSetup(t)
	)

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorID
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddrNoProject,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	project := sample.Project(r, projectID)
	project.CoordinatorID = coordID
	err := tk.ProjectKeeper.Project.Set(ctx, projectID, project)
	require.NoError(t, err)

	projectMainnetInitialized := sample.Project(r, projectMainnetInitializedID)
	projectMainnetInitialized.CoordinatorID = coordID
	projectMainnetInitialized.MainnetInitialized = true
	err = tk.ProjectKeeper.Project.Set(ctx, projectMainnetInitializedID, projectMainnetInitialized)
	require.NoError(t, err)

	projectEmptySupply := sample.Project(r, projectEmptySupplyID)
	projectEmptySupply.CoordinatorID = coordID
	projectEmptySupply.TotalSupply = sdk.NewCoins()
	err = tk.ProjectKeeper.Project.Set(ctx, projectEmptySupplyID, projectEmptySupply)
	require.NoError(t, err)

	projectIncorrectCoord := sample.Project(r, projectIncorrectCoordID)
	projectIncorrectCoord.CoordinatorID = coordID
	err = tk.ProjectKeeper.Project.Set(ctx, projectIncorrectCoordID, projectIncorrectCoord)
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgInitializeMainnet
		err  error
	}{
		{
			name: "should allow initialize mainnet",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
		},
		{
			name: "should fail if project not found",
			msg: types.MsgInitializeMainnet{
				ProjectID:      1000,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail if mainnet already initialized",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectMainnetInitializedID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "should fail if project has empty supply",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectEmptySupplyID,
				Coordinator:    coordAddr,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "should fail with non-existent coordinator",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectIncorrectCoordID,
				Coordinator:    sample.Address(r),
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should fail with invalid coordinator",
			msg: types.MsgInitializeMainnet{
				ProjectID:      projectIncorrectCoordID,
				Coordinator:    coordAddrNoProject,
				SourceHash:     sample.String(r, 30),
				SourceURL:      sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: profiletypes.ErrCoordinatorInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ts.ProjectSrv.InitializeMainnet(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			project, err := tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectID)
			require.NoError(t, err)
			require.True(t, project.MainnetInitialized)
			require.EqualValues(t, res.MainnetID, project.MainnetID)

			// Chain is in launch module
			chain, err := tk.LaunchKeeper.GetChain(ctx, project.MainnetID)
			require.NoError(t, err)
			require.True(t, chain.HasProject)
			require.True(t, chain.IsMainnet)
			require.EqualValues(t, tc.msg.ProjectID, chain.ProjectID)

			// Mainnet ID is listed in project chains
			projectChains, err := tk.ProjectKeeper.GetProjectChains(ctx, tc.msg.ProjectID)
			require.NoError(t, err)
			require.Contains(t, projectChains.Chains, project.MainnetID)
		})
	}
}
