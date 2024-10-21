package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/project/types"
)

const (
	accountWithoutProjectRoute = "account-without-project"
	projectSharesRoute         = "project-shares"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, accountWithoutProjectRoute,
		AccountWithoutProjectInvariant(k))
	ir.RegisterRoute(types.ModuleName, projectSharesRoute,
		ProjectSharesInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := AccountWithoutProjectInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return ProjectSharesInvariant(k)(ctx)
	}
}

// AccountWithoutProjectInvariant invariant that checks if
// the `MainnetAccount` project exist.
func AccountWithoutProjectInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all, err := k.GetAllMainnetAccount(ctx)
		if err != nil {
			return "", false
		}
		for _, acc := range all {
			if _, err := k.GetProject(ctx, acc.ProjectId); err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, accountWithoutProjectRoute,
					fmt.Sprintf("%s: %d", err, acc.ProjectId),
				), true
			}
		}
		return "", false
	}
}

// ProjectSharesInvariant invariant that checks, for all projects, if the amount of allocated shares is equal to
// the sum of `MainnetVestingAccount` and `MainnetAccount` shares plus
// the amount of vouchers in circulation plus
// the total shares of special allocations
func ProjectSharesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		accountSharesByProject := make(map[uint64]types.Shares)

		// get all mainnet account shares
		accounts, err := k.GetAllMainnetAccount(ctx)
		if err != nil {
			return "", false
		}
		for _, acc := range accounts {
			if _, ok := accountSharesByProject[acc.ProjectId]; !ok {
				accountSharesByProject[acc.ProjectId] = types.EmptyShares()
			}
			accountSharesByProject[acc.ProjectId] = types.IncreaseShares(
				accountSharesByProject[acc.ProjectId],
				acc.Shares,
			)
		}

		projects, err := k.Projects(ctx)
		if err != nil {
			return "", false
		}
		for _, project := range projects {
			projectID := project.ProjectId
			expectedAllocatedSharesShares := accountSharesByProject[projectID]

			// read existing denoms from allocated shares of the project to check possible minted vouchers
			allocated, err := types.SharesToVouchers(project.GetAllocatedShares(), projectID)
			if err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, projectSharesRoute,
					fmt.Sprintf("project %d: allocated shares can't be converted to vouchers %s",
						projectID,
						err.Error(),
					),
				), true
			}

			// get the supply for the circulating vouchers
			vouchers := sdk.NewCoins()
			for _, a := range allocated {
				voucherSupply := k.bankKeeper.GetSupply(ctx, a.Denom)
				vouchers = vouchers.Add(voucherSupply)
			}

			// convert to shares and add to the project shares - since we are converting shares to vouchers earlier,
			// this conversion back to shares will never fail by design, thus we can ignore the error
			vShares, _ := types.VouchersToShares(vouchers, projectID)
			expectedAllocatedSharesShares = types.IncreaseShares(expectedAllocatedSharesShares, vShares)

			// increase expected shares with special allocations
			expectedAllocatedSharesShares = types.IncreaseShares(
				expectedAllocatedSharesShares,
				project.SpecialAllocations.TotalShares(),
			)

			if !types.IsEqualShares(expectedAllocatedSharesShares, project.GetAllocatedShares()) {
				return sdk.FormatInvariant(
					types.ModuleName, projectSharesRoute,
					fmt.Sprintf("project %d: expected allocated shares: %s, actual allocated shares: %s",
						projectID,
						expectedAllocatedSharesShares.String(),
						project.GetAllocatedShares().String(),
					),
				), true
			}
		}
		return "", false
	}
}
