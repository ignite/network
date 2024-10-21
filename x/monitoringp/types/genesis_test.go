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
	for _, tc := range []struct {
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
				PortId: types.PortID,
				ConsumerClientId: &types.ConsumerClientID{
					ClientId: "29",
				},
				Params: types.DefaultParams(),
				ConnectionChannelId: &types.ConnectionChannelID{
					ChannelId: "67",
				},
				MonitoringInfo: &types.MonitoringInfo{},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "should prevent invalid params",
			genState: &types.GenesisState{
				PortId: types.PortID,
				ConsumerClientId: &types.ConsumerClientID{
					ClientId: "29",
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
				PortId: types.PortID,
				ConsumerClientId: &types.ConsumerClientID{
					ClientId: "29",
				},
				Params: types.DefaultParams(),
				ConnectionChannelId: &types.ConnectionChannelID{
					ChannelId: "67",
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
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
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
