package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	launchtypes "github.com/ignite/network/x/launch/types"
	"github.com/ignite/network/x/reward/keeper"
	"github.com/ignite/network/x/reward/types"
)

// FindRandomChainWithCoordBalance find a random chain from store
func FindRandomChainWithCoordBalance(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	bk types.BankKeeper,
	hasRewardPool,
	checkBalance bool,
	wantCoins sdk.Coins,
) (chain launchtypes.Chain, found bool) {
	chains, err := k.GetLaunchKeeper().Chains(ctx)
	if err != nil {
		return launchtypes.Chain{}, false
	}
	r.Shuffle(len(chains), func(i, j int) {
		chains[i], chains[j] = chains[j], chains[i]
	})
	for _, c := range chains {
		_, err := k.RewardPool.Get(ctx, c.LaunchID)
		if hasRewardPool != (err != nil) {
			continue
		}

		// chain cannot be launch triggered
		if c.LaunchTriggered || c.IsMainnet {
			continue
		}

		// check if the coordinator is still in the store and active
		coordinator, err := k.GetProfileKeeper().GetCoordinator(ctx, c.CoordinatorID)
		if err != nil || !coordinator.Active {
			continue
		}

		coordinatorAccAddr, err := k.AddressCodec().StringToBytes(coordinator.Address)
		if err != nil {
			continue
		}

		if checkBalance {
			balance := bk.SpendableCoins(ctx, coordinatorAccAddr)
			if !balance.IsAllGTE(wantCoins) {
				continue
			}
		}

		chain = c
		found = true
		break
	}
	return chain, found
}
