package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
)

func TestKeeper_AddChainToProject(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should fail if project does not exist", func(t *testing.T) {
		err := tk.ProjectKeeper.AddChainToProject(ctx, 0, 0)
		require.Error(t, err)
	})

	// Chains can be added
	projectID := uint64(0)
	t.Run("should allow adding chains to project", func(t *testing.T) {
		err := tk.ProjectKeeper.Project.Set(ctx, projectID, sample.Project(r, 0))
		require.NoError(t, err)
		err = tk.ProjectKeeper.AddChainToProject(ctx, projectID, 0)
		require.NoError(t, err)
		err = tk.ProjectKeeper.AddChainToProject(ctx, projectID, 1)
		require.NoError(t, err)
		err = tk.ProjectKeeper.AddChainToProject(ctx, projectID, 2)
		require.NoError(t, err)

		projectChains, err := tk.ProjectKeeper.GetProjectChains(ctx, projectID)
		require.NoError(t, err)
		require.EqualValues(t, projectChains.ProjectId, uint64(0))
		require.Len(t, projectChains.Chains, 3)
		require.EqualValues(t, []uint64{0, 1, 2}, projectChains.Chains)
	})

	// Can't add an existing chain
	t.Run("should prevent adding existing chain to project", func(t *testing.T) {
		err := tk.ProjectKeeper.AddChainToProject(ctx, projectID, 0)
		require.Error(t, err)
	})
}
