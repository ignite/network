package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
)

func TestIsProjectMainnetLaunchTriggered(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should show project has mainnet with launch triggered", func(t *testing.T) {
		var err error
		projectMainnetLaunched := sample.Project(r, 0)
		projectMainnetLaunched.MainnetInitialized = true
		chainLaunched := sample.Chain(r, 0, 0)
		chainLaunched.LaunchTriggered = true
		chainLaunched.IsMainnet = true
		projectMainnetLaunched.MainnetId, err = tk.LaunchKeeper.AppendChain(ctx, chainLaunched)
		require.NoError(t, err)
		projectMainnetLaunched.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, projectMainnetLaunched)
		require.NoError(t, err)
		res, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetLaunched.ProjectId)
		require.NoError(t, err)
		require.True(t, res)
	})

	t.Run("should show project has mainnet with launch not triggered", func(t *testing.T) {
		var err error
		projectMainnetInitialized := sample.Project(r, 1)
		projectMainnetInitialized.MainnetInitialized = true
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTriggered = false
		chain.IsMainnet = true
		projectMainnetInitialized.MainnetId, err = tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		projectMainnetInitialized.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, projectMainnetInitialized)
		require.NoError(t, err)
		res, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetInitialized.ProjectId)
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("should show project with mainnnet not initialized", func(t *testing.T) {
		var err error
		projectMainnetNotInitialized := sample.Project(r, 2)
		projectMainnetNotInitialized.MainnetInitialized = false
		projectMainnetNotInitialized.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, projectMainnetNotInitialized)
		require.NoError(t, err)
		res, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetNotInitialized.ProjectId)
		require.NoError(t, err)
		require.False(t, res)
	})

	t.Run("should show mainnet not found", func(t *testing.T) {
		var err error
		projectMainnetNotFound := sample.Project(r, 3)
		projectMainnetNotFound.MainnetInitialized = true
		projectMainnetNotFound.MainnetId = 1000
		projectMainnetNotFound.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, projectMainnetNotFound)
		require.NoError(t, err)
		_, err = tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectMainnetNotFound.ProjectId)
		require.Error(t, err)
	})

	t.Run("should show associated mainnet chain is not mainnet", func(t *testing.T) {
		var err error
		projectInvalidMainnet := sample.Project(r, 4)
		projectInvalidMainnet.MainnetInitialized = true
		chainNoMainnet := sample.Chain(r, 0, 0)
		chainNoMainnet.LaunchTriggered = false
		chainNoMainnet.IsMainnet = false
		projectInvalidMainnet.MainnetId, err = tk.LaunchKeeper.AppendChain(ctx, chainNoMainnet)
		require.NoError(t, err)
		projectInvalidMainnet.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, projectInvalidMainnet)
		require.NoError(t, err)
		_, err = tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, projectInvalidMainnet.ProjectId)
		require.Error(t, err)
	})

	t.Run("should show project not found", func(t *testing.T) {
		_, err := tk.ProjectKeeper.IsProjectMainnetLaunchTriggered(ctx, 1000)
		require.Error(t, err)
	})
}
