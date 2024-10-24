package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/network/x/monitoringc/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			name:     "should allow valid default genesis",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			name: "should allow valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIdList: []types.VerifiedClientID{
					{LaunchId: 0, ClientIdList: []string{"0"}},
					{LaunchId: 1, ClientIdList: []string{"1", "2"}},
				},
				ProviderClientIdList: []types.ProviderClientID{
					{LaunchId: 0, ClientId: "0"},
					{LaunchId: 1, ClientId: "2"},
				},
				LaunchIdFromVerifiedClientIdList: []types.LaunchIDFromVerifiedClientID{
					{LaunchId: 0, ClientId: "0"},
					{LaunchId: 1, ClientId: "1"},
				},
				LaunchIdFromChannelIdList: []types.LaunchIDFromChannelID{
					{LaunchId: 0, ChannelId: "0"},
					{LaunchId: 1, ChannelId: "1"},
				},
				MonitoringHistoryList: []types.MonitoringHistory{
					{
						LaunchId: 0,
					},
					{
						LaunchId: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			name: "should prevent invalid portID",
			genState: &types.GenesisState{
				PortId: "",
			},
			valid: false,
		},
		{
			name: "should prevent duplicated verifiedClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIdList: []types.VerifiedClientID{
					{
						LaunchId:     0,
						ClientIdList: []string{"0"},
					},
					{
						LaunchId:     0,
						ClientIdList: []string{"1", "2"},
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated clientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIdList: []types.VerifiedClientID{
					{
						LaunchId:     0,
						ClientIdList: []string{"0", "0"},
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated providerClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIdList: []types.VerifiedClientID{
					{
						LaunchId:     0,
						ClientIdList: []string{"0"},
					},
				},
				ProviderClientIdList: []types.ProviderClientID{
					{
						LaunchId: 0,
						ClientId: "0",
					},
					{
						LaunchId: 0,
						ClientId: "0",
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated launchIDFromVerifiedClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIdList: []types.VerifiedClientID{
					{
						LaunchId:     0,
						ClientIdList: []string{"0"},
					},
				},
				LaunchIdFromVerifiedClientIdList: []types.LaunchIDFromVerifiedClientID{
					{
						ClientId: "0",
						LaunchId: 0,
					},
					{
						ClientId: "0",
						LaunchId: 0,
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent provider client id without valid client id",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIdList: []types.VerifiedClientID{
					{LaunchId: 0, ClientIdList: []string{"0"}},
					{LaunchId: 1, ClientIdList: []string{"1", "2"}},
				},
				ProviderClientIdList: []types.ProviderClientID{
					{LaunchId: 0, ClientId: "0"},
					{LaunchId: 1, ClientId: "3"},
				},
				LaunchIdFromVerifiedClientIdList: []types.LaunchIDFromVerifiedClientID{
					{LaunchId: 0, ClientId: "0"},
					{LaunchId: 1, ClientId: "2"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			name: "should prevent launch id from verified client id without valid client id",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIdList: []types.VerifiedClientID{
					{LaunchId: 0, ClientIdList: []string{"0"}},
					{LaunchId: 1, ClientIdList: []string{"1", "2"}},
				},
				ProviderClientIdList: []types.ProviderClientID{
					{LaunchId: 0, ClientId: "0"},
					{LaunchId: 1, ClientId: "2"},
				},
				LaunchIdFromVerifiedClientIdList: []types.LaunchIDFromVerifiedClientID{
					{LaunchId: 0, ClientId: "1"},
					{LaunchId: 1, ClientId: "1"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			name: "should prevent duplicated launchIDFromChannelID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				LaunchIdFromChannelIdList: []types.LaunchIDFromChannelID{
					{
						ChannelId: "0",
					},
					{
						ChannelId: "0",
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated monitoringHistory",
			genState: &types.GenesisState{
				PortId: types.PortID,
				MonitoringHistoryList: []types.MonitoringHistory{
					{
						LaunchId: 0,
					},
					{
						LaunchId: 0,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
