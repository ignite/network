package keeper

import (
	"testing"

	"cosmossdk.io/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	portkeeper "github.com/cosmos/ibc-go/v8/modules/core/05-port/keeper"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	launchkeeper "github.com/ignite/network/x/launch/keeper"
	launchtypes "github.com/ignite/network/x/launch/types"
	monitoringptypes "github.com/ignite/network/x/monitoringp/types"
	participationkeeper "github.com/ignite/network/x/participation/keeper"
	participationtypes "github.com/ignite/network/x/participation/types"
	profilekeeper "github.com/ignite/network/x/profile/keeper"
	projectkeeper "github.com/ignite/network/x/project/keeper"
	projecttypes "github.com/ignite/network/x/project/types"
	rewardkeeper "github.com/ignite/network/x/reward/keeper"
	rewardtypes "github.com/ignite/network/x/reward/types"
)

// NewTestSetupWithMonitoringp returns a test keepers struct and servers struct with the monitoring provider module
func NewTestSetupWithMonitoringp(t testing.TB) (sdk.Context, TestKeepers, TestMsgServers) {
	return NewTestSetupWithIBCMocksMonitoringp(t, []Connection{}, []Channel{})
}

// NewTestSetupWithIBCMocksMonitoringp returns a keeper of the monitoring provider module for testing purpose with mocks for IBC keepers
func NewTestSetupWithIBCMocksMonitoringp(
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
	portKeeper := portkeeper.NewKeeper(scopedKeeper)
	monitoringProviderKeeper := initializer.Monitoringp(
		stakingKeeper,
		ibcKeeper,
		*capabilityKeeper,
		portKeeper,
		connectionMock,
		channelMock,
	)
	fundraisingKeeper := initializer.Fundraising(authKeeper, bankKeeper, distrKeeper)
	profileKeeper := initializer.Profile()
	launchKeeper := initializer.Launch(profileKeeper, distrKeeper)
	rewardKeeper := initializer.Reward(authKeeper, bankKeeper, profileKeeper, launchKeeper)
	projectKeeper := initializer.Project(launchKeeper, profileKeeper, bankKeeper, distrKeeper)
	participationKeeper := initializer.Participation(fundraisingKeeper, stakingKeeper)
	launchKeeper.SetProjectKeeper(projectKeeper)

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
	require.NoError(t, monitoringProviderKeeper.Params.Set(ctx, monitoringptypes.DefaultParams()))
	setIBCDefaultParams(ctx, ibcKeeper)

	profileSrv := profilekeeper.NewMsgServerImpl(profileKeeper)
	launchSrv := launchkeeper.NewMsgServerImpl(launchKeeper)
	projectSrv := projectkeeper.NewMsgServerImpl(projectKeeper)
	rewardSrv := rewardkeeper.NewMsgServerImpl(rewardKeeper)
	participationSrv := participationkeeper.NewMsgServerImpl(participationKeeper)

	// set max shares - only set during app InitGenesis
	require.NoError(t, projectKeeper.TotalShares.Set(ctx, networktypes.TotalShareNumber))

	return ctx, TestKeepers{
			T:                        t,
			ProjectKeeper:            projectKeeper,
			LaunchKeeper:             launchKeeper,
			ProfileKeeper:            profileKeeper,
			RewardKeeper:             rewardKeeper,
			MonitoringProviderKeeper: monitoringProviderKeeper,
			BankKeeper:               bankKeeper,
			IBCKeeper:                ibcKeeper,
			StakingKeeper:            stakingKeeper,
			FundraisingKeeper:        fundraisingKeeper,
			ParticipationKeeper:      participationKeeper,
		}, TestMsgServers{
			T:                t,
			ProfileSrv:       profileSrv,
			LaunchSrv:        launchSrv,
			ProjectSrv:       projectSrv,
			RewardSrv:        rewardSrv,
			ParticipationSrv: participationSrv,
		}
}
