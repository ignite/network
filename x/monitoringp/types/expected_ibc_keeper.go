package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
)

// ClientKeeper is imported to add the ability to create IBC Client from the module
type ClientKeeper interface {
	CreateClient(ctx sdk.Context, clientState exported.ClientState, consensusState exported.ConsensusState) (string, error)
	// TODO check IBC keeper methods
	// CreateClient(ctx context.Context, clientType string, clientState, consensusState []byte) (string, error)
}

// ConnectionKeeper is imported to check client ID during IBC handshake
type ConnectionKeeper interface {
	GetConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool)
}

// ChannelKeeper defines the expected IBC channel keeper.
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, portID, channelID string) (channeltypes.Channel, bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(
		ctx sdk.Context,
		channelCap *capabilitytypes.Capability,
		sourcePort string,
		sourceChannel string,
		timeoutHeight clienttypes.Height,
		timeoutTimestamp uint64,
		data []byte,
	) (uint64, error)
	ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error
}

// PortKeeper defines the expected IBC port keeper.
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}

// ScopedKeeper defines the expected IBC scoped keeper.
type ScopedKeeper interface {
	GetCapability(ctx sdk.Context, name string) (*capabilitytypes.Capability, bool)
	AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool
	ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error
}
