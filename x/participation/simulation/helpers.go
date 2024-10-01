package simulation

import (
	"math/rand"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"

	"github.com/ignite/network/x/participation/keeper"
	"github.com/ignite/network/x/participation/types"
)

// RandomAuctionWithdrawEnabled returns random auction where used allocations can be withdrawn at blockTime
func RandomAuctionWithdrawEnabled(
	ctx sdk.Context,
	r *rand.Rand,
	fk types.FundraisingKeeper,
	k keeper.Keeper,
) (auction fundraisingtypes.AuctionI, found bool) {
	auctions, err := fk.Auctions(ctx)
	if err != nil || len(auctions) == 0 {
		return auction, false
	}
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}
	if len(auctions) == 0 {
		return auction, false
	}

	r.Shuffle(len(auctions), func(i, j int) {
		auctions[i], auctions[j] = auctions[j], auctions[i]
	})

	for _, a := range auctions {
		// if auction cancelled, withdraw is always enabled
		if a.GetStatus() == fundraisingtypes.AuctionStatusCancelled {
			return a, true
		}

		// check if withdrawal delay has passed and hence withdraw is enabled
		if ctx.BlockTime().After(a.GetStartTime().Add(params.WithdrawalDelay)) {
			return a, true
		}
	}

	return auction, false
}

// RandomAccWithAvailableAllocations returns random account that has at least the desired amount of available allocations
// and can still participate in the specified auction
func RandomAccWithAvailableAllocations(ctx sdk.Context, r *rand.Rand,
	k keeper.Keeper,
	accs []simtypes.Account,
	desired sdkmath.Int,
	auctionID uint64,
) (simtypes.Account, sdkmath.Int, bool) {
	// Randomize the set
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	// account must have allocations but not already have participated
	for _, acc := range accs {
		amt, err := k.GetAvailableAllocations(ctx, acc.Address.String())
		if err != nil {
			continue
		}

		if amt.GTE(desired) {
			_, err := k.AuctionUsedAllocations.Get(ctx, collections.Join(acc.Address, auctionID))
			if err != nil {
				continue
			}

			return acc, amt, true
		}
	}

	return simtypes.Account{}, sdkmath.ZeroInt(), false
}

// RandomAccWithAuctionUsedAllocationsNotWithdrawn returns random account that has used allocations for the given
// auction that have not yet been withdrawn
func RandomAccWithAuctionUsedAllocationsNotWithdrawn(
	ctx sdk.Context,
	r *rand.Rand,
	k keeper.Keeper,
	accs []simtypes.Account,
	auctionID uint64,
) (simtypes.Account, bool) {
	// Randomize the set
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	// account must have used allocations for this auction that have not yet been withdrawn
	for _, acc := range accs {
		usedAllocations, err := k.AuctionUsedAllocations.Get(ctx, collections.Join(acc.Address, auctionID))
		if err != nil || usedAllocations.Withdrawn {
			continue
		}

		return acc, true
	}

	return simtypes.Account{}, false
}

func RandomTierFromList(r *rand.Rand, tierList []types.Tier) (types.Tier, bool) {
	if len(tierList) == 0 {
		return types.Tier{}, false
	}

	index := r.Intn(len(tierList))
	return tierList[index], true
}

// RandomAuctionParticipationEnabled returns random auction where participation is enabled
func RandomAuctionParticipationEnabled(
	ctx sdk.Context,
	r *rand.Rand,
	fk types.FundraisingKeeper,
	k keeper.Keeper,
) (auction fundraisingtypes.AuctionI, found bool) {
	auctions, err := fk.Auctions(ctx)
	if err != nil || len(auctions) == 0 {
		return auction, false
	}

	r.Shuffle(len(auctions), func(i, j int) {
		auctions[i], auctions[j] = auctions[j], auctions[i]
	})

	for _, a := range auctions {
		if a.GetStatus() != fundraisingtypes.AuctionStatusStandBy {
			continue
		}
		ok, err := k.IsRegistrationEnabled(ctx, a.GetStartTime())
		if err != nil || !ok {
			continue
		}
		auction = a
		found = true
	}

	return
}
