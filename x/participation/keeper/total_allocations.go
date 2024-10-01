package keeper

import (
	"context"
	"math"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/participation/types"
)

// GetTotalAllocations returns the number of available allocations based on delegations
func (k Keeper) GetTotalAllocations(ctx context.Context, address string) (sdkmath.Int, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return sdkmath.ZeroInt(), ignterrors.Critical("failed to get participation params")
	}

	allocationPriceBondedDec := sdkmath.LegacyNewDecFromInt(params.AllocationPrice.Bonded)

	accAddr, err := k.addressCodec.StringToBytes(address)
	if err != nil {
		return sdkmath.ZeroInt(), sdkerrors.Wrapf(types.ErrInvalidAddress, err.Error())
	}

	// count total shares for account
	totalDel := sdkmath.LegacyZeroDec()
	dels, err := k.stakingKeeper.GetDelegatorDelegations(ctx, accAddr, math.MaxUint16)
	if err != nil {
		return sdkmath.ZeroInt(), sdkerrors.Wrapf(types.ErrInvalidDelegations, err.Error())
	}

	for _, del := range dels {
		totalDel = totalDel.Add(del.GetShares())
	}

	numAlloc := totalDel.Quo(allocationPriceBondedDec)
	if numAlloc.IsNegative() {
		return sdkmath.ZeroInt(), types.ErrInvalidAllocationAmount
	}

	return numAlloc.TruncateInt(), nil
}
