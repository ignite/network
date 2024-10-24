package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgCreateCoordinator(address, identity, website, details string) *MsgCreateCoordinator {
	return &MsgCreateCoordinator{
		Address:  address,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

// Type returns the msg type.
func (msg *MsgCreateCoordinator) Type() string {
	return sdk.MsgTypeURL(msg)
}

func NewMsgUpdateCoordinatorDescription(address, identity, website, details string) *MsgUpdateCoordinatorDescription {
	return &MsgUpdateCoordinatorDescription{
		Address:  address,
		Identity: identity,
		Website:  website,
		Details:  details,
	}
}

// Type returns the msg type.
func (msg *MsgUpdateCoordinatorDescription) Type() string {
	return sdk.MsgTypeURL(msg)
}

func (msg *MsgUpdateCoordinatorDescription) ValidateBasic() error {
	if msg.Details == "" &&
		msg.Identity == "" &&
		msg.Website == "" {
		return sdkerrors.Wrap(ErrEmptyDescription, msg.Address)
	}
	return nil
}

func NewMsgUpdateCoordinatorAddress(address string, newAddress string) *MsgUpdateCoordinatorAddress {
	return &MsgUpdateCoordinatorAddress{
		Address:    address,
		NewAddress: newAddress,
	}
}

// Type returns the msg type.
func (msg *MsgUpdateCoordinatorAddress) Type() string {
	return sdk.MsgTypeURL(msg)
}

func (msg *MsgUpdateCoordinatorAddress) ValidateBasic() error {
	if msg.Address == msg.NewAddress {
		return sdkerrors.Wrapf(ErrDupAddress,
			"address is equal to new address (%s)", msg.Address)
	}
	return nil
}

func NewMsgDisableCoordinator(address string) *MsgDisableCoordinator {
	return &MsgDisableCoordinator{
		Address: address,
	}
}

// Type returns the msg type.
func (msg *MsgDisableCoordinator) Type() string {
	return sdk.MsgTypeURL(msg)
}
