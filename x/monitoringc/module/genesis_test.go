package monitoringc_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	monitoringc "github.com/ignite/network/x/monitoringc/module"
	"github.com/ignite/network/x/monitoringc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortID: types.PortID,
		LaunchIDFromChannelIDList: []types.LaunchIDFromChannelID{
			{
				ChannelID: "0",
			},
			{
				ChannelID: "1",
			},
		},
		LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
			{
				ClientID: "0",
			},
			{
				ClientID: "1",
			},
		},
		MonitoringHistoryList: []types.MonitoringHistory{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		VerifiedClientIDList: []types.VerifiedClientID{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		ProviderClientIDList: []types.ProviderClientID{
			{
				LaunchID: 0,
			},
			{
				LaunchID: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.MonitoringcKeeper(t)
	err := monitoringc.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := monitoringc.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortID, got.PortID)

	require.ElementsMatch(t, genesisState.LaunchIDFromChannelIDList, got.LaunchIDFromChannelIDList)
	require.ElementsMatch(t, genesisState.LaunchIDFromVerifiedClientIDList, got.LaunchIDFromVerifiedClientIDList)
	require.ElementsMatch(t, genesisState.MonitoringHistoryList, got.MonitoringHistoryList)
	require.ElementsMatch(t, genesisState.VerifiedClientIDList, got.VerifiedClientIDList)
	require.ElementsMatch(t, genesisState.ProviderClientIDList, got.ProviderClientIDList)
	// this line is used by starport scaffolding # genesis/test/assert
}
