package keeper

import (
	"context"
	"math/rand"
	"testing"

	"cosmossdk.io/core/address"
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/keeper"
	"github.com/ignite/network/x/profile/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func ProfileKeeper(t testing.TB) (keeper.Keeper, sdk.Context, address.Codec) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	addressCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())

	k := keeper.NewKeeper(
		cdc,
		addressCodec,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	if err := k.Params.Set(ctx, types.DefaultParams()); err != nil {
		t.Fatalf("failed to set params: %v", err)
	}

	return k, ctx, addressCodec
}

// CreateCoordinator creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateCoordinator(ctx context.Context, r *rand.Rand) (id uint64, address sdk.AccAddress) {
	return tm.CreateCoordinatorWithAddr(ctx, r, sample.Address(r))
}

// CreateCoordinatorWithAddr creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateCoordinatorWithAddr(ctx context.Context, r *rand.Rand, address string) (uint64, sdk.AccAddress) {
	var (
		addr        = sdk.MustAccAddressFromBech32(address)
		description = sample.CoordinatorDescription(r)
	)
	res, err := tm.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:  address,
		Identity: description.Identity,
		Website:  description.Website,
		Details:  description.Details,
	})
	require.NoError(tm.T, err)
	return res.CoordinatorId, addr
}
