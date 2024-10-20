package reward_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	reward "github.com/ignite/network/x/reward/module"
	"github.com/ignite/network/x/reward/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RewardPoolList: []types.RewardPool{
			{
				LaunchId: 0,
			},
			{
				LaunchId: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.RewardKeeper(t)
	err := reward.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := reward.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RewardPoolList, got.RewardPoolList)
	// this line is used by starport scaffolding # genesis/test/assert
}
