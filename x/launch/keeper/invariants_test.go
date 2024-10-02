package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	launch "github.com/ignite/network/x/launch/types"
)

func TestDuplicatedAccountInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should not break with valid state", func(t *testing.T) {
		launchID := uint64(0)
		vestingAddress := sample.AccAddress(r)
		vestingAccount := sample.VestingAccount(r, launchID, vestingAddress.String())
		err := tk.LaunchKeeper.VestingAccount.Set(ctx, collections.Join(launchID, vestingAddress), vestingAccount)
		require.NoError(t, err)

		genesisAddress := sample.AccAddress(r)
		genesisAccount := sample.GenesisAccount(r, launchID, genesisAddress.String())
		err = tk.LaunchKeeper.GenesisAccount.Set(ctx, collections.Join(launchID, genesisAddress), genesisAccount)
		require.NoError(t, err)

		msg, broken := keeper.DuplicatedAccountInvariant(tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with duplicated account", func(t *testing.T) {
		launchID := uint64(0)
		addr := sample.AccAddress(r)
		err := tk.LaunchKeeper.VestingAccount.Set(ctx, collections.Join(launchID, addr), sample.VestingAccount(r, launchID, addr.String()))
		require.NoError(t, err)
		err = tk.LaunchKeeper.GenesisAccount.Set(ctx, collections.Join(launchID, addr), sample.GenesisAccount(r, launchID, addr.String()))
		require.NoError(t, err)

		msg, broken := keeper.DuplicatedAccountInvariant(tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestInvalidChainInvariant(t *testing.T) {
	t.Run("should not break with valid state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		chain := sample.Chain(r, 0, 0)
		project := sample.Project(r, 0)
		projectID, err := tk.ProjectKeeper.AppendProject(ctx, project)
		require.NoError(t, err)
		chain.ProjectID = projectID
		chain.HasProject = true
		_, err = tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		msg, broken := keeper.InvalidChainInvariant(tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with an invalid chain", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		chain := sample.Chain(r, 0, 0)
		chain.GenesisChainID = "_invalid_"
		_, err := tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		msg, broken := keeper.InvalidChainInvariant(tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should break with a chain that does not have a valid associated project", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		chain := sample.Chain(r, 0, 0)
		chain.HasProject = true
		chain.ProjectID = 1000
		_, err := tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		msg, broken := keeper.InvalidChainInvariant(tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestUnknownRequestTypeInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should not break with valid state", func(t *testing.T) {
		_, err := tk.LaunchKeeper.AppendRequest(ctx, sample.Request(r, 0, sample.Address(r)))
		require.NoError(t, err)
		msg, broken := keeper.UnknownRequestTypeInvariant(tk.LaunchKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should break with an invalid request", func(t *testing.T) {
		_, err := tk.LaunchKeeper.AppendRequest(ctx, sample.RequestWithContent(r, 0,
			launch.NewGenesisAccount(0, sample.Address(r), sdk.NewCoins()),
		))
		require.NoError(t, err)
		msg, broken := keeper.UnknownRequestTypeInvariant(tk.LaunchKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
