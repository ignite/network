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
		Address: address,
		Description: ValidatorDescription{
			Identity:        identity,
			Moniker:         moniker,
			Website:         website,
			SecurityContact: securityContact,
			Details:         details,
		},
	}
}

// Type returns the msg type.
func (msg *MsgUpdateValidatorDescription) Type() string {
	return sdk.MsgTypeURL(msg)
}

// TODO fixme
func (msg *MsgAddValidatorOperatorAddress) GetSigners() []sdk.AccAddress {
	validatorAddress, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	operatorAddress, err := sdk.AccAddressFromBech32(msg.OperatorAddress)
	if err != nil {
		panic(err)
	}

	// validator must prove ownership of both address
	return []sdk.AccAddress{validatorAddress, operatorAddress}
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
	if msg.Description.Details == "" &&
		msg.Description.Moniker == "" &&
		msg.Description.Identity == "" &&
		msg.Description.Website == "" &&
		msg.Description.SecurityContact == "" {
		return sdkerrors.Wrap(ErrEmptyDescription, msg.Address)
	}
	return nil
}
