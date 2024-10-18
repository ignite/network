package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/types"
)

func TestMsgAddValidatorOperatorAddress(t *testing.T) {
	var (
		ctx, tk, ts = testkeeper.NewTestSetup(t)
		valAddr     = sample.Address(r)
		opAddr      = sample.Address(r)
	)

	err := tk.ProfileKeeper.Validator.Set(ctx, valAddr, types.Validator{
		Address:           valAddr,
		Description:       types.ValidatorDescription{},
		OperatorAddresses: []string{opAddr},
	})
	require.NoError(t, err)

	tests := []struct {
		name   string
		msg    *types.MsgAddValidatorOperatorAddress
		newVal bool
		err    error
	}{
		{
			name: "should allow associating a new operator address to a validator",
			msg: &types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: valAddr,
				OperatorAddress:  sample.Address(r),
			},
		},
		{
			name: "should allow creating a new validator if it doesn't exist",
			msg: &types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: sample.Address(r),
				OperatorAddress:  sample.Address(r),
			},
			newVal: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ts.ProfileSrv.AddValidatorOperatorAddress(ctx, tt.msg)
			require.NoError(t, err)

			validator, err := tk.ProfileKeeper.GetValidator(ctx, tt.msg.ValidatorAddress)
			require.NoError(t, err, "validator was not saved")
			require.Equal(t, tt.msg.ValidatorAddress, validator.Address)
			require.True(t, validator.HasOperatorAddress(tt.msg.OperatorAddress))

			// check that original address still exists if we appended to existing validator
			if !tt.newVal {
				require.True(t, validator.HasOperatorAddress(opAddr))
			}

			valByOpAddr, err := tk.ProfileKeeper.GetValidatorByOperatorAddress(ctx, tt.msg.OperatorAddress)
			require.NoError(t, err, "validator by operator address was not saved")
			require.Equal(t, tt.msg.ValidatorAddress, valByOpAddr.ValidatorAddress)
			require.Equal(t, tt.msg.OperatorAddress, valByOpAddr.OperatorAddress)
		})
	}
}

func TestMsgUpdateValidatorDescription(t *testing.T) {
	var (
		addr1       = sample.Address(r)
		addr2       = sample.Address(r)
		ctx, tk, ts = testkeeper.NewTestSetup(t)
	)
	tests := []struct {
		name string
		msg  types.MsgUpdateValidatorDescription
		err  error
	}{
		{
			name: "should allow creating a new validator if doesn't exist",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr1,
				Description: sample.ValidatorDescription(addr1),
			},
		}, {
			name: "should allow updating an existing validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr1,
				Description: sample.ValidatorDescription(addr2),
			},
		}, {
			name: "should allow creating a second validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:     addr2,
				Description: sample.ValidatorDescription(addr2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldValidator, oldErr := tk.ProfileKeeper.GetValidator(ctx, tt.msg.Address)

			_, err := ts.ProfileSrv.UpdateValidatorDescription(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			validator, err := tk.ProfileKeeper.GetValidator(ctx, tt.msg.Address)
			require.NoError(t, err, "validator not found")
			require.EqualValues(t, tt.msg.Address, validator.Address)

			if len(tt.msg.Description.Identity) > 0 {
				require.EqualValues(t, tt.msg.Description.Identity, validator.Description.Identity)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Identity, oldValidator.Description.Identity)
			}

			if len(tt.msg.Description.Website) > 0 {
				require.EqualValues(t, tt.msg.Description.Website, validator.Description.Website)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Website, oldValidator.Description.Website)
			}

			if len(tt.msg.Description.Details) > 0 {
				require.EqualValues(t, tt.msg.Description.Details, validator.Description.Details)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Details, oldValidator.Description.Details)
			}

			if len(tt.msg.Description.Moniker) > 0 {
				require.EqualValues(t, tt.msg.Description.Moniker, validator.Description.Moniker)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Moniker, oldValidator.Description.Moniker)
			}

			if len(tt.msg.Description.SecurityContact) > 0 {
				require.EqualValues(t, tt.msg.Description.SecurityContact, validator.Description.SecurityContact)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.SecurityContact, oldValidator.Description.SecurityContact)
			}
		})
	}
}
