package reward

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/reward/keeper"
	"github.com/ignite/network/x/reward/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) error {
	// Set all the rewardPool
	for _, elem := range genState.RewardPoolList {
		if err := k.RewardPool.Set(ctx, elem.LaunchId, elem); err != nil {
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

	if err := k.RewardPool.Walk(ctx, nil, func(_ uint64, val types.RewardPool) (stop bool, err error) {
		genesis.RewardPoolList = append(genesis.RewardPoolList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
