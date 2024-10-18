package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/keeper"
	"github.com/ignite/network/x/profile/types"
)

func TestCoordinatorAddrNotFoundInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should not break with valid state", func(t *testing.T) {
		var err error
		coordinator := sample.Coordinator(r, sample.Address(r))
		coordinator.CoordinatorID, err = tk.ProfileKeeper.AppendCoordinator(ctx, coordinator)
		require.NoError(t, err)
		acc := sample.AccAddress(r)
		err = tk.ProfileKeeper.CoordinatorByAddress.Set(ctx, acc, types.CoordinatorByAddress{
			Address:       acc.String(),
			CoordinatorID: coordinator.CoordinatorID,
		})
		require.NoError(t, err)
		msg, broken := keeper.CoordinatorAddrNotFoundInvariant(tk.ProfileKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should not break with coordinator not found from coordinator by address", func(t *testing.T) {
		acc := sample.AccAddress(r)
		err := tk.ProfileKeeper.CoordinatorByAddress.Set(ctx, acc, types.CoordinatorByAddress{
			Address:       acc.String(),
			CoordinatorID: 10,
		})
		require.NoError(t, err)
		msg, broken := keeper.CoordinatorAddrNotFoundInvariant(tk.ProfileKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
