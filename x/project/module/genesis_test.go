package project_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	project "github.com/ignite/network/x/project/module"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
*/

func TestGenesis(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("should allow importing and exporting genesis", func(t *testing.T) {
		genesisState := sample.ProjectGenesisStateWithAccounts(r)

		err := project.InitGenesis(ctx, tk.ProjectKeeper, genesisState)
		require.NoError(t, err)
		got, err := project.ExportGenesis(ctx, tk.ProjectKeeper)
		require.NoError(t, err)

		require.ElementsMatch(t, genesisState.ProjectChainsList, got.ProjectChainsList)
		require.ElementsMatch(t, genesisState.ProjectList, got.ProjectList)
		require.Equal(t, genesisState.ProjectCount, got.ProjectCount)
		require.ElementsMatch(t, genesisState.MainnetAccountList, got.MainnetAccountList)
		require.Equal(t, genesisState.Params, got.Params)
		maxShares, err := tk.ProjectKeeper.TotalShares.Get(ctx)
		require.NoError(t, err)
		require.Equal(t, uint64(networktypes.TotalShareNumber), maxShares)
	})

	// this line is used by starport scaffolding # genesis/test/assert
}
