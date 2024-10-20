package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	launchtypes "github.com/ignite/network/x/launch/types"
	"github.com/ignite/network/x/project/types"
)

type LaunchKeeper interface {
	GetChain(ctx context.Context, launchID uint64) (launchtypes.Chain, error)
	CreateNewChain(
		ctx context.Context,
		coordinatorID uint64,
		genesisChainID,
		sourceURL,
		sourceHash string,
		initialGenesis launchtypes.InitialGenesis,
		hasProject bool,
		projectID uint64,
		isMainnet bool,
		accountBalance sdk.Coins,
		metadata []byte,
	) (uint64, error)
}

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		addressCodec address.Codec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message.
		// Typically, this should be the x/gov module account.
		authority string

		Schema         collections.Schema
		Params         collections.Item[types.Params]
		MainnetAccount collections.Map[collections.Pair[uint64, sdk.AccAddress], types.MainnetAccount]
		ProjectSeq     collections.Sequence
		Project        collections.Map[uint64, types.Project]
		ProjectChains  collections.Map[uint64, types.ProjectChains]
		TotalShares    collections.Item[uint64]
		// this line is used by starport scaffolding # collection/type

		distributionKeeper types.DistributionKeeper
		profileKeeper      types.ProfileKeeper
		bankKeeper         types.BankKeeper
		launchKeeper       LaunchKeeper
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
	bankKeeper types.BankKeeper,
	launchKeeper LaunchKeeper,
) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:                cdc,
		addressCodec:       addressCodec,
		storeService:       storeService,
		authority:          authority,
		logger:             logger,
		distributionKeeper: distributionKeeper,
		profileKeeper:      profileKeeper,
		bankKeeper:         bankKeeper,
		launchKeeper:       launchKeeper,
		Params:             collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		MainnetAccount:     collections.NewMap(sb, types.MainnetAccountKey, "mainnet_account", collections.PairKeyCodec(collections.Uint64Key, sdk.LengthPrefixedAddressKey(sdk.AccAddressKey)), codec.CollValue[types.MainnetAccount](cdc)),
		ProjectSeq:         collections.NewSequence(sb, types.ProjectCountKey, "project_seq"),
		Project:            collections.NewMap(sb, types.ProjectKey, "project", collections.Uint64Key, codec.CollValue[types.Project](cdc)),
		ProjectChains:      collections.NewMap(sb, types.ProjectChainsKey, "project_chains", collections.Uint64Key, codec.CollValue[types.ProjectChains](cdc)),
		TotalShares:        collections.NewItem(sb, types.TotalSharesKey, "total_shares", collections.Uint64Value),
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
