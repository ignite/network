package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/reward/types"
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
				RewardPoolList: []types.RewardPool{
					sample.RewardPool(r, 1),
					sample.RewardPool(r, 2),
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated rewardPool",
			genState: &types.GenesisState{
				RewardPoolList: []types.RewardPool{
					sample.RewardPool(r, 1),
					sample.RewardPool(r, 1),
				},
			},
			valid: false,
		},
		{
			desc: "should prevent invalid rewardPool",
			genState: &types.GenesisState{
				RewardPoolList: []types.RewardPool{
					sample.RewardPool(r, 1),
					{}, // invalid reward pool
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
