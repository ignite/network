package keeper

import (
	"context"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/profile/types"
)

func (k msgServer) UpdateValidatorDescription(ctx context.Context, msg *types.MsgUpdateValidatorDescription) (*types.MsgUpdateValidatorDescriptionResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if _, err := k.addressCodec.StringToBytes(msg.Address); err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid address %s", err.Error())
	}

	// Check if the validator address is already in the store
	validator, err := k.Validator.Get(ctx, msg.Address)
	if errors.IsOf(err, collections.ErrNotFound) {
		validator = types.Validator{
			Address:     msg.Address,
			Description: types.ValidatorDescription{},
		}
	} else if err != nil {
		return nil, err
	}

	if len(msg.Identity) > 0 {
		validator.Description.Identity = msg.Identity
	}
	if len(msg.Website) > 0 {
		validator.Description.Website = msg.Website
	}
	if len(msg.Details) > 0 {
		validator.Description.Details = msg.Details
	}
	if len(msg.Moniker) > 0 {
		validator.Description.Moniker = msg.Moniker
	}
	if len(msg.SecurityContact) > 0 {
		validator.Description.SecurityContact = msg.SecurityContact
	}

	if errors.IsOf(err, collections.ErrNotFound) {
		err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
			&types.EventValidatorCreated{
				Address:           validator.Address,
				OperatorAddresses: validator.OperatorAddresses,
			})
	}

	if err := k.Validator.Set(ctx, validator.Address, validator); err != nil {
		return nil, err
	}

	return &types.MsgUpdateValidatorDescriptionResponse{}, nil
}

func (k msgServer) AddValidatorOperatorAddress(ctx context.Context, msg *types.MsgAddValidatorOperatorAddress) (*types.MsgAddValidatorOperatorAddressResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if _, err := k.addressCodec.StringToBytes(msg.ValidatorAddress); err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid validator address %s", err.Error())
	}

	valAddr := msg.ValidatorAddress
	opAddr := msg.OperatorAddress

	validator := types.Validator{
		Address:           valAddr,
		OperatorAddresses: []string{opAddr},
		Description:       types.ValidatorDescription{},
	}

	// get the current validator to eventually overwrite description and add opAddr
	validatorStore, err := k.Validator.Get(ctx, valAddr)
	if !errors.IsOf(err, collections.ErrNotFound) {
		validator.Description = validatorStore.Description
		validator = validatorStore.AddValidatorOperatorAddress(opAddr)
	}

	// store validator information
	if err := k.Validator.Set(ctx, validator.Address, validator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set validator")
	}
	if err := k.ValidatorByOperatorAddress.Set(ctx, opAddr, types.ValidatorByOperatorAddress{
		OperatorAddress:  opAddr,
		ValidatorAddress: valAddr,
	}); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set validator by operator")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if !errors.IsOf(err, collections.ErrNotFound) {
		err = sdkCtx.EventManager().EmitTypedEvent(
			&types.EventValidatorOperatorAddressesUpdated{
				Address:           validator.Address,
				OperatorAddresses: validator.OperatorAddresses,
			})
	} else {
		err = sdkCtx.EventManager().EmitTypedEvent(
			&types.EventValidatorCreated{
				Address:           validator.Address,
				OperatorAddresses: validator.OperatorAddresses,
			})
	}

	return &types.MsgAddValidatorOperatorAddressResponse{}, nil
}
