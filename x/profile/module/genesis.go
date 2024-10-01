package profile

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/profile/keeper"
	"github.com/ignite/network/x/profile/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) error {
	// Set all the validator
	for _, elem := range genState.ValidatorList {
		if err := k.Validator.Set(ctx, elem.Address, elem); err != nil {
			return err
		}
	}

	// Set all the validatorByOperatorAddress
	for _, elem := range genState.ValidatorsByOperatorAddress {
		if err := k.ValidatorByOperatorAddress.Set(ctx, elem.OperatorAddress, elem); err != nil {
			return err
		}
	}

	// Set all the coordinator
	for _, elem := range genState.CoordinatorList {
		if err := k.Coordinator.Set(ctx, elem.CoordinatorID, elem); err != nil {
			return err
		}
	}

	// Set coordinator count
	if err := k.CoordinatorSeq.Set(ctx, genState.CoordinatorCount); err != nil {
		return err
	}

	// Set all the coordinatorByAddress
	for _, elem := range genState.CoordinatorsByAddress {
		address, err := k.AddressCodec().StringToBytes(elem.Address)
		if err != nil {
			return err
		}
		if err := k.CoordinatorByAddress.Set(ctx, address, elem); err != nil {
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

	// Cordinators
	err = k.Coordinator.Walk(ctx, nil, func(key uint64, elem types.Coordinator) (bool, error) {
		genesis.CoordinatorList = append(genesis.CoordinatorList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.CoordinatorCount, err = k.CoordinatorSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	err = k.CoordinatorByAddress.Walk(ctx, nil, func(key sdk.AccAddress, elem types.CoordinatorByAddress) (bool, error) {
		genesis.CoordinatorsByAddress = append(genesis.CoordinatorsByAddress, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	// ListValidator
	if err := k.Validator.Walk(ctx, nil, func(_ string, val types.Validator) (stop bool, err error) {
		genesis.ValidatorList = append(genesis.ValidatorList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.ValidatorByOperatorAddress.Walk(ctx, nil, func(_ string, val types.ValidatorByOperatorAddress) (stop bool, err error) {
		genesis.ValidatorsByOperatorAddress = append(genesis.ValidatorsByOperatorAddress, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
