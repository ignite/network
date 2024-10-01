package project

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) error {
	// Set all the mainnetAccount
	for _, elem := range genState.MainnetAccountList {
		address, err := k.AddressCodec().StringToBytes(elem.Address)
		if err != nil {
			return err
		}

		if err := k.MainnetAccount.Set(ctx, collections.Join(elem.ProjectID, sdk.AccAddress(address)), elem); err != nil {
			return err
		}
	}

	// Set all the project
	for _, elem := range genState.ProjectList {
		if err := k.Project.Set(ctx, elem.ProjectID, elem); err != nil {
			return err
		}
	}

	// Set project count
	if err := k.ProjectSeq.Set(ctx, genState.ProjectCount); err != nil {
		return err
	}
	// Set all the projectChains
	for _, elem := range genState.ProjectChainsList {
		if err := k.ProjectChains.Set(ctx, elem.ProjectID, elem); err != nil {
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

	if err := k.MainnetAccount.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], val types.MainnetAccount) (stop bool, err error) {
		genesis.MainnetAccountList = append(genesis.MainnetAccountList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	err = k.Project.Walk(ctx, nil, func(key uint64, elem types.Project) (bool, error) {
		genesis.ProjectList = append(genesis.ProjectList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.ProjectCount, err = k.ProjectSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	if err := k.ProjectChains.Walk(ctx, nil, func(_ uint64, val types.ProjectChains) (stop bool, err error) {
		genesis.ProjectChainsList = append(genesis.ProjectChainsList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
