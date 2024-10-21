package keeper

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ignite/network/testutil/sample"
)

// DelegateN creates N delegations from the same address
func (tk TestKeepers) DelegateN(ctx sdk.Context, r *rand.Rand, address string, shareAmt int64, n int) ([]stakingtypes.Delegation, sdkmath.LegacyDec, error) {
	items := make([]stakingtypes.Delegation, n)
	totalShares := sdkmath.LegacyZeroDec()

	for i := range items {
		item, err := tk.Delegate(ctx, r, address, shareAmt)
		if err != nil {
			return nil, totalShares, err
		}

		items[i] = item
		totalShares = totalShares.Add(items[i].Shares)
	}

	return items, totalShares, nil
}

// Delegate creates a sample delegation and sets it in the keeper
func (tk TestKeepers) Delegate(ctx sdk.Context, r *rand.Rand, address string, amt int64) (stakingtypes.Delegation, error) {
	del := sample.Delegation(tk.T, r, address)
	del.Shares = sdkmath.LegacyNewDec(amt)
	return del, tk.StakingKeeper.SetDelegation(ctx, del)
}
