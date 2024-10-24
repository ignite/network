package profile_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	profile "github.com/ignite/network/x/profile/module"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
*/

func TestGenesis(t *testing.T) {
	r := sample.Rand()
	genesisState := sample.ProfileGenesisState(r)
	k, ctx, _ := keepertest.ProfileKeeper(t)
	err := profile.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := profile.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.Params, got.Params)
	require.ElementsMatch(t, genesisState.CoordinatorList, got.CoordinatorList)
	require.Equal(t, genesisState.CoordinatorCount, got.CoordinatorCount)
	require.ElementsMatch(t, genesisState.ValidatorList, got.ValidatorList)
	// this line is used by starport scaffolding # genesis/test/assert
}
