package keeper_test

import (
	"fmt"
	"testing"

	"cosmossdk.io/collections"
	"github.com/stretchr/testify/require"

	tc "github.com/ignite/network/testutil/constructor"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func TestAccountWithoutProjectInvariant(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	t.Run("should allow valid case", func(t *testing.T) {
		var err error
		project := sample.Project(r, 0)
		project.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, project)
		require.NoError(t, err)

		addr := sample.AccAddress(r)
		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(project.ProjectId, addr), sample.MainnetAccount(r, project.ProjectId, addr.String()))
		require.NoError(t, err)

		msg, broken := keeper.AccountWithoutProjectInvariant(tk.ProjectKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should prevent invalid case", func(t *testing.T) {
		addr := sample.AccAddress(r)
		projectID := uint64(100)
		err := tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(projectID, addr), sample.MainnetAccount(r, projectID, addr.String()))
		require.NoError(t, err)

		msg, broken := keeper.AccountWithoutProjectInvariant(tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})
}

func TestProjectSharesInvariant(t *testing.T) {
	t.Run("should allow valid case", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		// create projects with some allocated shares
		projectID1, projectID2 := uint64(1), uint64(2)
		project := sample.Project(r, projectID1)
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			tc.Shares(t, "100foo,200bar"),
		)
		err := tk.ProjectKeeper.Project.Set(ctx, projectID1, project)
		require.NoError(t, err)

		project = sample.Project(r, projectID2)
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			tc.Shares(t, "10000foo"),
		)
		err = tk.ProjectKeeper.Project.Set(ctx, projectID2, project)
		require.NoError(t, err)

		// mint vouchers
		voucherFoo, voucherBar := types.VoucherDenom(projectID1, "foo"), types.VoucherDenom(projectID1, "bar")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("50%s,100%s", voucherFoo, voucherBar)))

		// mint vouchers for another project
		voucherFoo = types.VoucherDenom(projectID2, "foo")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("5000%s", voucherFoo)))

		// add accounts with shares
		addr1 := sample.AccAddress(r)
		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(projectID1, addr1), types.MainnetAccount{
			ProjectId: projectID1,
			Address:   addr1.String(),
			Shares:    tc.Shares(t, "20foo,40bar"),
		})
		require.NoError(t, err)

		addr2 := sample.AccAddress(r)
		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(projectID1, addr2), types.MainnetAccount{
			ProjectId: projectID1,
			Address:   addr2.String(),
			Shares:    tc.Shares(t, "30foo,60bar"),
		})
		require.NoError(t, err)

		addr3 := sample.AccAddress(r)
		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(projectID2, addr3), types.MainnetAccount{
			ProjectId: projectID2,
			Address:   addr3.String(),
			Shares:    tc.Shares(t, "5000foo"),
		})
		require.NoError(t, err)

		msg, broken := keeper.ProjectSharesInvariant(tk.ProjectKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should allow project with empty allocated share is valid", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		projectID := uint64(3)
		err := tk.ProjectKeeper.Project.Set(ctx, projectID, sample.Project(r, projectID))
		require.NoError(t, err)

		msg, broken := keeper.ProjectSharesInvariant(tk.ProjectKeeper)(ctx)
		require.False(t, broken, msg)
	})

	t.Run("should prevent allocated shares cannot be converted to vouchers", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		projectID := uint64(4)
		project := sample.Project(r, projectID)
		coins := tc.Coins(t, "100foo,200bar")
		shares := make(types.Shares, len(coins))
		for i, coin := range coins {
			shares[i] = coin
		}
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			shares,
		)
		err := tk.ProjectKeeper.Project.Set(ctx, projectID, project)
		require.NoError(t, err)

		msg, broken := keeper.ProjectSharesInvariant(tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should prevent invalid allocated shares", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		projectID := uint64(4)
		project := sample.Project(r, projectID)
		project.AllocatedShares = types.IncreaseShares(
			project.AllocatedShares,
			tc.Shares(t, "100foo,200bar"),
		)
		err := tk.ProjectKeeper.Project.Set(ctx, projectID, project)
		require.NoError(t, err)

		// mint vouchers
		voucherFoo, voucherBar := types.VoucherDenom(projectID, "foo"), types.VoucherDenom(projectID, "bar")
		tk.Mint(ctx, sample.Address(r), tc.Coins(t, fmt.Sprintf("99%s,200%s", voucherFoo, voucherBar)))

		msg, broken := keeper.ProjectSharesInvariant(tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})

	t.Run("should prevent project with special allocations not tracked by allocated shares", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetup(t)
		project := sample.Project(r, 3)
		project.SpecialAllocations.GenesisDistribution = types.IncreaseShares(
			project.SpecialAllocations.GenesisDistribution,
			sample.Shares(r),
		)
		err := tk.ProjectKeeper.Project.Set(ctx, project.ProjectId, project)
		require.NoError(t, err)

		msg, broken := keeper.ProjectSharesInvariant(tk.ProjectKeeper)(ctx)
		require.True(t, broken, msg)
	})
}
