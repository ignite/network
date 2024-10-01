package participation

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/participation/keeper"
	"github.com/ignite/network/x/participation/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) error {
	// Set all the auctionUsedAllocations
	for _, elem := range genState.AuctionUsedAllocationsList {
		address, err := k.AddressCodec().StringToBytes(elem.Address)
		if err != nil {
			return err
		}
		if err := k.AuctionUsedAllocations.Set(ctx, collections.Join(sdk.AccAddress(address), elem.AuctionID), elem); err != nil {
			return err
		}
	}
	// Set all the usedAllocations
	for _, elem := range genState.UsedAllocationsList {
		if err := k.UsedAllocations.Set(ctx, elem.Address, elem); err != nil {
			return err
		}
	}

	// this line is used by starport scaffolding # genesis/module/init

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	if err := k.AuctionUsedAllocations.Walk(ctx, nil, func(_ collections.Pair[sdk.AccAddress, uint64], val types.AuctionUsedAllocations) (stop bool, err error) {
		genesis.AuctionUsedAllocationsList = append(genesis.AuctionUsedAllocationsList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.UsedAllocations.Walk(ctx, nil, func(_ string, val types.UsedAllocations) (stop bool, err error) {
		genesis.UsedAllocationsList = append(genesis.UsedAllocationsList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
