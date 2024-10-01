package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/participation/types"
)

const (
	mismatchUsedAllocationsRoute = "mismatch-used-allocations"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, mismatchUsedAllocationsRoute,
		MismatchUsedAllocationsInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return MismatchUsedAllocationsInvariant(k)(ctx)
	}
}

// MismatchUsedAllocationsInvariant invariant that checks if the number of used allocations in `UsedAllocations`
// is different from the sum of per-auction used allocations in `AuctionUsedAllocations`
func MismatchUsedAllocationsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all, err := k.AllUsedAllocations(ctx)
		if err != nil {
			return "", false
		}
		for _, usedAllocs := range all {
			auctionUsedAllocs, err := k.AllAuctionUsedAllocations(ctx, usedAllocs.Address)
			if err != nil {
				return "", false
			}
			sum := sdkmath.ZeroInt()
			for _, auction := range auctionUsedAllocs {
				if !auction.Withdrawn {
					sum = sum.Add(auction.NumAllocations)
				}
			}
			if !sum.Equal(usedAllocs.NumAllocations) {
				return sdk.FormatInvariant(
					types.ModuleName, mismatchUsedAllocationsRoute,
					"total used allocations not equal to sum of per-auction used allocations",
				), true
			}
		}
		return "", false
	}
}
