package monitoringc_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	monitoringc "github.com/ignite/network/x/monitoringc/module"
	"github.com/ignite/network/x/monitoringc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		VerifiedClientIdList: []types.VerifiedClientID{
			{
				LaunchId:     0,
				ClientIdList: []string{"0"},
			},
			{
				LaunchId:     1,
				ClientIdList: []string{"0"},
			},
		},
		ProviderClientIdList: []types.ProviderClientID{
			{
				LaunchId: 0,
			},
			{
				LaunchId: 1,
			},
		},
		LaunchIdFromVerifiedClientIdList: []types.LaunchIDFromVerifiedClientID{
			{
				ClientId: "0",
			},
			{
				ClientId: "1",
			},
		},
		LaunchIdFromChannelIdList: []types.LaunchIDFromChannelID{
			{
				ChannelId: "0",
			},
			{
				ChannelId: "1",
			},
		},
		MonitoringHistoryList: []types.MonitoringHistory{
			{
				LaunchId: 0,
			},
			{
				LaunchId: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should allow import and export of genesis", func(t *testing.T) {
		err := monitoringc.InitGenesis(ctx, tk.MonitoringConsumerKeeper, genesisState)
		require.NoError(t, err)
		got, err := monitoringc.ExportGenesis(ctx, tk.MonitoringConsumerKeeper)
		require.NoError(t, err)
		require.NotNil(t, got)

		nullify.Fill(&genesisState)
		nullify.Fill(got)

		require.Equal(t, genesisState.PortId, got.PortId)

		require.ElementsMatch(t, genesisState.VerifiedClientIdList, got.VerifiedClientIdList)
		require.ElementsMatch(t, genesisState.ProviderClientIdList, got.ProviderClientIdList)
		require.ElementsMatch(t, genesisState.LaunchIdFromVerifiedClientIdList, got.LaunchIdFromVerifiedClientIdList)
		require.ElementsMatch(t, genesisState.LaunchIdFromChannelIdList, got.LaunchIdFromChannelIdList)
		require.ElementsMatch(t, genesisState.MonitoringHistoryList, got.MonitoringHistoryList)
		// this line is used by starport scaffolding # genesis/test/assert
	})
}
