package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ignite/network/x/reward/types"
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

		Schema     collections.Schema
		Params     collections.Item[types.Params]
		RewardPool collections.Map[uint64, types.RewardPool]
		// this line is used by starport scaffolding # collection/type

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
		profileKeeper types.ProfileKeeper
		launchKeeper  types.LaunchKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	profileKeeper types.ProfileKeeper,
	launchKeeper types.LaunchKeeper,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:           cdc,
		addressCodec:  addressCodec,
		storeService:  storeService,
		authority:     authority,
		logger:        logger,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		profileKeeper: profileKeeper,
		launchKeeper:  launchKeeper,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		RewardPool:    collections.NewMap(sb, types.RewardPoolKey, "reward_pool", collections.Uint64Key, codec.CollValue[types.RewardPool](cdc)),
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

// GetProfileKeeper gets the profile keeper interface of the module
func (k *Keeper) GetProfileKeeper() types.ProfileKeeper {
	return k.profileKeeper
}

// GetLaunchKeeper gets the profile keeper interface of the module
func (k *Keeper) GetLaunchKeeper() types.LaunchKeeper {
	return k.launchKeeper
}
