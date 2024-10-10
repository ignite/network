package participation_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	participation "github.com/ignite/network/x/participation/module"
)

/*
// We use a genesis template from sample package, therefore this placeholder is not used
// this line is used by starport scaffolding # genesis/test/state
*/

func TestGenesis(t *testing.T) {
	r := sample.Rand()
	genesisState := sample.ParticipationGenesisStateWithAllocations(r)

	k, ctx, _ := keepertest.ParticipationKeeper(t)
	err := participation.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := participation.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Params, got.Params)
	require.ElementsMatch(t, genesisState.AuctionUsedAllocationsList, got.AuctionUsedAllocationsList)
	require.ElementsMatch(t, genesisState.UsedAllocationsList, got.UsedAllocationsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
