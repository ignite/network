package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/monitoringc/types"
)

func TestMsgCreateClient_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateClient
		err  error
	}{
		{
			name: "should validate valid message",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(0),
				UnbondingPeriod: networktypes.DefaultUnbondingPeriod,
				RevisionHeight:  networktypes.DefaultRevisionHeight,
			},
		},
		{
			name: "should prevent invalid consensus state",
			msg: types.MsgCreateClient{
				Creator:  sample.Address(r),
				LaunchID: 0,
				ConsensusState: networktypes.NewConsensusState(
					"2022-01-12T07:56:35.394367Z",
					"foo",
					"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
				),
				ValidatorSet:    sample.ValidatorSet(1),
				UnbondingPeriod: networktypes.DefaultUnbondingPeriod,
				RevisionHeight:  networktypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidConsensusState,
		},
		{
			name: "should prevent invalid validator set",
			msg: types.MsgCreateClient{
				Creator:        sample.Address(r),
				LaunchID:       0,
				ConsensusState: sample.ConsensusState(0),
				ValidatorSet: networktypes.NewValidatorSet(
					networktypes.NewValidator(
						"foo",
						0,
						100,
					),
				),
				UnbondingPeriod: networktypes.DefaultUnbondingPeriod,
				RevisionHeight:  networktypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidValidatorSet,
		},
		{
			name: "should prevent validator set not matching consensus state",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(1),
				UnbondingPeriod: networktypes.DefaultUnbondingPeriod,
				RevisionHeight:  networktypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidValidatorSetHash,
		},
		{
			name: "should prevent unbonding period lower than minimal",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(0),
				UnbondingPeriod: networktypes.MinimalUnbondingPeriod - 1,
				RevisionHeight:  networktypes.DefaultRevisionHeight,
			},
			err: types.ErrInvalidUnbondingPeriod,
		},
		{
			name: "should prevent zero revision height",
			msg: types.MsgCreateClient{
				Creator:         sample.Address(r),
				LaunchID:        0,
				ConsensusState:  sample.ConsensusState(0),
				ValidatorSet:    sample.ValidatorSet(0),
				UnbondingPeriod: networktypes.MinimalUnbondingPeriod,
				RevisionHeight:  0,
			},
			err: types.ErrInvalidRevisionHeight,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
