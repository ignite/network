package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	connectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ignterrors "github.com/ignite/modules/pkg/errors"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	launchtypes "github.com/ignite/network/x/launch/types"
	"github.com/ignite/network/x/monitoringc/types"
)

// testSetupWithFooClient returns a test setup with monitoring keeper containing necessary IBC mocks for a client with ID foo
func testSetupWithFooClient(t *testing.T) (context.Context, testkeeper.TestKeepers, testkeeper.TestMsgServers) {
	return testkeeper.NewTestSetupWithIBCMocks(
		t,
		[]testkeeper.Connection{
			{
				ConnID: "foo",
				Conn: connectiontypes.ConnectionEnd{
					ClientId: "foo",
				},
			},
		},
		[]testkeeper.Channel{
			{
				ChannelID: "foo",
				Channel: channeltypes.Channel{
					ConnectionHops: []string{"foo"},
				},
			},
		},
	)
}

func TestKeeper_VerifyClientIDFromChannelID(t *testing.T) {
	t.Run("should return no error if the client is verified and provider has no connection yet", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		clientID := "foo"
		err := tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(sdkCtx, clientID, types.LaunchIDFromVerifiedClientID{
			LaunchID: 1,
			ClientID: clientID,
		})
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(sdkCtx, "foo")
		require.NoError(t, err)
	})

	t.Run("should fail if connection doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(sdkCtx, "bar")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})

	t.Run("should fail if the client is not verified", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		err := tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(sdkCtx, "foo")
		require.ErrorIs(t, err, types.ErrClientNotVerified)
	})

	t.Run("should fail if the provider already has an established connection", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		launchID := uint64(1)
		err := tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(sdkCtx, "foo", types.LaunchIDFromVerifiedClientID{
			LaunchID: launchID,
			ClientID: "foo",
		})
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.ProviderClientID.Set(sdkCtx, launchID, types.ProviderClientID{
			LaunchID: launchID,
			ClientID: "bar",
		})
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.VerifyClientIDFromConnID(sdkCtx, "foo")
		require.ErrorIs(t, err, types.ErrConnectionAlreadyEstablished)
	})
}

func TestKeeper_RegisterProviderClientIDFromChannelID(t *testing.T) {
	t.Run("should register the client id", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		clientID := "foo"
		launchID := uint64(1)
		err := tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(sdkCtx, clientID, types.LaunchIDFromVerifiedClientID{
			LaunchID: launchID,
			ClientID: clientID,
		})
		require.NoError(t, err)

		chain := launchtypes.Chain{
			LaunchID: launchID,
		}
		err = tk.LaunchKeeper.Chain.Set(sdkCtx, launchID, chain)
		require.NoError(t, err)

		err = tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(sdkCtx, "foo")
		require.NoError(t, err)

		// check that the chain is properly set to have MonitoringConnected be true
		chain, err = tk.LaunchKeeper.GetChain(sdkCtx, launchID)
		require.NoError(t, err)
		require.True(t, chain.MonitoringConnected)

		// the provider client ID should be created
		pCid, err := tk.MonitoringConsumerKeeper.ProviderClientID.Get(sdkCtx, 1)
		require.NoError(t, err)
		require.EqualValues(t, launchID, pCid.LaunchID)
		require.EqualValues(t, "foo", pCid.ClientID)

		// the channel ID is associated with the correct launch ID
		launchIDFromChanID, err := tk.MonitoringConsumerKeeper.LaunchIDFromChannelID.Get(sdkCtx, "foo")
		require.NoError(t, err)
		require.EqualValues(t, launchID, launchIDFromChanID.LaunchID)
		require.EqualValues(t, "foo", launchIDFromChanID.ChannelID)
	})

	t.Run("should fail if the channel doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(sdkCtx, "bar")
		require.ErrorIs(t, err, channeltypes.ErrChannelNotFound)
	})

	t.Run("should fail with critical error if the client is not verified", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(sdkCtx, "foo")
		require.ErrorIs(t, err, ignterrors.ErrCritical)
	})

	t.Run("should fail if the provider already has an established connection", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		launchID := uint64(1)
		err := tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(sdkCtx, "foo", types.LaunchIDFromVerifiedClientID{
			LaunchID: launchID,
			ClientID: "foo",
		})
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.ProviderClientID.Set(sdkCtx, launchID, types.ProviderClientID{
			LaunchID: launchID,
			ClientID: "bar",
		})
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(sdkCtx, "foo")
		require.ErrorIs(t, err, types.ErrConnectionAlreadyEstablished)
	})

	t.Run("should fail if monitoring connection already enabled on chain", func(t *testing.T) {
		ctx, tk, _ := testSetupWithFooClient(t)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		launchID := uint64(1)
		err := tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(sdkCtx, "foo", types.LaunchIDFromVerifiedClientID{
			LaunchID: launchID,
			ClientID: "foo",
		})
		require.NoError(t, err)
		chain := launchtypes.Chain{
			LaunchID:            launchID,
			MonitoringConnected: true,
		}
		err = tk.LaunchKeeper.Chain.Set(sdkCtx, launchID, chain)
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(sdkCtx, "foo")
		require.ErrorIs(t, err, launchtypes.ErrChainMonitoringConnected)
	})

	t.Run("should fail if the channel has more than 1 hop connection", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocks(
			t,
			[]testkeeper.Connection{},
			[]testkeeper.Channel{
				{
					ChannelID: "foo",
					Channel: channeltypes.Channel{
						ConnectionHops: []string{"foo", "bar"},
					},
				},
			},
		)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(sdkCtx, "foo")
		require.ErrorIs(t, err, channeltypes.ErrTooManyConnectionHops)
	})

	t.Run("should fail if the connection doesn't exist", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithIBCMocks(
			t,
			[]testkeeper.Connection{},
			[]testkeeper.Channel{
				{
					ChannelID: "foo",
					Channel: channeltypes.Channel{
						ConnectionHops: []string{"foo"},
					},
				},
			},
		)
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		err := tk.MonitoringConsumerKeeper.RegisterProviderClientIDFromChannelID(sdkCtx, "foo")
		require.ErrorIs(t, err, connectiontypes.ErrConnectionNotFound)
	})
}
