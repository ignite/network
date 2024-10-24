package launch_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	launch "github.com/ignite/network/x/launch/module"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
*/

func TestGenesis(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("should allow import and export the genesis state", func(t *testing.T) {
		genesisState := sample.LaunchGenesisState(r)
		err := launch.InitGenesis(ctx, tk.LaunchKeeper, genesisState)
		require.NoError(t, err)
		got, err := launch.ExportGenesis(ctx, tk.LaunchKeeper)
		require.NoError(t, err)

		// Compare lists
		require.ElementsMatch(t, genesisState.ChainList, got.ChainList)
		require.Equal(t, genesisState.ChainCount, got.ChainCount)

		require.ElementsMatch(t, genesisState.GenesisAccountList, got.GenesisAccountList)
		require.ElementsMatch(t, genesisState.VestingAccountList, got.VestingAccountList)
		require.ElementsMatch(t, genesisState.GenesisValidatorList, got.GenesisValidatorList)
		require.ElementsMatch(t, genesisState.RequestList, got.RequestList)
		require.ElementsMatch(t, genesisState.RequestCounters, got.RequestCounters)

		require.Equal(t, genesisState.Params, got.Params)
	})
	// this line is used by starport scaffolding # genesis/test/assert
}
