package keeper

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	"github.com/ignite/network/x/monitoringc/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		addressCodec address.Codec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message.
		// Typically, this should be the x/gov module account.
		authority string

		Schema                       collections.Schema
		Params                       collections.Item[types.Params]
		LaunchIDFromChannelID        collections.Map[string, types.LaunchIDFromChannelID]
		LaunchIDFromVerifiedClientID collections.Map[string, types.LaunchIDFromVerifiedClientID]
		MonitoringHistory            collections.Map[uint64, types.MonitoringHistory]
		VerifiedClientID             collections.Map[uint64, types.VerifiedClientID]
		ProviderClientID             collections.Map[uint64, types.ProviderClientID]
		// this line is used by starport scaffolding # collection/type

		ibcKeeperFn        func() *ibckeeper.Keeper
		capabilityScopedFn func(string) capabilitykeeper.ScopedKeeper

		launchKeeper types.LaunchKeeper
		rewardKeeper types.RewardKeeper

		clientKeeper     types.ClientKeeper
		portKeeper       types.PortKeeper
		connectionKeeper types.ConnectionKeeper
		channelKeeper    types.ChannelKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	ibcKeeperFn func() *ibckeeper.Keeper,
	capabilityScopedFn func(string) capabilitykeeper.ScopedKeeper,
	launchKeeper types.LaunchKeeper,
	rewardKeeper types.RewardKeeper,
) *Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := &Keeper{
		cdc:                          cdc,
		addressCodec:                 addressCodec,
		storeService:                 storeService,
		authority:                    authority,
		logger:                       logger,
		ibcKeeperFn:                  ibcKeeperFn,
		capabilityScopedFn:           capabilityScopedFn,
		launchKeeper:                 launchKeeper,
		rewardKeeper:                 rewardKeeper,
		Params:                       collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		LaunchIDFromChannelID:        collections.NewMap(sb, types.LaunchIDFromChannelIDKey, "launch_id_from_channel_id", collections.StringKey, codec.CollValue[types.LaunchIDFromChannelID](cdc)),
		LaunchIDFromVerifiedClientID: collections.NewMap(sb, types.LaunchIDFromVerifiedClientIDKey, "launch_id_from_verified_client_id", collections.StringKey, codec.CollValue[types.LaunchIDFromVerifiedClientID](cdc)),
		MonitoringHistory:            collections.NewMap(sb, types.MonitoringHistoryKey, "monitoring_history", collections.Uint64Key, codec.CollValue[types.MonitoringHistory](cdc)),
		VerifiedClientID:             collections.NewMap(sb, types.VerifiedClientIDKey, "verified_client_id", collections.Uint64Key, codec.CollValue[types.VerifiedClientID](cdc)),
		ProviderClientID:             collections.NewMap(sb, types.ProviderClientIDKey, "provider_client_id", collections.Uint64Key, codec.CollValue[types.ProviderClientID](cdc)),
		// this line is used by starport scaffolding # collection/instantiate
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// AddressCodec returns the address codec.
func (k Keeper) AddressCodec() address.Codec {
	return k.addressCodec
}

// ----------------------------------------------------------------------------
// IBC Keeper Logic
// ----------------------------------------------------------------------------

// SetIBCKeeper sets the IBC Keeper
func (k *Keeper) SetIBCKeeper(ibcKeeper *ibckeeper.Keeper) {
	k.SetClientKeeper(ibcKeeper.ClientKeeper)
	k.SetPortKeeper(ibcKeeper.PortKeeper)
	k.SetConnectionKeeper(ibcKeeper.ConnectionKeeper)
	k.SetChannelKeeper(ibcKeeper.ChannelKeeper)
}

// SetClientKeeper sets IBC client keeper
func (k *Keeper) SetClientKeeper(clientKeeper types.ClientKeeper) {
	k.clientKeeper = clientKeeper
}

// SetPortKeeper sets IBC port keeper
func (k *Keeper) SetPortKeeper(portKeeper types.PortKeeper) {
	k.portKeeper = portKeeper
}

// SetConnectionKeeper sets IBC connection keeper
func (k *Keeper) SetConnectionKeeper(connectionKeeper types.ConnectionKeeper) {
	k.connectionKeeper = connectionKeeper
}

// SetChannelKeeper sets IBC channel keeper
func (k *Keeper) SetChannelKeeper(channelKeeper types.ChannelKeeper) {
	k.channelKeeper = channelKeeper
}

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.ScopedKeeper().GetCapability(ctx, capName)
	if !ok {
		return errorsmod.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.ibcKeeperFn().ChannelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// ShouldBound checks if the IBC app module can be bound to the desired port
func (k Keeper) ShouldBound(ctx sdk.Context, portID string) bool {
	scopedKeeper := k.ScopedKeeper()
	if scopedKeeper == nil {
		return false
	}
	_, ok := scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return !ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.ibcKeeperFn().PortKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the IBC app module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the IBC app module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.ScopedKeeper().AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.ScopedKeeper().ClaimCapability(ctx, cap, name)
}

// ScopedKeeper returns the ScopedKeeper
func (k Keeper) ScopedKeeper() exported.ScopedKeeper {
	return k.capabilityScopedFn(types.ModuleName)
}

// ClearVerifiedClientIdList removes a set of verifiedClientID in the store from its launch ID
func (k Keeper) ClearVerifiedClientIdList(ctx context.Context, launchID uint64) error {
	verifiedClientID, err := k.VerifiedClientID.Get(ctx, launchID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil
		}
		return err
	}
	if err := k.VerifiedClientID.Remove(ctx, launchID); err != nil {
		return err
	}

	for _, clientID := range verifiedClientID.GetClientIdList() {
		if err := k.LaunchIDFromVerifiedClientID.Remove(ctx, clientID); err != nil {
			return err
		}
	}
	return nil
}
