package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/network/x/project/types"
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
				MainnetAccountList: []types.MainnetAccount{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 1,
					},
				},
				ProjectList: []types.Project{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 1,
					},
				},
				ProjectCount: 2,
				ProjectChainsList: []types.ProjectChains{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated mainnetAccount",
			genState: &types.GenesisState{
				MainnetAccountList: []types.MainnetAccount{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated project",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid project count",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					{
						ProjectID: 1,
					},
				},
				ProjectCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated projectChains",
			genState: &types.GenesisState{
				ProjectChainsList: []types.ProjectChains{
					{
						ProjectID: 0,
					},
					{
						ProjectID: 0,
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
