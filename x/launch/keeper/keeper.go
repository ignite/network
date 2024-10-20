package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/launch/types"
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

		Schema           collections.Schema
		Params           collections.Item[types.Params]
		ChainSeq         collections.Sequence
		Chain            collections.Map[uint64, types.Chain]
		GenesisAccount   collections.Map[collections.Pair[uint64, sdk.AccAddress], types.GenesisAccount]
		GenesisValidator collections.Map[collections.Pair[uint64, sdk.AccAddress], types.GenesisValidator]
		VestingAccount   collections.Map[collections.Pair[uint64, sdk.AccAddress], types.VestingAccount]
		RequestSeq       collections.Map[uint64, uint64]
		Request          collections.Map[collections.Pair[uint64, uint64], types.Request]
		ParamChange      collections.Map[collections.Pair[uint64, string], types.ParamChange]
		// this line is used by starport scaffolding # collection/type

		distributionKeeper types.DistributionKeeper
		profileKeeper      types.ProfileKeeper
		projectKeeper      types.ProjectKeeper
		monitoringcKeeper  types.MonitoringConsumerKeeper
		hooks              types.LaunchHooks
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	distributionKeeper types.DistributionKeeper,
	profileKeeper types.ProfileKeeper,
) *Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := &Keeper{
		cdc:                cdc,
		addressCodec:       addressCodec,
		storeService:       storeService,
		authority:          authority,
		logger:             logger,
		distributionKeeper: distributionKeeper,
		profileKeeper:      profileKeeper,
		Params:             collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		ChainSeq:           collections.NewSequence(sb, types.ChainCountKey, "chain_seq"),
		Chain:              collections.NewMap(sb, types.ChainKey, "chain", collections.Uint64Key, codec.CollValue[types.Chain](cdc)),
		GenesisAccount:     collections.NewMap(sb, types.GenesisAccountKey, "genesis_account", collections.PairKeyCodec(collections.Uint64Key, sdk.LengthPrefixedAddressKey(sdk.AccAddressKey)), codec.CollValue[types.GenesisAccount](cdc)),
		GenesisValidator:   collections.NewMap(sb, types.GenesisValidatorKey, "genesis_validator", collections.PairKeyCodec(collections.Uint64Key, sdk.LengthPrefixedAddressKey(sdk.AccAddressKey)), codec.CollValue[types.GenesisValidator](cdc)),
		VestingAccount:     collections.NewMap(sb, types.VestingAccountKey, "vesting_account", collections.PairKeyCodec(collections.Uint64Key, sdk.LengthPrefixedAddressKey(sdk.AccAddressKey)), codec.CollValue[types.VestingAccount](cdc)),
		RequestSeq:         collections.NewMap(sb, types.RequestCountKey, "request_seq", collections.Uint64Key, collections.Uint64Value),
		Request:            collections.NewMap(sb, types.RequestKey, "request", collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key), codec.CollValue[types.Request](cdc)),
		ParamChange:        collections.NewMap(sb, types.ParamChangeKey, "param_change", collections.PairKeyCodec(collections.Uint64Key, collections.StringKey), codec.CollValue[types.ParamChange](cdc)),
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

// SetProjectKeeper sets the project keeper interface of the module
func (k *Keeper) SetProjectKeeper(projectKeeper types.ProjectKeeper) {
	if k.projectKeeper != nil {
		panic("project keeper already set for launch module")
	}
	k.projectKeeper = projectKeeper
}

// GetProjectKeeper gets the project keeper interface of the module
func (k *Keeper) GetProjectKeeper() types.ProjectKeeper {
	return k.projectKeeper
}

// SetMonitoringcKeeper sets the monitoring consumer keeper interface of the module
func (k *Keeper) SetMonitoringcKeeper(monitoringcKeeper types.MonitoringConsumerKeeper) {
	if k.monitoringcKeeper != nil {
		panic("monitoring consumer keeper already set for launch module")
	}
	k.monitoringcKeeper = monitoringcKeeper
}

// GetMonitoringcKeeper gets the monitoring consumer keeper interface of the module
func (k *Keeper) GetMonitoringcKeeper() types.MonitoringConsumerKeeper {
	return k.monitoringcKeeper
}

// GetProfileKeeper gets the profile keeper interface of the module
func (k *Keeper) GetProfileKeeper() types.ProfileKeeper {
	return k.profileKeeper
}

// SetHooks sets the fundraising hooks.
func (k *Keeper) SetHooks(hooks types.LaunchHooks) error {
	if k.hooks != nil {
		return errors.New("cannot set launch hooks twice")
	}
	k.hooks = hooks
	return nil
}

// EnableMonitoringConnection sets a chain with MonitoringConnected set to true
func (k Keeper) EnableMonitoringConnection(ctx context.Context, launchID uint64) error {
	chain, err := k.GetChain(ctx, launchID)
	if err != nil {
		return err
	}

	if chain.MonitoringConnected {
		return types.ErrChainMonitoringConnected
	}
	chain.MonitoringConnected = true
	return k.Chain.Set(ctx, chain.LaunchId, chain)
}
