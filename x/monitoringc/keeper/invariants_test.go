package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/monitoringc/keeper"
	"github.com/ignite/network/x/monitoringc/types"
)

func TestMissingVerifiedClientIDInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should allow valid case", func(t *testing.T) {
		n := sample.Uint64(r)
		launchID := sample.Uint64(r)
		for i := uint64(0); i < n; i++ {
			clientID := sample.AlphaString(r, 10)
			err := tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, clientID)
			require.NoError(t, err)
			err = tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(ctx, clientID, types.LaunchIDFromVerifiedClientID{
				ClientId: clientID,
				LaunchId: launchID,
			})
			require.NoError(t, err)
		}
		msg, broken := keeper.MissingVerifiedClientIDInvariant(tk.MonitoringConsumerKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should prevent invalid case", func(t *testing.T) {
		n := sample.Uint64(r)
		launchID := sample.Uint64(r)
		for i := uint64(0); i < n; i++ {
			clientID := sample.AlphaString(r, 10)
			err := tk.MonitoringConsumerKeeper.AddVerifiedClientID(ctx, launchID, clientID)
			require.NoError(t, err)
		}
		msg, broken := keeper.MissingVerifiedClientIDInvariant(tk.MonitoringConsumerKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
