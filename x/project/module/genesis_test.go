package project_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	project "github.com/ignite/network/x/project/module"
	"github.com/ignite/network/x/project/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MainnetAccountList: []types.MainnetAccount{
			{
				ProjectID: 0,
			},
			{
				ProjectID: 1,
			},
		},
		ProjectList: []types.Project{
			{
				ProjectID: 0,
			},
			{
				ProjectID: 1,
			},
		},
		ProjectCount: 2,
		ProjectChainsList: []types.ProjectChains{
			{
				ProjectID: 0,
			},
			{
				ProjectID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.ProjectKeeper(t)
	err := project.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := project.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MainnetAccountList, got.MainnetAccountList)
	require.ElementsMatch(t, genesisState.ProjectList, got.ProjectList)
	require.Equal(t, genesisState.ProjectCount, got.ProjectCount)
	require.ElementsMatch(t, genesisState.ProjectChainsList, got.ProjectChainsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
