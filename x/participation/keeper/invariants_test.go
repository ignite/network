package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/participation/keeper"
	"github.com/ignite/network/x/participation/types"
)

func TestMismatchUsedAllocationsInvariant(t *testing.T) {
	var (
		ctx, tk, _        = testkeeper.NewTestSetup(t)
		addr              = sample.Address(r)
		auctionUsedAllocs = []types.AuctionUsedAllocations{
			{
				Address:        addr,
				AuctionId:      1,
				NumAllocations: sdkmath.OneInt(),
				Withdrawn:      false,
			},
			{
				Address:        addr,
				AuctionId:      2,
				NumAllocations: sdkmath.OneInt(),
				Withdrawn:      false,
			},
			{
				Address:        addr,
				AuctionId:      3,
				NumAllocations: sdkmath.NewInt(5),
				Withdrawn:      true,
			},
		}
		invalidUsedAllocs = types.UsedAllocations{
			Address:        addr,
			NumAllocations: sdkmath.NewInt(7),
		}
		validUsedAllocs = types.UsedAllocations{
			Address:        addr,
			NumAllocations: sdkmath.NewInt(2),
		}
	)
	byteAddr, err := tk.ParticipationKeeper.AddressCodec().StringToBytes(addr)
	require.NoError(t, err)
	accAddr := sdk.AccAddress(byteAddr)

	t.Run("should allow valid case", func(t *testing.T) {
		err := tk.ParticipationKeeper.UsedAllocations.Set(ctx, accAddr.String(), validUsedAllocs)
		require.NoError(t, err)
		for _, auction := range auctionUsedAllocs {
			err = tk.ParticipationKeeper.AuctionUsedAllocations.Set(ctx, collections.Join(accAddr, auction.AuctionId), auction)
			require.NoError(t, err)
		}
		_, isValid := keeper.MismatchUsedAllocationsInvariant(tk.ParticipationKeeper)(ctx)
		require.False(t, isValid)
	})

	t.Run("should prevent invalid case", func(t *testing.T) {
		err := tk.ParticipationKeeper.UsedAllocations.Set(ctx, accAddr.String(), invalidUsedAllocs)
		require.NoError(t, err)
		for _, auction := range auctionUsedAllocs {
			err = tk.ParticipationKeeper.AuctionUsedAllocations.Set(ctx, collections.Join(accAddr, auction.AuctionId), auction)
			require.NoError(t, err)
			require.NoError(t, err)
		}
		_, isValid := keeper.MismatchUsedAllocationsInvariant(tk.ParticipationKeeper)(ctx)
		require.True(t, isValid)
	})
}
