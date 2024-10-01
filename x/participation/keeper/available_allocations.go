package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	"github.com/pkg/errors"
)

// GetAvailableAllocations returns the number of allocations that are unused
func (k Keeper) GetAvailableAllocations(ctx context.Context, address string) (sdkmath.Int, error) {
	numTotalAlloc, err := k.GetTotalAllocations(ctx, address)
	if err != nil {
		return sdkmath.ZeroInt(), err
	}

	usedAlloc, err := k.UsedAllocations.Get(ctx, address)
	if errors.Is(err, collections.ErrNotFound) {
		return numTotalAlloc, nil
	} else if err != nil {
		return sdkmath.ZeroInt(), err
	}

	// return 0 if result would be negative
	if usedAlloc.NumAllocations.GT(numTotalAlloc) {
		return sdkmath.ZeroInt(), nil
	}

	availableAlloc := numTotalAlloc.Sub(usedAlloc.NumAllocations)

	return availableAlloc, nil
}
