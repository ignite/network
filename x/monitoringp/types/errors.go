package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/monitoringp module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1102, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1103, "invalid version")
)
