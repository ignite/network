package types

import sdkerrors "cosmossdk.io/errors"

// DONTCOVER

// x/monitoringp module sentinel errors
var (
	ErrInvalidSigner                 = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrInvalidPacketTimeout          = sdkerrors.Register(ModuleName, 1102, "invalid packet timeout")
	ErrInvalidVersion                = sdkerrors.Register(ModuleName, 1103, "invalid version")
	ErrInvalidClientState            = sdkerrors.Register(ModuleName, 1104, "invalid client state")
	ErrInvalidConsensusState         = sdkerrors.Register(ModuleName, 1105, "invalid consensus state")
	ErrClientCreationFailure         = sdkerrors.Register(ModuleName, 1106, "failed to create an IBC client")
	ErrInvalidHandshake              = sdkerrors.Register(ModuleName, 1107, "invalid handshake")
	ErrNoConsumerClient              = sdkerrors.Register(ModuleName, 1108, "consumer IBC client doesn't exist")
	ErrConsumerConnectionEstablished = sdkerrors.Register(ModuleName, 1109, "consumer connection already established")
	ErrInvalidClient                 = sdkerrors.Register(ModuleName, 1110, "invalid IBC client")
	ErrJSONUnmarshal                 = sdkerrors.Register(ModuleName, 1111, "failed to unmarshal JSON")
	ErrJSONMarshal                   = sdkerrors.Register(ModuleName, 1112, "failed to marshal JSON")
	ErrNotImplemented                = sdkerrors.Register(ModuleName, 1113, "not implemented")
	ErrUnrecognizedAckType           = sdkerrors.Register(ModuleName, 1114, "unrecognized acknowledgement type")
	ErrUnrecognizedPacketType        = sdkerrors.Register(ModuleName, 1115, "unrecognized packet type")
	ErrCannotCloseChannel            = sdkerrors.Register(ModuleName, 1116, "user cannot close channel")
)
