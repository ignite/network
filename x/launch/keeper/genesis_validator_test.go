package keeper_test

import (
	"context"
	"encoding/base64"
	"testing"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

func createNGenesisValidatorByLaunchID(keeper *keeper.Keeper, ctx context.Context, launchID int) []types.GenesisValidator {
	items := make([]types.GenesisValidator, launchID)
	for i := range items {
		addr := sample.AccAddress(r)
		items[i] = sample.GenesisValidator(r, uint64(launchID), addr.String())
		_ = keeper.GenesisValidator.Set(ctx, collections.Join(items[i].LaunchId, addr), items[i])
	}
	return items
}

func TestKeeper_GetValidatorsAndTotalDelegation(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	launchID := 10
	validators := createNGenesisValidatorByLaunchID(tk.LaunchKeeper, ctx, launchID)
	totalSelfDelegation := sdkmath.LegacyZeroDec()
	validatorMap := make(map[string]types.GenesisValidator)
	for _, validator := range validators {
		consPubKey := base64.StdEncoding.EncodeToString(validator.ConsPubKey)
		validatorMap[consPubKey] = validator
		totalSelfDelegation = totalSelfDelegation.Add(sdkmath.LegacyNewDecFromInt(validator.SelfDelegation.Amount))
	}

	t.Run("should get a map of genesis validator and the total delegation", func(t *testing.T) {
		val, got, err := tk.LaunchKeeper.GetValidatorsAndTotalDelegation(ctx, uint64(launchID))
		require.NoError(t, err)
		require.Equal(t, totalSelfDelegation, got)
		require.EqualValues(t, validatorMap, val)
	})
}
