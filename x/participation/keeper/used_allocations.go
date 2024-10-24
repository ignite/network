package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/participation/types"
)

// AllUsedAllocations returns all UsedAllocations.
func (k Keeper) AllUsedAllocations(ctx context.Context) ([]types.UsedAllocations, error) {
	usedAllocations := make([]types.UsedAllocations, 0)
	err := k.UsedAllocations.Walk(ctx, nil, func(_ string, value types.UsedAllocations) (bool, error) {
		usedAllocations = append(usedAllocations, value)
		return false, nil
	})
	return usedAllocations, err
}

// AllAuctionUsedAllocations returns all AuctionUsedAllocations.
func (k Keeper) AllAuctionUsedAllocations(ctx context.Context, address string) ([]types.AuctionUsedAllocations, error) {
	accAddress, err := k.addressCodec.StringToBytes(address)
	if err != nil {
		return nil, err
	}
	auctionUsedAllocations := make([]types.AuctionUsedAllocations, 0)
	rng := collections.NewPrefixedPairRange[sdk.AccAddress, uint64](accAddress)
	err = k.AuctionUsedAllocations.Walk(ctx, rng, func(_ collections.Pair[sdk.AccAddress, uint64], value types.AuctionUsedAllocations) (bool, error) {
		auctionUsedAllocations = append(auctionUsedAllocations, value)
		return false, nil
	})
	return auctionUsedAllocations, err
}
