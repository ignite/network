package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/profile/types"
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

		Schema                     collections.Schema
		Params                     collections.Item[types.Params]
		CoordinatorSeq             collections.Sequence
		Coordinator                collections.Map[uint64, types.Coordinator]
		CoordinatorByAddress       collections.Map[sdk.AccAddress, types.CoordinatorByAddress]
		Validator                  collections.Map[string, types.Validator]
		ValidatorByOperatorAddress collections.Map[string, types.ValidatorByOperatorAddress]
		// this line is used by starport scaffolding # collection/type
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:                        cdc,
		addressCodec:               addressCodec,
		storeService:               storeService,
		authority:                  authority,
		logger:                     logger,
		Params:                     collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		CoordinatorSeq:             collections.NewSequence(sb, types.CoordinatorCountKey, "coordinator_seq"),
		Coordinator:                collections.NewMap(sb, types.CoordinatorKey, "coordinator", collections.Uint64Key, codec.CollValue[types.Coordinator](cdc)),
		CoordinatorByAddress:       collections.NewMap(sb, types.CoordinatorByAddressKey, "coordinator_by_address", sdk.LengthPrefixedAddressKey(sdk.AccAddressKey), codec.CollValue[types.CoordinatorByAddress](cdc)),
		Validator:                  collections.NewMap(sb, types.ValidatorKey, "validator", collections.StringKey, codec.CollValue[types.Validator](cdc)),
		ValidatorByOperatorAddress: collections.NewMap(sb, types.ValidatorByOperatorAddressKey, "validator_by_operator_address", collections.StringKey, codec.CollValue[types.ValidatorByOperatorAddress](cdc)),
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
