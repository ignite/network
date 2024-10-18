package monitoringp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	monitoringp "github.com/ignite/network/x/monitoringp/module"
	"github.com/ignite/network/x/monitoringp/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortID: types.PortID,
		MonitoringInfo: &types.MonitoringInfo{
			Transmitted:     false,
			SignatureCounts: networktypes.SignatureCounts{},
		},
		ConnectionChannelID: &types.ConnectionChannelID{
			ChannelID: "98",
		},
		ConsumerClientID: &types.ConsumerClientID{
			ClientID: "42",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.MonitoringpKeeper(t)
	err := monitoringp.InitGenesis(ctx, k, genesisState)
	require.NoError(t, err)
	got, err := monitoringp.ExportGenesis(ctx, k)
	require.NoError(t, err)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortID, got.PortID)

	require.Equal(t, genesisState.MonitoringInfo, got.MonitoringInfo)
	require.Equal(t, genesisState.ConnectionChannelID, got.ConnectionChannelID)
	require.Equal(t, genesisState.ConsumerClientID, got.ConsumerClientID)
	// this line is used by starport scaffolding # genesis/test/assert
}
