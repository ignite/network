package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/monitoringc module sentinel errors
var (
	ErrInvalidSigner                = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrInvalidVersion               = sdkerrors.Register(ModuleName, 1102, "invalid version")
	ErrInvalidClientState           = sdkerrors.Register(ModuleName, 1103, "invalid client state")
	ErrInvalidConsensusState        = sdkerrors.Register(ModuleName, 1104, "invalid consensus state")
	ErrInvalidValidatorSet          = sdkerrors.Register(ModuleName, 1105, "invalid validator set")
	ErrInvalidValidatorSetHash      = sdkerrors.Register(ModuleName, 1106, "invalid validator set hash")
	ErrClientCreationFailure        = sdkerrors.Register(ModuleName, 1107, "failed to create IBC client")
	ErrInvalidHandshake             = sdkerrors.Register(ModuleName, 1108, "invalid handshake")
	ErrClientNotVerified            = sdkerrors.Register(ModuleName, 1109, "ibc client not verified")
	ErrConnectionAlreadyEstablished = sdkerrors.Register(ModuleName, 1110, "ibc connection already established")
	ErrInvalidUnbondingPeriod       = sdkerrors.Register(ModuleName, 1111, "invalid unbonding period")
	ErrInvalidRevisionHeight        = sdkerrors.Register(ModuleName, 1112, "invalid revision height")
	ErrVerifiedClientIdListNotFound = sdkerrors.Register(ModuleName, 1113, "verified client IDs not found")
	ErrInvalidClientCreatorAddress  = sdkerrors.Register(ModuleName, 1114, "invalid client creator address")
	ErrCannotCloseChannel           = sdkerrors.Register(ModuleName, 1115, "user cannot close channel")
	ErrJSONUnmarshal                = sdkerrors.Register(ModuleName, 1116, "failed to unmarshal JSON")
	ErrJSONMarshal                  = sdkerrors.Register(ModuleName, 1117, "failed to marshal JSON")
	ErrUnrecognizedPacketType       = sdkerrors.Register(ModuleName, 1118, "unrecognized packet type")
	ErrInvalidPacketTimeout         = sdkerrors.Register(ModuleName, 1119, "invalid packet timeout")
)
