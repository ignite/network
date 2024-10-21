package launch

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

// InitGenesis initializes the launch module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genState types.GenesisState) error {
	// Set all the chain
	for _, elem := range genState.ChainList {
		if err := k.Chain.Set(ctx, elem.LaunchId, elem); err != nil {
			return err
		}
	}

	// Set chain count
	if err := k.ChainSeq.Set(ctx, genState.ChainCount); err != nil {
		return err
	}

	// Set all the genesisAccount
	for _, elem := range genState.GenesisAccountList {
		address, err := k.AddressCodec().StringToBytes(elem.Address)
		if err != nil {
			return err
		}
		if err := k.GenesisAccount.Set(ctx, collections.Join(elem.LaunchId, sdk.AccAddress(address)), elem); err != nil {
			return err
		}
	}

	// Set all the vestingAccount
	for _, elem := range genState.VestingAccountList {
		address, err := k.AddressCodec().StringToBytes(elem.Address)
		if err != nil {
			return err
		}
		if err := k.VestingAccount.Set(ctx, collections.Join(elem.LaunchId, sdk.AccAddress(address)), elem); err != nil {
			return err
		}
	}

	// Set all the genesisValidator
	for _, elem := range genState.GenesisValidatorList {
		address, err := k.AddressCodec().StringToBytes(elem.Address)
		if err != nil {
			return err
		}
		if err := k.GenesisValidator.Set(ctx, collections.Join(elem.LaunchId, sdk.AccAddress(address)), elem); err != nil {
			return err
		}
	}

	// Set all the paramChange
	for _, elem := range genState.ParamChangeList {
		if err := k.ParamChange.Set(ctx, collections.Join(elem.LaunchId, types.ParamChangeSubKey(elem.Module, elem.Param)), elem); err != nil {
			return err
		}
	}

	// Set all the request
	for _, elem := range genState.RequestList {
		if err := k.Request.Set(ctx, collections.Join(elem.LaunchId, elem.RequestId), elem); err != nil {
			return err
		}
	}

	// Set all request counter
	for _, elem := range genState.RequestCounters {
		if err := k.RequestSeq.Set(ctx, elem.LaunchId, elem.Counter); err != nil {
			return err
		}
	}

	// this line is used by starport scaffolding # genesis/module/init

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	err = k.Chain.Walk(ctx, nil, func(key uint64, elem types.Chain) (bool, error) {
		genesis.ChainList = append(genesis.ChainList, elem)
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	genesis.ChainCount, err = k.ChainSeq.Peek(ctx)
	if err != nil {
		return nil, err
	}

	if err := k.GenesisAccount.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], val types.GenesisAccount) (stop bool, err error) {
		genesis.GenesisAccountList = append(genesis.GenesisAccountList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.VestingAccount.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], val types.VestingAccount) (stop bool, err error) {
		genesis.VestingAccountList = append(genesis.VestingAccountList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.GenesisValidator.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], val types.GenesisValidator) (stop bool, err error) {
		genesis.GenesisValidatorList = append(genesis.GenesisValidatorList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.ParamChange.Walk(ctx, nil, func(_ collections.Pair[uint64, string], val types.ParamChange) (stop bool, err error) {
		genesis.ParamChangeList = append(genesis.ParamChangeList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.Request.Walk(ctx, nil, func(_ collections.Pair[uint64, uint64], val types.Request) (stop bool, err error) {
		genesis.RequestList = append(genesis.RequestList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	// Get request counts
	for _, elem := range genesis.ChainList {
		// Get request count
		counter, err := k.GetRequestCounter(ctx, elem.LaunchId)
		if err != nil {
			return nil, err
		}
		genesis.RequestCounters = append(genesis.RequestCounters, types.RequestCounter{
			LaunchId: elem.LaunchId,
			Counter:  counter,
		})
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
