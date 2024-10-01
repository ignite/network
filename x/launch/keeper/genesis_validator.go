package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/launch/types"
)

// AllGenesisValidator returns all GenesisValidator.
func (k Keeper) AllGenesisValidator(ctx context.Context) ([]types.GenesisValidator, error) {
	genesisValidators := make([]types.GenesisValidator, 0)
	err := k.GenesisValidator.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], value types.GenesisValidator) (bool, error) {
		genesisValidators = append(genesisValidators, value)
		return false, nil
	})
	return genesisValidators, err
}
