package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/network/x/monitoringc/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated launchIDFromChannelID",
			genState: &types.GenesisState{
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
			desc: "duplicated launchIDFromVerifiedClientID",
			genState: &types.GenesisState{
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{
						ClientID: "0",
					},
					{
						ClientID: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated monitoringHistory",
			genState: &types.GenesisState{
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
		{
			desc: "duplicated verifiedClientID",
			genState: &types.GenesisState{
				VerifiedClientIDList: []types.VerifiedClientID{
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
		{
			desc: "duplicated providerClientID",
			genState: &types.GenesisState{
				ProviderClientIDList: []types.ProviderClientID{
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
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
