package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/launch/types"
)

// AllVestingAccount returns all VestingAccount.
func (k Keeper) AllVestingAccount(ctx context.Context) ([]types.VestingAccount, error) {
	vestingAccount := make([]types.VestingAccount, 0)
	err := k.VestingAccount.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], value types.VestingAccount) (bool, error) {
		vestingAccount = append(vestingAccount, value)
		return false, nil
	})
	return vestingAccount, err
}
