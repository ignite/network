package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/profile module sentinel errors
var (
	ErrInvalidSigner                      = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrCoordinatorAlreadyExist            = sdkerrors.Register(ModuleName, 1102, "coordinator address already exist")
	ErrCoordinatorAddressNotFound         = sdkerrors.Register(ModuleName, 1103, "coordinator address not found")
	ErrCoordinatorInvalid                 = sdkerrors.Register(ModuleName, 1104, "invalid coordinator")
	ErrEmptyDescription                   = sdkerrors.Register(ModuleName, 1105, "you must provide at least one description parameter")
	ErrDupAddress                         = sdkerrors.Register(ModuleName, 1106, "address is duplicated")
	ErrCoordinatorInactive                = sdkerrors.Register(ModuleName, 1107, "inactive coordinator")
	ErrInvalidCoordinatorAddress          = sdkerrors.Register(ModuleName, 1108, "invalid coordinator address")
	ErrInvalidValidatorAddress            = sdkerrors.Register(ModuleName, 1109, "invalid validator address")
	ErrInvalidOperatorAddress             = sdkerrors.Register(ModuleName, 1110, "invalid operator address")
	ErrValidatorNotFound                  = sdkerrors.Register(ModuleName, 1111, "validator not found")
	ErrValidatorByOperatorAddressNotFound = sdkerrors.Register(ModuleName, 1112, "validator by operator address not found")
	ErrCoordinatorNotFound                = sdkerrors.Register(ModuleName, 1113, "coordinator not found")
)
