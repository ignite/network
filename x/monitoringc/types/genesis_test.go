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
				PortID: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "1"},
				},
				LaunchIDFromChannelIDList: []types.LaunchIDFromChannelID{
					{LaunchID: 0, ChannelID: "0"},
					{LaunchID: 1, ChannelID: "1"},
				},
				MonitoringHistoryList: []types.MonitoringHistory{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			name: "should prevent invalid portID",
			genState: &types.GenesisState{
				PortID: "",
			},
			valid: false,
		},
		{
			name: "should prevent duplicated verifiedClientID",
			genState: &types.GenesisState{
				PortID: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
					{
						LaunchID:  0,
						ClientIDs: []string{"1", "2"},
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated clientID",
			genState: &types.GenesisState{
				PortID: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0", "0"},
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated providerClientID",
			genState: &types.GenesisState{
				PortID: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{
						LaunchID: 0,
						ClientID: "0",
					},
					{
						LaunchID: 0,
						ClientID: "0",
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated launchIDFromVerifiedClientID",
			genState: &types.GenesisState{
				PortID: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{
						ClientID: "0",
						LaunchID: 0,
					},
					{
						ClientID: "0",
						LaunchID: 0,
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent provider client id without valid client id",
			genState: &types.GenesisState{
				PortID: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "3"},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			name: "should prevent launch id from verified client id without valid client id",
			genState: &types.GenesisState{
				PortID: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "1"},
					{LaunchID: 1, ClientID: "1"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			name: "should prevent duplicated launchIDFromChannelID",
			genState: &types.GenesisState{
				PortID: types.PortID,
				LaunchIDFromChannelIDList: []types.LaunchIDFromChannelID{
					{
						ChannelID: "0",
					},
					{
						ChannelID: "0",
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated monitoringHistory",
			genState: &types.GenesisState{
				PortID: types.PortID,
				MonitoringHistoryList: []types.MonitoringHistory{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 0,
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
