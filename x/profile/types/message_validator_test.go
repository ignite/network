package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/types"
)

func TestMsgAddValidatorOperatorAddress_ValidateBasic(t *testing.T) {
	sampleAddr := sample.Address(r)

	tests := []struct {
		name string
		msg  types.MsgAddValidatorOperatorAddress
		err  error
	}{
		{
			name: "should validate different addresses for Network validator and operator address",
			msg: types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: sample.Address(r),
				OperatorAddress:  sample.Address(r),
			},
		},
		{
			name: "should prevent validate same address for Network validator and operator address",
			msg: types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: sampleAddr,
				OperatorAddress:  sampleAddr,
			},
			err: types.ErrDupAddress,
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

func TestMsgUpdateValidatorDescription_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  types.MsgUpdateValidatorDescription
		err  error
	}{
		{
			name: "should prevent validate emtpy description",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr,
				Description: types.ValidatorDescription{},
			},
			err: types.ErrEmptyDescription,
		}, {
			name: "should validate valid message",
			msg: types.MsgUpdateValidatorDescription{
				Address: sample.Address(r),
				Description: types.ValidatorDescription{
					Identity:        "identity",
					Moniker:         "moniker",
					Website:         "website",
					SecurityContact: "security-contact",
					Details:         "details",
				},
			},
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
