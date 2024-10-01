package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorDescription
		err  error
	}{
		{
			name: "should prevent validate invalid coordinator address",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: "invalid address",
			},
			err: types.ErrInvalidCoordinatorAddress,
		},
		{
			name: "should prevent validate empty description",
			msg: types.MsgUpdateCoordinatorDescription{
				Address:     addr,
				Description: types.CoordinatorDescription{},
			},
			err: types.ErrEmptyDescription,
		},
		{
			name: "should validate valid message",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: sample.Address(r),
				Description: types.CoordinatorDescription{
					Identity: "identity",
					Website:  "website",
					Details:  "details",
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

func TestMsgUpdateCoordinatorAddress_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorAddress
		err  error
	}{
		{
			name: "should prevent validate invalid coordinator address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    "invalid address",
				NewAddress: sample.Address(r),
			},
			err: types.ErrInvalidCoordinatorAddress,
		},
		{
			name: "should prevent validate invalid new address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    sample.Address(r),
				NewAddress: "invalid address",
			},
			err: types.ErrInvalidCoordinatorAddress,
		},
		{
			name: "should prevent validate similar new address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: addr,
			},
			err: types.ErrDupAddress,
		},
		{
			name: "should validate different addresses",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    sample.Address(r),
				NewAddress: sample.Address(r),
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
