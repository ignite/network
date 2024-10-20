package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgUpdateValidatorDescription(
	address,
	identity,
	moniker,
	website,
	securityContact,
	details string,
) *MsgUpdateValidatorDescription {
	return &MsgUpdateValidatorDescription{
		Address:         address,
		Identity:        identity,
		Moniker:         moniker,
		Website:         website,
		SecurityContact: securityContact,
		Details:         details,
	}
}

// Type returns the msg type.
func (msg *MsgUpdateValidatorDescription) Type() string {
	return sdk.MsgTypeURL(msg)
}

func (msg *MsgAddValidatorOperatorAddress) ValidateBasic() error {
	if msg.ValidatorAddress == msg.OperatorAddress {
		return sdkerrors.Wrapf(ErrDupAddress, "validator profile address and operator address must be different")
	}
	return nil
}

func NewMsgAddValidatorOperatorAddress(validatorAddress string, operatorAddress string) *MsgAddValidatorOperatorAddress {
	return &MsgAddValidatorOperatorAddress{
		ValidatorAddress: validatorAddress,
		OperatorAddress:  operatorAddress,
	}
}

// Type returns the msg type.
func (msg *MsgAddValidatorOperatorAddress) Type() string {
	return sdk.MsgTypeURL(msg)
}

func (msg *MsgUpdateValidatorDescription) ValidateBasic() error {
	if msg.Details == "" &&
		msg.Moniker == "" &&
		msg.Identity == "" &&
		msg.Website == "" &&
		msg.SecurityContact == "" {
		return sdkerrors.Wrap(ErrEmptyDescription, msg.Address)
	}
	return nil
}
