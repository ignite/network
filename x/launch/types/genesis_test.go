package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/network/x/launch/types"
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
				ChainList: []types.Chain{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 1,
					},
				},
				ChainCount: 2,
				GenesisAccountList: []types.GenesisAccount{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 1,
					},
				},
				GenesisValidatorList: []types.GenesisValidator{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 1,
					},
				},
				VestingAccountList: []types.VestingAccount{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 1,
					},
				},
				RequestList: []types.Request{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 1,
					},
				},
				RequestCounters: []types.RequestCounter{
					{
						LaunchID: 0,
						Counter:  1,
					},
					{
						LaunchID: 1,
						Counter:  2,
					},
				},
				ParamChangeList: []types.ParamChange{
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
			desc: "duplicated chain",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
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
			desc: "invalid chain count",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
					{
						LaunchID: 1,
					},
				},
				ChainCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated genesisAccount",
			genState: &types.GenesisState{
				GenesisAccountList: []types.GenesisAccount{
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
			desc: "duplicated genesisValidator",
			genState: &types.GenesisState{
				GenesisValidatorList: []types.GenesisValidator{
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
			desc: "duplicated vestingAccount",
			genState: &types.GenesisState{
				VestingAccountList: []types.VestingAccount{
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
			desc: "duplicated request",
			genState: &types.GenesisState{
				RequestList: []types.Request{
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
			desc: "invalid request count",
			genState: &types.GenesisState{
				RequestList: []types.Request{
					{
						LaunchID: 1,
					},
				},
				RequestCounters: []types.RequestCounter{
					{
						LaunchID: 0,
						Counter:  1,
					},
					{
						LaunchID: 1,
						Counter:  2,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated paramChange",
			genState: &types.GenesisState{
				ParamChangeList: []types.ParamChange{
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
