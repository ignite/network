package participation_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	participation "github.com/ignite/network/x/participation/module"
	"github.com/ignite/network/x/participation/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AuctionUsedAllocationsList: []types.AuctionUsedAllocations{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		UsedAllocationsList: []types.UsedAllocations{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.ParticipationKeeper(t)
	err := participation.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := participation.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AuctionUsedAllocationsList, got.AuctionUsedAllocationsList)
	require.ElementsMatch(t, genesisState.UsedAllocationsList, got.UsedAllocationsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
