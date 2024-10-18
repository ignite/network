package simulation_test

import (
	"fmt"
	"math/rand"
	"testing"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profilekeeper "github.com/ignite/network/x/profile/keeper"
	profiletypes "github.com/ignite/network/x/profile/types"
	simproject "github.com/ignite/network/x/project/simulation"
	projecttypes "github.com/ignite/network/x/project/types"
)

// populateCoordinators populates the profile keeper with some coordinators from simulation accounts
func populateCoordinators(
	t *testing.T,
	r *rand.Rand,
	ctx sdk.Context,
	pk profilekeeper.Keeper,
	accs []simtypes.Account,
	coordNb int,
) (coordIDs []uint64) {
	require.LessOrEqual(t, coordNb, len(accs))
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})
	for i := 0; i < coordNb; i++ {
		coordID, err := pk.AppendCoordinator(ctx, profiletypes.Coordinator{
			Address:     accs[i].Address.String(),
			Description: sample.CoordinatorDescription(r),
			Active:      true,
		})
		require.NoError(t, err)
		coordIDs = append(coordIDs, coordID)
	}

	return
}

func TestGetCoordSimAccount(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	t.Run("should return no coordinator", func(t *testing.T) {
		_, _, found := simproject.GetCoordSimAccount(r, ctx, tk.ProfileKeeper, accs)
		require.False(t, found)
	})

	populateCoordinators(t, r, ctx, tk.ProfileKeeper, accs, 10)

	t.Run("should find coordinators", func(t *testing.T) {
		acc, coordID, found := simproject.GetCoordSimAccount(r, ctx, tk.ProfileKeeper, accs)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, err := tk.ProfileKeeper.GetCoordinator(ctx, coordID)
		require.NoError(t, err)
	})
}

func TestGetCoordSimAccountWithProjectID(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()
	var err error

	t.Run("should find no project", func(t *testing.T) {
		_, _, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			tk.ProjectKeeper,
			accs,
			false,
			false,
		)
		require.False(t, found)
	})

	coords := populateCoordinators(t, r, ctx, tk.ProfileKeeper, accs, 10)

	t.Run("should find one project with mainnet launch triggered", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			0,
			sample.AlphaString(r, 5),
			coords[1],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		prjt.MainnetInitialized = true
		chain := sample.Chain(r, 0, coords[1])
		chain.LaunchTriggered = true
		chain.IsMainnet = true
		prjt.MainnetID, err = tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		_, err = tk.ProjectKeeper.AppendProject(ctx, prjt)
		require.NoError(t, err)
		_, _, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			tk.ProjectKeeper,
			accs,
			false,
			true,
		)
		require.False(t, found)
	})

	t.Run("should find a project", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			1,
			sample.AlphaString(r, 5),
			coords[0],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		prjt.MainnetInitialized = true
		chain := sample.Chain(r, 0, coords[1])
		chain.LaunchTriggered = false
		chain.IsMainnet = true
		prjt.MainnetID, err = tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		_, err = tk.ProjectKeeper.AppendProject(ctx, prjt)
		require.NoError(t, err)
		acc, id, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			tk.ProjectKeeper,
			accs,
			false,
			true,
		)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, err = tk.ProjectKeeper.GetProject(ctx, id)
		require.NoError(t, err)
		require.EqualValues(t, id, prjt.ProjectID)
	})

	t.Run("should find a project with no mainnet initialized", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			2,
			sample.AlphaString(r, 5),
			coords[1],
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		idNoMainnet, err := tk.ProjectKeeper.AppendProject(ctx, prjt)
		require.NoError(t, err)
		acc, id, found := simproject.GetCoordSimAccountWithProjectID(
			r,
			ctx,
			tk.ProfileKeeper,
			tk.ProjectKeeper,
			accs,
			true,
			false,
		)
		require.True(t, found)
		require.Contains(t, accs, acc)
		_, err = tk.ProjectKeeper.GetProject(ctx, id)
		require.NoError(t, err)
		require.EqualValues(t, idNoMainnet, id)
		require.EqualValues(t, prjt.ProjectID, id)
		require.False(t, prjt.MainnetInitialized)
	})
}

func TestGetSharesFromProject(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()

	t.Run("should find no project", func(t *testing.T) {
		_, found := simproject.GetSharesFromProject(r, ctx, tk.ProjectKeeper, 0)
		require.False(t, found)
	})

	t.Run("should find no shares remaining for the project", func(t *testing.T) {
		prjt := projecttypes.NewProject(
			0,
			sample.AlphaString(r, 5),
			0,
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		)
		shares, err := projecttypes.NewShares(fmt.Sprintf(
			"%[1]dfoo,%[1]dbar,%[1]dtoto",
			networktypes.TotalShareNumber,
		))
		require.NoError(t, err)
		prjt.AllocatedShares = shares
		prjtSharesReached, err := tk.ProjectKeeper.AppendProject(ctx, prjt)
		require.NoError(t, err)

		_, found := simproject.GetSharesFromProject(r, ctx, tk.ProjectKeeper, prjtSharesReached)
		require.False(t, found)
	})

	t.Run("should find project with available shares", func(t *testing.T) {
		prjtID, err := tk.ProjectKeeper.AppendProject(ctx, projecttypes.NewProject(
			1,
			sample.AlphaString(r, 5),
			0,
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
			sample.Duration(r).Milliseconds(),
		))
		require.NoError(t, err)
		shares, found := simproject.GetSharesFromProject(r, ctx, tk.ProjectKeeper, prjtID)
		require.True(t, found)
		require.NotEqualValues(t, projecttypes.EmptyShares(), shares)
	})
}

func TestGetAccountWithVouchers(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()
	var err error

	mint := func(addr sdk.AccAddress, coins sdk.Coins) {
		require.NoError(t, tk.BankKeeper.MintCoins(ctx, projecttypes.ModuleName, coins))
		require.NoError(t, tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, projecttypes.ModuleName, addr, coins))
	}

	t.Run("should find no account", func(t *testing.T) {
		_, _, _, found := simproject.GetAccountWithVouchers(r, ctx, tk.BankKeeper, tk.ProjectKeeper, accs, false)
		require.False(t, found)
	})

	t.Run("should find account with vouchers for a project with launch triggered", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 0)
		project.MainnetInitialized = true
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTriggered = true
		chain.IsMainnet = true
		project.MainnetID, err = tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		project.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, project)
		require.NoError(t, err)
		mint(acc.Address, sample.Vouchers(r, project.ProjectID))
		prjtID, acc, coins, found := simproject.GetAccountWithVouchers(r, ctx, tk.BankKeeper, tk.ProjectKeeper, accs, false)
		require.True(t, found)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.False(t, coins.Empty())
		require.Contains(t, accs, acc)
	})

	t.Run("should find account with vouchers", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 1)
		project.MainnetInitialized = false
		project.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, project)
		require.NoError(t, err)
		mint(acc.Address, sample.Vouchers(r, project.ProjectID))
		prjtID, acc, coins, found := simproject.GetAccountWithVouchers(r, ctx, tk.BankKeeper, tk.ProjectKeeper, accs, true)
		require.True(t, found)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.False(t, coins.Empty())
		require.Contains(t, accs, acc)
	})
}

func TestGetAccountWithShares(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	accs := sample.SimAccounts()
	var err error

	t.Run("should find no account", func(t *testing.T) {
		_, _, _, found := simproject.GetAccountWithShares(r, ctx, tk.ProjectKeeper, accs, false)
		require.False(t, found)
	})

	t.Run("should not find account not part of sim accounts", func(t *testing.T) {
		sampleAddr := sample.AccAddress(r)
		projectID := uint64(10)
		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(projectID, sampleAddr), projecttypes.MainnetAccount{
			ProjectID: projectID,
			Address:   sampleAddr.String(),
			Shares:    sample.Shares(r),
		})
		require.NoError(t, err)
		_, _, _, found := simproject.GetAccountWithShares(r, ctx, tk.ProjectKeeper, accs, false)
		require.False(t, found)
		err = tk.ProjectKeeper.MainnetAccount.Remove(ctx, collections.Join(projectID, sampleAddr))
		require.NoError(t, err)
	})

	t.Run("should find account from project with launched mainnet can be retrieved", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 0)
		project.MainnetInitialized = true
		chain := sample.Chain(r, 0, 0)
		chain.LaunchTriggered = true
		chain.IsMainnet = true
		project.MainnetID, err = tk.LaunchKeeper.AppendChain(ctx, chain)
		require.NoError(t, err)
		project.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, project)
		require.NoError(t, err)
		share := sample.Shares(r)
		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(project.ProjectID, acc.Address), projecttypes.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:   acc.Address.String(),
			Shares:    share,
		})
		require.NoError(t, err)
		prjtID, acc, shareRetrieved, found := simproject.GetAccountWithShares(r, ctx, tk.ProjectKeeper, accs, false)
		require.True(t, found)
		require.Contains(t, accs, acc)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.EqualValues(t, share, shareRetrieved)
	})

	t.Run("should find account from project", func(t *testing.T) {
		acc, _ := simtypes.RandomAcc(r, accs)
		project := sample.Project(r, 1)
		project.MainnetInitialized = false
		project.ProjectID, err = tk.ProjectKeeper.AppendProject(ctx, project)
		require.NoError(t, err)
		share := sample.Shares(r)
		err = tk.ProjectKeeper.MainnetAccount.Set(ctx, collections.Join(project.ProjectID, acc.Address), projecttypes.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:   acc.Address.String(),
			Shares:    share,
		})
		require.NoError(t, err)
		prjtID, acc, shareRetrieved, found := simproject.GetAccountWithShares(r, ctx, tk.ProjectKeeper, accs, true)
		require.True(t, found)
		require.Contains(t, accs, acc)
		require.EqualValues(t, project.ProjectID, prjtID)
		require.EqualValues(t, share, shareRetrieved)
	})
}
