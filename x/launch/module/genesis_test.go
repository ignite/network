package launch_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	launch "github.com/ignite/network/x/launch/module"
	"github.com/ignite/network/x/launch/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ChainList: []types.Chain{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		ChainCount: 2,
		GenesisAccountList: []types.GenesisAccount{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		GenesisValidatorList: []types.GenesisValidator{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		VestingAccountList: []types.VestingAccount{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		RequestList: []types.Request{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		RequestCounters: []types.RequestCounter{
			{
				LaunchID: 0,
				Counter:  1,
			},
			{
				LaunchID: 1,
				Counter:  2,
			},
		},
		ParamChangeList: []types.ParamChange{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.LaunchKeeper(t)
	err := launch.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := launch.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ChainList, got.ChainList)
	require.Equal(t, genesisState.ChainCount, got.ChainCount)
	require.ElementsMatch(t, genesisState.GenesisAccountList, got.GenesisAccountList)
	require.ElementsMatch(t, genesisState.GenesisValidatorList, got.GenesisValidatorList)
	require.ElementsMatch(t, genesisState.VestingAccountList, got.VestingAccountList)
	require.ElementsMatch(t, genesisState.RequestList, got.RequestList)
	require.ElementsMatch(t, genesisState.RequestCounters, got.RequestCounters)
	require.ElementsMatch(t, genesisState.ParamChangeList, got.ParamChangeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
