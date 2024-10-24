package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/reward/types"
)

// ListRewardPool returns all rewardPool
func (k Keeper) ListRewardPool(ctx sdk.Context) ([]types.RewardPool, error) {
	rewardPools := make([]types.RewardPool, 0)
	err := k.RewardPool.Walk(ctx, nil, func(_ uint64, rewardPool types.RewardPool) (bool, error) {
		rewardPools = append(rewardPools, rewardPool)
		return false, nil
	})
	return rewardPools, err
}
