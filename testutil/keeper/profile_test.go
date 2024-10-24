package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
)

func TestTestMsgServers_CreateCoordinator(t *testing.T) {
	ctx, tk, tm := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	id, addr := tm.CreateCoordinator(ctx, r)
	coordinator, err := tk.ProfileKeeper.GetCoordinator(ctx, id)
	require.NoError(t, err)
	require.Equal(t, id, coordinator.CoordinatorId)
	require.Equal(t, addr.String(), coordinator.Address)
	coordinatorByAddress, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, addr)
	require.NoError(t, err)
	require.Equal(t, id, coordinatorByAddress.CoordinatorId)
	require.Equal(t, addr.String(), coordinatorByAddress.Address)
}
