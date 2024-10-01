package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/participation/types"
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

		Schema                 collections.Schema
		Params                 collections.Item[types.Params]
		AuctionUsedAllocations collections.Map[collections.Pair[sdk.AccAddress, uint64], types.AuctionUsedAllocations]
		UsedAllocations        collections.Map[string, types.UsedAllocations]
		// this line is used by starport scaffolding # collection/type

		stakingKeeper     types.StakingKeeper
		fundraisingKeeper types.FundraisingKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
	stakingKeeper types.StakingKeeper,
	fundraisingKeeper types.FundraisingKeeper,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:                    cdc,
		addressCodec:           addressCodec,
		storeService:           storeService,
		authority:              authority,
		logger:                 logger,
		stakingKeeper:          stakingKeeper,
		fundraisingKeeper:      fundraisingKeeper,
		Params:                 collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		AuctionUsedAllocations: collections.NewMap(sb, types.AuctionUsedAllocationsKey, "auctionUsedAllocations", collections.PairKeyCodec(sdk.LengthPrefixedAddressKey(sdk.AccAddressKey), collections.Uint64Key), codec.CollValue[types.AuctionUsedAllocations](cdc)),
		UsedAllocations:        collections.NewMap(sb, types.UsedAllocationsKey, "usedAllocations", collections.StringKey, codec.CollValue[types.UsedAllocations](cdc)),
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
