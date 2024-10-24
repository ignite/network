package keeper

import (
	"context"
	"encoding/base64"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/launch/types"
)

// GetValidatorsAndTotalDelegation returns the genesisValidator map by
// consensus address and total of self delegation
func (k Keeper) GetValidatorsAndTotalDelegation(
	ctx context.Context,
	launchID uint64,
) (map[string]types.GenesisValidator, sdkmath.LegacyDec, error) {
	validators := make(map[string]types.GenesisValidator)
	totalDelegation := sdkmath.LegacyZeroDec()
	rng := collections.NewPrefixedPairRange[uint64, sdk.AccAddress](launchID)
	err := k.GenesisValidator.Walk(ctx, rng, func(key collections.Pair[uint64, sdk.AccAddress], val types.GenesisValidator) (bool, error) {
		consPubKey := base64.StdEncoding.EncodeToString(val.ConsPubKey)
		validators[consPubKey] = val
		totalDelegation = totalDelegation.Add(sdkmath.LegacyNewDecFromInt(val.SelfDelegation.Amount))
		return false, nil
	})
	return validators, totalDelegation, err
}

// AllGenesisAccount returns all GenesisAccount.
func (k Keeper) AllGenesisAccount(ctx context.Context) ([]types.GenesisAccount, error) {
	genesisAccounts := make([]types.GenesisAccount, 0)
	err := k.GenesisAccount.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], value types.GenesisAccount) (bool, error) {
		genesisAccounts = append(genesisAccounts, value)
		return false, nil
	})
	return genesisAccounts, err
}
