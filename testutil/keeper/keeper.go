// Package keeper provides methods to initialize SDK keepers with local storage for test purposes
package keeper

import (
	"testing"
	"time"

	"cosmossdk.io/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	portkeeper "github.com/cosmos/ibc-go/v8/modules/core/05-port/keeper"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	claimkeeper "github.com/ignite/modules/x/claim/keeper"
	claimtypes "github.com/ignite/modules/x/claim/types"
	fundraisingkeeper "github.com/ignite/modules/x/fundraising/keeper"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	"github.com/ignite/network/testutil/keeper/mocks"
	launchkeeper "github.com/ignite/network/x/launch/keeper"
	launchtypes "github.com/ignite/network/x/launch/types"
	monitoringckeeper "github.com/ignite/network/x/monitoringc/keeper"
	monitoringctypes "github.com/ignite/network/x/monitoringc/types"
	monitoringpkeeper "github.com/ignite/network/x/monitoringp/keeper"
	participationkeeper "github.com/ignite/network/x/participation/keeper"
	participationtypes "github.com/ignite/network/x/participation/types"
	profilekeeper "github.com/ignite/network/x/profile/keeper"
	profiletypes "github.com/ignite/network/x/profile/types"
	projectkeeper "github.com/ignite/network/x/project/keeper"
	projecttypes "github.com/ignite/network/x/project/types"
	rewardkeeper "github.com/ignite/network/x/reward/keeper"
	rewardtypes "github.com/ignite/network/x/reward/types"
)

var (
	// ExampleTimestamp is a timestamp used as the current time for the context of the keepers returned from the package
	ExampleTimestamp = time.Date(2020, time.January, 1, 12, 0, 0, 0, time.UTC)

	// ExampleHeight is a block height used as the current block height for the context of test keeper
	ExampleHeight = int64(1111)
)

// HookMocks holds mocks for the module hooks
type HookMocks struct {
	LaunchHooksMock *mocks.LaunchHooks
}

// TestKeepers holds all keepers used during keeper tests for all modules
type TestKeepers struct {
	T                        testing.TB
	ProjectKeeper            projectkeeper.Keeper
	LaunchKeeper             *launchkeeper.Keeper
	ProfileKeeper            profilekeeper.Keeper
	RewardKeeper             rewardkeeper.Keeper
	MonitoringConsumerKeeper *monitoringckeeper.Keeper
	MonitoringProviderKeeper monitoringpkeeper.Keeper
	AccountKeeper            authkeeper.AccountKeeper
	BankKeeper               bankkeeper.Keeper
	DistrKeeper              distrkeeper.Keeper
	IBCKeeper                *ibckeeper.Keeper
	StakingKeeper            *stakingkeeper.Keeper
	FundraisingKeeper        fundraisingkeeper.Keeper
	ParticipationKeeper      participationkeeper.Keeper
	ClaimKeeper              claimkeeper.Keeper
	HooksMocks               HookMocks
}

// TestMsgServers holds all message servers used during keeper tests for all modules
type TestMsgServers struct {
	T                testing.TB
	ProfileSrv       profiletypes.MsgServer
	LaunchSrv        launchtypes.MsgServer
	ProjectSrv       projecttypes.MsgServer
	RewardSrv        rewardtypes.MsgServer
	MonitoringcSrv   monitoringctypes.MsgServer
	ParticipationSrv participationtypes.MsgServer
	ClaimSrv         claimtypes.MsgServer
}

// SetupOption represents an option that can be provided to NewTestSetup
type SetupOption func(*setupOptions)

// setupOptions represents the set of SetupOption
type setupOptions struct {
	LaunchHooksMock bool
}

// WithLaunchHooksMock sets a mock for the hooks in testing launch keeper
func WithLaunchHooksMock() func(*setupOptions) {
	return func(o *setupOptions) {
		o.LaunchHooksMock = true
	}
}

// NewTestSetup returns initialized instances of all the keepers and message servers of the modules
func NewTestSetup(t testing.TB, options ...SetupOption) (sdk.Context, TestKeepers, TestMsgServers) {
	// setup options
	var so setupOptions
	for _, option := range options {
		option(&so)
	}

	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingKeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	distrKeeper := initializer.Distribution(authKeeper, bankKeeper, stakingKeeper)
	upgradeKeeper := initializer.Upgrade()
	scopedKeeper := capabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	ibcKeeper := initializer.IBC(paramKeeper, stakingKeeper, scopedKeeper, upgradeKeeper)
	fundraisingKeeper := initializer.Fundraising(authKeeper, bankKeeper, distrKeeper)
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, distrKeeper)
	rewardKeeper := initializer.Reward(authKeeper, bankKeeper, profileKeeper, launchKeeper)
	projectKeeper := initializer.Project(launchKeeper, profileKeeper, bankKeeper, distrKeeper)
	participationKeeper := initializer.Participation(fundraisingKeeper, stakingKeeper)
	err := launchKeeper.SetProjectKeeper(projectKeeper)
	require.NoError(t, err)
	portKeeper := portkeeper.NewKeeper(scopedKeeper)
	monitoringConsumerKeeper, err := initializer.Monitoringc(
		ibcKeeper,
		*capabilityKeeper,
		portKeeper,
		launchKeeper,
		rewardKeeper,
		[]Connection{},
		[]Channel{},
	)
	require.NoError(t, err)
	err = launchKeeper.SetMonitoringcKeeper(monitoringConsumerKeeper)
	require.NoError(t, err)
	claimKeeper := initializer.Claim(authKeeper, distrKeeper, bankKeeper)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time:   ExampleTimestamp,
		Height: ExampleHeight,
	}, false, log.NewNopLogger())

	// Initialize community pool

	require.NoError(t, distrKeeper.FeePool.Set(ctx, distrtypes.InitialFeePool()))

	// Initialize params
	require.NoError(t, distrKeeper.Params.Set(ctx, distrtypes.DefaultParams()))
	require.NoError(t, stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams()))
	require.NoError(t, launchKeeper.Params.Set(ctx, launchtypes.DefaultParams()))
	require.NoError(t, rewardKeeper.Params.Set(ctx, rewardtypes.DefaultParams()))
	require.NoError(t, projectKeeper.Params.Set(ctx, projecttypes.DefaultParams()))
	fundraisingParams := fundraisingtypes.DefaultParams()
	fundraisingParams.AuctionCreationFee = sdk.NewCoins()
	require.NoError(t, fundraisingKeeper.Params.Set(ctx, fundraisingParams))
	require.NoError(t, participationKeeper.Params.Set(ctx, participationtypes.DefaultParams()))
	require.NoError(t, monitoringConsumerKeeper.Params.Set(ctx, monitoringctypes.DefaultParams()))
	require.NoError(t, claimKeeper.Params.Set(ctx, claimtypes.DefaultParams()))
	setIBCDefaultParams(ctx, ibcKeeper)

	// Set hooks
	var hooksMocks HookMocks
	if so.LaunchHooksMock {
		var err error
		launchHooksMock := mocks.NewLaunchHooks(t)
		err = launchKeeper.SetHooks(launchHooksMock)
		require.NoError(t, err)
		hooksMocks.LaunchHooksMock = launchHooksMock
	}

	profileSrv := profilekeeper.NewMsgServerImpl(profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(launchKeeper)
	projectSrv := projectkeeper.NewMsgServerImpl(projectKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(rewardKeeper)
	monitoringcSrv := monitoringckeeper.NewMsgServerImpl(monitoringConsumerKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(participationKeeper)
	claimSrv := claimkeeper.NewMsgServerImpl(claimKeeper)

	// set max shares - only set during app InitGenesis
	require.NoError(t, projectKeeper.TotalShares.Set(ctx, networktypes.TotalShareNumber))

	return ctx, TestKeepers{
			T:                        t,
			ProjectKeeper:            projectKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringConsumerKeeper: monitoringConsumerKeeper,
			AccountKeeper:            authKeeper,
			BankKeeper:               bankKeeper,
			DistrKeeper:              distrKeeper,
			IBCKeeper:                ibcKeeper,
			StakingKeeper:            stakingKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
			ClaimKeeper:              claimKeeper,
			HooksMocks:               hooksMocks,
		}, TestMsgServers{
			T:                t,
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			ProjectSrv:       projectSrv,
			RewardSrv:        rewardSrv,
			MonitoringcSrv:   monitoringcSrv,
			ParticipationSrv: participationSrv,
			ClaimSrv:         claimSrv,
		}
}

// NewTestSetupWithIBCMocks returns a keeper of the monitoring consumer module for testing purpose with mocks for IBC keepers
func NewTestSetupWithIBCMocks(
	t testing.TB,
	connectionMock []Connection,
	channelMock []Channel,
) (sdk.Context, TestKeepers, TestMsgServers) {
	initializer := newInitializer()

	paramKeeper := initializer.Param()
	capabilityKeeper := initializer.Capability()
	authKeeper := initializer.Auth(paramKeeper)
	bankKeeper := initializer.Bank(paramKeeper, authKeeper)
	stakingKeeper := initializer.Staking(authKeeper, bankKeeper, paramKeeper)
	distrKeeper := initializer.Distribution(authKeeper, bankKeeper, stakingKeeper)
	upgradeKeeper := initializer.Upgrade()
	scopedKeeper := capabilityKeeper.ScopeToModule(ibcexported.ModuleName)
	ibcKeeper := initializer.IBC(paramKeeper, stakingKeeper, scopedKeeper, upgradeKeeper)
	fundraisingKeeper := initializer.Fundraising(authKeeper, bankKeeper, distrKeeper)
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, distrKeeper)
	rewardKeeper := initializer.Reward(authKeeper, bankKeeper, profileKeeper, launchKeeper)
	projectKeeper := initializer.Project(launchKeeper, profileKeeper, bankKeeper, distrKeeper)
	participationKeeper := initializer.Participation(fundraisingKeeper, stakingKeeper)
	err := launchKeeper.SetProjectKeeper(projectKeeper)
	require.NoError(t, err)
	portKeeper := portkeeper.NewKeeper(scopedKeeper)
	monitoringConsumerKeeper, err := initializer.Monitoringc(
		ibcKeeper,
		*capabilityKeeper,
		portKeeper,
		launchKeeper,
		rewardKeeper,
		connectionMock,
		channelMock,
	)
	require.NoError(t, err)
	err = launchKeeper.SetMonitoringcKeeper(monitoringConsumerKeeper)
	require.NoError(t, err)
	claimKeeper := initializer.Claim(authKeeper, distrKeeper, bankKeeper)
	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	// Create a context using a custom timestamp
	ctx := sdk.NewContext(initializer.StateStore, tmproto.Header{
		Time:   ExampleTimestamp,
		Height: ExampleHeight,
	}, false, log.NewNopLogger())

	// Initialize community pool
	require.NoError(t, distrKeeper.FeePool.Set(ctx, distrtypes.InitialFeePool()))

	// Initialize params
	require.NoError(t, distrKeeper.Params.Set(ctx, distrtypes.DefaultParams()))
	require.NoError(t, stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams()))
	require.NoError(t, launchKeeper.Params.Set(ctx, launchtypes.DefaultParams()))
	require.NoError(t, rewardKeeper.Params.Set(ctx, rewardtypes.DefaultParams()))
	require.NoError(t, projectKeeper.Params.Set(ctx, projecttypes.DefaultParams()))
	require.NoError(t, fundraisingKeeper.Params.Set(ctx, fundraisingtypes.DefaultParams()))
	require.NoError(t, participationKeeper.Params.Set(ctx, participationtypes.DefaultParams()))
	require.NoError(t, monitoringConsumerKeeper.Params.Set(ctx, monitoringctypes.DefaultParams()))
	require.NoError(t, claimKeeper.Params.Set(ctx, claimtypes.DefaultParams()))
	setIBCDefaultParams(ctx, ibcKeeper)

	profileSrv := profilekeeper.NewMsgServerImpl(profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(launchKeeper)
	projectSrv := projectkeeper.NewMsgServerImpl(projectKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(rewardKeeper)
	monitoringcSrv := monitoringckeeper.NewMsgServerImpl(monitoringConsumerKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(participationKeeper)

	// set max shares - only set during app InitGenesis
	require.NoError(t, projectKeeper.TotalShares.Set(ctx, networktypes.TotalShareNumber))

	return ctx, TestKeepers{
			T:                        t,
			ProjectKeeper:            projectKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringConsumerKeeper: monitoringConsumerKeeper,
			AccountKeeper:            authKeeper,
			BankKeeper:               bankKeeper,
			IBCKeeper:                ibcKeeper,
			StakingKeeper:            stakingKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
			ClaimKeeper:              claimKeeper,
		}, TestMsgServers{
			T:                t,
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			ProjectSrv:       projectSrv,
			RewardSrv:        rewardSrv,
			MonitoringcSrv:   monitoringcSrv,
			ParticipationSrv: participationSrv,
		}
}

// setIBCDefaultParams set default params for IBC client and connection keepers
func setIBCDefaultParams(ctx sdk.Context, ibcKeeper *ibckeeper.Keeper) {
	ibcKeeper.ClientKeeper.SetParams(ctx, ibcclienttypes.DefaultParams())
	ibcKeeper.ConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())
	ibcKeeper.ClientKeeper.SetNextClientSequence(ctx, 0)
}
