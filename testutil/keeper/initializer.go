package keeper

import (
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	upgradekeeper "cosmossdk.io/x/upgrade/keeper"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	portkeeper "github.com/cosmos/ibc-go/v8/modules/core/05-port/keeper"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	claimkeeper "github.com/ignite/modules/x/claim/keeper"
	claimtypes "github.com/ignite/modules/x/claim/types"
	fundraisingkeeper "github.com/ignite/modules/x/fundraising/keeper"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"
	minttypes "github.com/ignite/modules/x/mint/types"

	networktypes "github.com/ignite/network/pkg/types"
	"github.com/ignite/network/testutil/sample"
	launchkeeper "github.com/ignite/network/x/launch/keeper"
	launchtypes "github.com/ignite/network/x/launch/types"
	monitoringckeeper "github.com/ignite/network/x/monitoringc/keeper"
	monitoringctypes "github.com/ignite/network/x/monitoringc/types"
	monitoringpkeeper "github.com/ignite/network/x/monitoringp/keeper"
	monitoringptypes "github.com/ignite/network/x/monitoringp/types"
	participationkeeper "github.com/ignite/network/x/participation/keeper"
	participationtypes "github.com/ignite/network/x/participation/types"
	profilekeeper "github.com/ignite/network/x/profile/keeper"
	profiletypes "github.com/ignite/network/x/profile/types"
	projectkeeper "github.com/ignite/network/x/project/keeper"
	projecttypes "github.com/ignite/network/x/project/types"
	rewardkeeper "github.com/ignite/network/x/reward/keeper"
	rewardtypes "github.com/ignite/network/x/reward/types"
)

var moduleAccountPerms = map[string][]string{
	authtypes.FeeCollectorName:     nil,
	distrtypes.ModuleName:          nil,
	minttypes.ModuleName:           {authtypes.Minter},
	ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	projecttypes.ModuleName:        {authtypes.Minter, authtypes.Burner},
	stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	rewardtypes.ModuleName:         {authtypes.Minter, authtypes.Burner},
	fundraisingtypes.ModuleName:    nil,
	claimtypes.ModuleName:          {authtypes.Minter, authtypes.Burner},
}

// initializer allows to initialize each module keeper
type initializer struct {
	Codec      codec.Codec
	Amino      *codec.LegacyAmino
	DB         *dbm.MemDB
	StateStore store.CommitMultiStore
}

type invalid struct{}

type s struct {
	I int
}

func createTestCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cdc.RegisterConcrete(s{}, "test/s", nil)
	cdc.RegisterConcrete(invalid{}, "test/invalid", nil)
	return cdc
}

func newInitializer() initializer {
	db := dbm.NewMemDB()
	return initializer{
		DB:         db,
		Codec:      sample.Codec(),
		Amino:      createTestCodec(),
		StateStore: store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics()),
	}
}

// ModuleAccountAddrs returns all the app's module account addresses.
func ModuleAccountAddrs(maccPerms map[string][]string) map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (i initializer) Param() paramskeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(paramstypes.StoreKey)
	tkeys := storetypes.NewTransientStoreKey(paramstypes.TStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(tkeys, storetypes.StoreTypeTransient, i.DB)

	return paramskeeper.NewKeeper(
		i.Codec,
		i.Amino,
		storeKey,
		tkeys,
	)
}

func (i initializer) Auth(paramKeeper paramskeeper.Keeper) authkeeper.AccountKeeper {
	storeKey := storetypes.NewKVStoreKey(authtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	paramKeeper.Subspace(authtypes.ModuleName)

	return authkeeper.NewAccountKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authtypes.ProtoBaseAccount,
		moduleAccountPerms,
		addresscodec.NewBech32Codec(sdk.Bech32MainPrefix),
		networktypes.AccountAddressPrefix,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Bank(paramKeeper paramskeeper.Keeper, authKeeper authkeeper.AccountKeeper) bankkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(banktypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	paramKeeper.Subspace(banktypes.ModuleName)
	modAccAddrs := ModuleAccountAddrs(moduleAccountPerms)

	return bankkeeper.NewBaseKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authKeeper,
		modAccAddrs,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		log.NewNopLogger(),
	)
}

func (i initializer) Capability() *capabilitykeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(capabilitytypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(capabilitytypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, i.DB)

	return capabilitykeeper.NewKeeper(i.Codec, storeKey, memStoreKey)
}

// create mock ProtocolVersionSetter for UpgradeKeeper

type ProtocolVersionSetter struct{}

func (vs ProtocolVersionSetter) SetProtocolVersion(uint64) {}

func (i initializer) Upgrade() *upgradekeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(upgradetypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	skipUpgradeHeights := make(map[int64]bool)
	vs := ProtocolVersionSetter{}

	return upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		runtime.NewKVStoreService(storeKey),
		i.Codec,
		"",
		vs,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Staking(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	paramKeeper paramskeeper.Keeper,
) *stakingkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(stakingtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	paramKeeper.Subspace(stakingtypes.ModuleName)

	return stakingkeeper.NewKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authKeeper,
		bankKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		addresscodec.NewBech32Codec("cosmosvaloper"),
		addresscodec.NewBech32Codec("cosmosvalcons"),
	)
}

func (i initializer) IBC(
	paramKeeper paramskeeper.Keeper,
	stakingKeeper *stakingkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	upgradeKeeper *upgradekeeper.Keeper,
) *ibckeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(ibcexported.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	return ibckeeper.NewKeeper(
		i.Codec,
		storeKey,
		paramKeeper.Subspace(ibcexported.ModuleName),
		stakingKeeper,
		upgradeKeeper,
		scopedKeeper,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Distribution(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	stakingKeeper *stakingkeeper.Keeper,
) distrkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(distrtypes.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	return distrkeeper.NewKeeper(
		i.Codec,
		runtime.NewKVStoreService(storeKey),
		authKeeper,
		bankKeeper,
		stakingKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Profile() profilekeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(profiletypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(profiletypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return profilekeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)
}

func (i initializer) Launch(
	profileKeeper profilekeeper.Keeper,
	distrKeeper distrkeeper.Keeper,
) launchkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(launchtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(launchtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return launchkeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		distrKeeper,
		profileKeeper,
	)
}

func (i initializer) Project(
	launchKeeper launchkeeper.Keeper,
	profileKeeper profilekeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	distrKeeper distrkeeper.Keeper,
) projectkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(projecttypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(projecttypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return projectkeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		distrKeeper,
		profileKeeper,
		bankKeeper,
		launchKeeper,
	)
}

func (i initializer) Reward(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	profileKeeper profilekeeper.Keeper,
	launchKeeper launchkeeper.Keeper,
) rewardkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(rewardtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(rewardtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return rewardkeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		authKeeper,
		bankKeeper,
		profileKeeper,
		launchKeeper,
	)
}

func (i initializer) Monitoringc(
	ibcKeeper *ibckeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
	portKeeper portkeeper.Keeper,
	launchKeeper launchkeeper.Keeper,
	rewardKeeper rewardkeeper.Keeper,
	connectionMock []Connection,
	channelMock []Channel,
) monitoringckeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(monitoringctypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(monitoringctypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	// check if ibc mocks should be used for connection and channel
	var (
		connKeeper    monitoringctypes.ConnectionKeeper = ibcKeeper.ConnectionKeeper
		channelKeeper monitoringctypes.ChannelKeeper    = ibcKeeper.ChannelKeeper
	)
	if len(connectionMock) != 0 {
		connKeeper = NewConnectionMock(connectionMock)
	}
	if len(channelMock) != 0 {
		channelKeeper = NewChannelMock(channelMock)
	}

	scopeModule := capabilityKeeper.ScopeToModule(monitoringctypes.ModuleName)

	k := monitoringckeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		func() *ibckeeper.Keeper {
			return &ibckeeper.Keeper{
				PortKeeper: &portKeeper,
			}
		},
		func(string) capabilitykeeper.ScopedKeeper {
			return scopeModule
		},
		launchKeeper,
		rewardKeeper,
	)
	k.SetIBCKeeper(ibcKeeper)
	k.SetConnectionKeeper(connKeeper)
	k.SetChannelKeeper(channelKeeper)
	return k
}

func (i initializer) Monitoringp(
	stakingKeeper *stakingkeeper.Keeper,
	ibcKeeper ibckeeper.Keeper,
	capabilityKeeper capabilitykeeper.Keeper,
	portKeeper portkeeper.Keeper,
	connectionMock []Connection,
	channelMock []Channel,
) monitoringpkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(monitoringptypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(monitoringptypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	// check if ibc mocks should be used for connection and channel
	var (
		connKeeper    monitoringctypes.ConnectionKeeper = ibcKeeper.ConnectionKeeper
		channelKeeper monitoringctypes.ChannelKeeper    = ibcKeeper.ChannelKeeper
	)
	if len(connectionMock) != 0 {
		connKeeper = NewConnectionMock(connectionMock)
	}
	if len(channelMock) != 0 {
		channelKeeper = NewChannelMock(channelMock)
	}

	scopeModule := capabilityKeeper.ScopeToModule(monitoringctypes.ModuleName)

	k := monitoringpkeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		func() *ibckeeper.Keeper {
			return &ibckeeper.Keeper{
				PortKeeper: &portKeeper,
			}
		},
		func(string) capabilitykeeper.ScopedKeeper {
			return scopeModule
		},
		stakingKeeper,
	)
	k.SetIBCKeepers(
		ibcKeeper.ClientKeeper,
		connKeeper,
		channelKeeper,
		ibcKeeper.PortKeeper,
	)
	return k
}

func (i initializer) Fundraising(
	authKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	disKeeper distrkeeper.Keeper,
) fundraisingkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(fundraisingtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(fundraisingtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return fundraisingkeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		authKeeper,
		bankKeeper,
		disKeeper,
	)
}

func (i initializer) Participation(
	fundraisingKeeper fundraisingkeeper.Keeper,
	stakingKeeper *stakingkeeper.Keeper,
) participationkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(participationtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(participationtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return participationkeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		stakingKeeper,
		fundraisingKeeper,
	)
}

func (i initializer) Claim(
	accountKeeper authkeeper.AccountKeeper,
	distrKeeper distrkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) claimkeeper.Keeper {
	storeKey := storetypes.NewKVStoreKey(claimtypes.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(claimtypes.MemStoreKey)

	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return claimkeeper.NewKeeper(
		i.Codec,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		accountKeeper,
		bankKeeper,
		distrKeeper,
	)
}
