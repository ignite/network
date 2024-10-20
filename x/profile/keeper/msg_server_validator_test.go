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
		desc1       = sample.ValidatorDescription(addr1)
		desc2       = sample.ValidatorDescription(addr2)
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
				Address:         addr1,
				Identity:        desc1.Identity,
				Moniker:         desc1.Moniker,
				Website:         desc1.Website,
				SecurityContact: desc1.SecurityContact,
				Details:         desc1.Details,
			},
		}, {
			name: "should allow updating an existing validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:         addr1,
				Identity:        desc2.Identity,
				Moniker:         desc2.Moniker,
				Website:         desc2.Website,
				SecurityContact: desc2.SecurityContact,
				Details:         desc2.Details,
			},
		}, {
			name: "should allow creating a second validator",
			msg: types.MsgUpdateValidatorDescription{
				Address:         addr2,
				Identity:        desc2.Identity,
				Moniker:         desc2.Moniker,
				Website:         desc2.Website,
				SecurityContact: desc2.SecurityContact,
				Details:         desc2.Details,
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

			if len(tt.msg.Identity) > 0 {
				require.EqualValues(t, tt.msg.Identity, validator.Description.Identity)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Identity, oldValidator.Description.Identity)
			}

			if len(tt.msg.Website) > 0 {
				require.EqualValues(t, tt.msg.Website, validator.Description.Website)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Website, oldValidator.Description.Website)
			}

			if len(tt.msg.Details) > 0 {
				require.EqualValues(t, tt.msg.Details, validator.Description.Details)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Details, oldValidator.Description.Details)
			}

			if len(tt.msg.Moniker) > 0 {
				require.EqualValues(t, tt.msg.Moniker, validator.Description.Moniker)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.Moniker, oldValidator.Description.Moniker)
			}

			if len(tt.msg.SecurityContact) > 0 {
				require.EqualValues(t, tt.msg.SecurityContact, validator.Description.SecurityContact)
			} else if oldErr != nil {
				require.EqualValues(t, oldValidator.Description.SecurityContact, oldValidator.Description.SecurityContact)
			}
		})
	}
}
