package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/monitoringp/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "should allow valid default genesis",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "should allow valid genesis state",
			genState: &types.GenesisState{
				PortID: types.PortID,
				ConsumerClientID: &types.ConsumerClientID{
					ClientID: "29",
				},
				Params: types.DefaultParams(),
				ConnectionChannelID: &types.ConnectionChannelID{
					ChannelID: "67",
				},
				MonitoringInfo: &types.MonitoringInfo{},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "should prevent invalid params",
			genState: &types.GenesisState{
				PortID: types.PortID,
				ConsumerClientID: &types.ConsumerClientID{
					ClientID: "29",
				},
				Params: types.NewParams(
					1000,
					"foo", // chain id should be <chain-name>-<revision-number>
					sample.ConsensusState(0),
					networktypes.DefaultUnbondingPeriod,
					1,
				),
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			desc: "should prevent invalid monitoring info",
			genState: &types.GenesisState{
				PortID: types.PortID,
				ConsumerClientID: &types.ConsumerClientID{
					ClientID: "29",
				},
				Params: types.DefaultParams(),
				ConnectionChannelID: &types.ConnectionChannelID{
					ChannelID: "67",
				},
				// Block count is lower than sum of relative signatures
				MonitoringInfo: &types.MonitoringInfo{
					SignatureCounts: networktypes.SignatureCounts{
						BlockCount: 1,
						Counts: []networktypes.SignatureCount{
							{
								OpAddress:          sample.Address(r),
								RelativeSignatures: sdkmath.LegacyNewDec(10),
							},
						},
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
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
