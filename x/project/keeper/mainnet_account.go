package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	"github.com/ignite/network/x/project/types"
)

func (k Keeper) GetMainnetAccount(ctx context.Context, projectID uint64, address sdk.AccAddress) (types.MainnetAccount, error) {
	acc, err := k.MainnetAccount.Get(ctx, collections.Join(projectID, address))
	if errors.Is(err, collections.ErrNotFound) {
		return types.MainnetAccount{}, types.ErrMainnetAccountNotFound
	}
	return acc, err
}

// GetAllMainnetAccount returns all MainnetAccount
func (k Keeper) GetAllMainnetAccount(ctx context.Context) ([]types.MainnetAccount, error) {
	accs := make([]types.MainnetAccount, 0)
	err := k.MainnetAccount.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], value types.MainnetAccount) (bool, error) {
		accs = append(accs, value)
		return false, nil
	})
	return accs, err
}

func (k Keeper) MainnetAccountBalance(ctx context.Context, projectID uint64, address sdk.AccAddress) (types.MainnetAccountBalance, error) {
	// get project and share information
	totalShareNumber, err := k.TotalShares.Get(ctx)
	if err != nil {
		return types.MainnetAccountBalance{}, err
	}

	project, err := k.GetProject(ctx, projectID)
	if err != nil {
		return types.MainnetAccountBalance{}, err
	}

	// get account balance
	acc, err := k.MainnetAccount.Get(ctx, collections.Join(projectID, address))
	if err != nil {
		return types.MainnetAccountBalance{}, err
	}

	balance, err := acc.Shares.CoinsFromTotalSupply(project.TotalSupply, totalShareNumber)
	if err != nil {
		return types.MainnetAccountBalance{}, err
	}

	return types.MainnetAccountBalance{
		ProjectID: acc.ProjectID,
		Address:   acc.Address,
		Coins:     balance,
	}, nil
}

func (k Keeper) ListMainnetAccountBalance(ctx context.Context, projectID uint64) ([]types.MainnetAccountBalance, error) {
	// get project and share information
	totalShareNumber, err := k.TotalShares.Get(ctx)
	if err != nil {
		return nil, err
	}

	project, err := k.GetProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	mainnetAccountBalances := make([]types.MainnetAccountBalance, 0)
	err = k.MainnetAccount.Walk(ctx, nil, func(_ collections.Pair[uint64, sdk.AccAddress], acc types.MainnetAccount) (stop bool, err error) {
		balance, err := acc.Shares.CoinsFromTotalSupply(project.TotalSupply, totalShareNumber)
		if err != nil {
			return true, err
		}

		// add the balance if not zero
		if !balance.IsZero() {
			mainnetAccountBalance := types.MainnetAccountBalance{
				ProjectID: acc.ProjectID,
				Address:   acc.Address,
				Coins:     balance,
			}
			mainnetAccountBalances = append(mainnetAccountBalances, mainnetAccountBalance)
		}

		return false, nil
	})
	return mainnetAccountBalances, err
}
