package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/monitoringc/types"
)

func TestVerifiedClientIDGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should allow get", func(t *testing.T) {
		items := createNVerifiedClientID(ctx, tk.MonitoringConsumerKeeper, 10)
		for _, item := range items {
			rst, err := tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, item.LaunchID)
			require.NoError(t, err)
			require.Equal(t,
				nullify.Fill(&item),
				nullify.Fill(&rst),
			)
		}
	})
}

func TestVerifiedClientIDClear(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should successfully clear entries", func(t *testing.T) {
		items := createNVerifiedClientID(ctx, tk.MonitoringConsumerKeeper, 1)
		launchID := items[0].LaunchID
		clientID := items[0].ClientIDs[0]

		err := tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(ctx, clientID, types.LaunchIDFromVerifiedClientID{
			ClientID: clientID,
			LaunchID: launchID,
		})
		require.NoError(t, err)
		rst, err := tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, launchID)
		require.NoError(t, err)
		require.Equal(t,
			nullify.Fill(&items[0]),
			nullify.Fill(&rst),
		)

		err = tk.MonitoringConsumerKeeper.ClearVerifiedClientIDs(ctx, launchID)
		require.NoError(t, err)
		_, err = tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, launchID)
		require.ErrorIs(t, err, collections.ErrNotFound)
		_, err = tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Get(ctx, clientID)
		require.ErrorIs(t, err, collections.ErrNotFound)
	})
}

func TestVerifiedClientIDGetAll(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should allow get all", func(t *testing.T) {
		items := createNVerifiedClientID(ctx, tk.MonitoringConsumerKeeper, 10)
		allVerifiedClientID, err := tk.MonitoringConsumerKeeper.AllVerifiedClientID(ctx)
		require.NoError(t, err)
		require.ElementsMatch(t,
			nullify.Fill(items),
			nullify.Fill(allVerifiedClientID),
		)
	})
}

func TestAddVerifiedClientID(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should update a verified client id", func(t *testing.T) {
		var (
			launchID         = uint64(1)
			newClientID      = "2"
			verifiedClientID = types.VerifiedClientID{
				LaunchID:  launchID,
				ClientIDs: []string{"1"},
			}
		)
		err := tk.MonitoringConsumerKeeper.VerifiedClientID.Set(ctx, launchID, verifiedClientID)
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, newClientID)
		require.NoError(t, err)
		got, err := tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, launchID)
		require.NoError(t, err)
		verifiedClientID.ClientIDs = append(verifiedClientID.ClientIDs, newClientID)
		require.Equal(t, verifiedClientID, got)
	})

	t.Run("should update a duplicated verified client id", func(t *testing.T) {
		var (
			launchID         = uint64(2)
			newClientID      = "2"
			verifiedClientID = types.VerifiedClientID{
				LaunchID:  launchID,
				ClientIDs: []string{"1", newClientID},
			}
		)
		err := tk.MonitoringConsumerKeeper.VerifiedClientID.Set(ctx, launchID, verifiedClientID)
		require.NoError(t, err)
		err = tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, newClientID)
		require.NoError(t, err)
		got, err := tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, launchID)
		require.NoError(t, err)
		require.Equal(t, verifiedClientID, got)
	})

	t.Run("should update a non exiting verified client id", func(t *testing.T) {
		verifiedClientID := types.VerifiedClientID{
			LaunchID:  3,
			ClientIDs: []string{"1"},
		}
		err := tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, verifiedClientID.LaunchID, verifiedClientID.ClientIDs[0])
		require.NoError(t, err)
		got, err := tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, verifiedClientID.LaunchID)
		require.NoError(t, err)
		require.Equal(t, verifiedClientID, got)
	})
}
