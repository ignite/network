package profile_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	profile "github.com/ignite/network/x/profile/module"
	"github.com/ignite/network/x/profile/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CoordinatorList: []types.Coordinator{
			{
				CoordinatorID: 0,
			},
			{
				CoordinatorID: 1,
			},
		},
		CoordinatorCount: 2,
		ValidatorList: []types.Validator{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.ProfileKeeper(t)
	err := profile.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := profile.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CoordinatorList, got.CoordinatorList)
	require.Equal(t, genesisState.CoordinatorCount, got.CoordinatorCount)
	require.ElementsMatch(t, genesisState.ValidatorList, got.ValidatorList)
	// this line is used by starport scaffolding # genesis/test/assert
}
