package simulation

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

// TypedMsg extends sdk.Msg with Type method
type TypedMsg interface {
	sdk.Msg
	Type() string
}

// ShareDenoms are the denom used for the shares in the simulation
var ShareDenoms = []string{"s/foo", "s/bar", "s/toto"}

// GetCoordSimAccount finds an account associated with a coordinator profile from simulation accounts
func GetCoordSimAccount(
	r *rand.Rand,
	ctx sdk.Context,
	pk types.ProfileKeeper,
	accs []simtypes.Account,
) (simtypes.Account, uint64, bool) {
	// Choose a random coordinator
	coords, err := pk.Coordinators(ctx)
	if err != nil {
		return simtypes.Account{}, 0, false
	}
	coordNb := len(coords)
	if coordNb == 0 {
		return simtypes.Account{}, 0, false
	}
	coord := coords[r.Intn(coordNb)]

	// Find the account linked to this address
	for _, acc := range accs {
		if acc.Address.String() == coord.Address && coord.Active {
			return acc, coord.CoordinatorID, true
		}
	}
	return simtypes.Account{}, 0, false
}

// GetCoordSimAccountWithProjectID finds an account associated with a coordinator profile from simulation accounts
// and a project created by this coordinator. The boolean flag `requireNoMainnetLaunchTriggered` is ignored if
// the flag `requireNoMainnetInitialized` is set to `true`
func GetCoordSimAccountWithProjectID(
	r *rand.Rand,
	ctx sdk.Context,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetInitialized bool,
	requireNoMainnetLaunchTriggered bool,
) (simtypes.Account, uint64, bool) {
	projects, err := k.Projects(ctx)
	if err != nil {
		return simtypes.Account{}, 0, false
	}
	prjtNb := len(projects)
	if prjtNb == 0 {
		return simtypes.Account{}, 0, false
	}

	r.Shuffle(prjtNb, func(i, j int) {
		projects[i], projects[j] = projects[j], projects[i]
	})

	// select first project after shuffle
	prjt := projects[0]
	// If a criteria is required for the project, we simply fetch the first one that satisfies the criteria
	if requireNoMainnetInitialized {
		var prjtFound bool
		for _, project := range projects {
			if !project.MainnetInitialized {
				prjt = project
				prjtFound = true
				break
			}
		}
		if !prjtFound {
			return simtypes.Account{}, 0, false
		}
	}
	if !requireNoMainnetInitialized && requireNoMainnetLaunchTriggered {
		var prjtFound bool
		for _, project := range projects {
			launched, _ := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
			if !launched {
				prjt = project
				prjtFound = true
				break
			}
		}
		if !prjtFound {
			return simtypes.Account{}, 0, false
		}
	}

	// Find the sim account of the project coordinator
	coord, err := pk.GetCoordinator(ctx, prjt.CoordinatorID)
	if err != nil {
		return simtypes.Account{}, 0, false
	}
	for _, acc := range accs {
		if acc.Address.String() == coord.Address && coord.Active {
			return acc, prjt.ProjectID, true
		}
	}

	return simtypes.Account{}, 0, false
}

// GetSharesFromProject returns a small portion of shares that can be minted as vouchers or added to an account
func GetSharesFromProject(r *rand.Rand, ctx sdk.Context, k keeper.Keeper, prjtID uint64) (types.Shares, bool) {
	prjt, err := k.GetProject(ctx, prjtID)
	if err != nil {
		return types.EmptyShares(), false
	}

	totalShares, err := k.TotalShares.Get(ctx)
	if err != nil {
		return types.EmptyShares(), false
	}

	var shares sdk.Coins
	for _, share := range ShareDenoms {
		remaining := int64(totalShares) - prjt.AllocatedShares.AmountOf(share)
		if remaining == 0 {
			continue
		}

		shareNb := r.Int63n(5000) + 10
		if shareNb > remaining {
			shareNb = remaining
		}
		shares = append(shares, sdk.NewCoin(share, sdkmath.NewInt(shareNb)))
	}

	// No shares can be distributed
	if shares.Empty() {
		return types.EmptyShares(), false
	}
	shares = shares.Sort()
	return types.Shares(shares), true
}

// GetAccountWithVouchers returns an account that has vouchers for a project
func GetAccountWithVouchers(
	r *rand.Rand,
	ctx sdk.Context,
	bk types.BankKeeper,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetLaunchTriggered bool,
) (prjtID uint64, account simtypes.Account, coins sdk.Coins, found bool) {
	var err error
	var accountAddr sdk.AccAddress

	// Parse all account balances and find one with vouchers
	bk.IterateAllBalances(ctx, func(addr sdk.AccAddress, coin sdk.Coin) bool {
		prjtID, err = types.VoucherProject(coin.Denom)
		if err != nil {
			return false
		}

		if requireNoMainnetLaunchTriggered {
			project, err := k.GetProject(ctx, prjtID)
			if err != nil {
				return false
			}
			launched, err := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
			if err != nil || launched {
				return false
			}
		}

		found = true
		accountAddr = addr
		return true
	})

	// No account has vouchers
	if !found {
		return 0, account, coins, false
	}

	// Fetch from the vouchers of the project owned by the account
	bk.IterateAccountBalances(ctx, accountAddr, func(coin sdk.Coin) bool {
		coinCampID, err := types.VoucherProject(coin.Denom)
		if err == nil && coinCampID == prjtID {
			// fetch a part of each voucher hold by the account
			amt, err := simtypes.RandPositiveInt(r, coin.Amount)
			if err == nil {
				coins = append(coins, sdk.NewCoin(coin.Denom, amt))
			}
		}
		return false
	})
	if coins.Empty() {
		return 0, account, coins, false
	}

	coins = coins.Sort()

	// Find the sim account
	for _, acc := range accs {
		if found = acc.Address.Equals(accountAddr); found {
			return prjtID, acc, coins, true
		}
	}
	return 0, account, coins, false
}

// GetAccountWithShares returns an account that contains allocated shares with its associated project
func GetAccountWithShares(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
	requireNoMainnetLaunchTriggered bool,
) (uint64, simtypes.Account, types.Shares, bool) {
	mainnetAccounts, err := k.GetAllMainnetAccount(ctx)
	if err != nil {
		return 0, simtypes.Account{}, nil, false
	}
	nb := len(mainnetAccounts)

	// No account have been created yet
	if nb == 0 {
		return 0, simtypes.Account{}, nil, false
	}

	r.Shuffle(nb, func(i, j int) {
		mainnetAccounts[i], mainnetAccounts[j] = mainnetAccounts[j], mainnetAccounts[i]
	})

	// select a mainnet account
	var mainnetAccount types.MainnetAccount
	for _, mAcc := range mainnetAccounts {
		if requireNoMainnetLaunchTriggered {
			project, err := k.GetProject(ctx, mAcc.ProjectID)
			if err != nil {
				continue
			}
			launched, _ := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
			if launched {
				continue
			}
		}
		mainnetAccount = mAcc
		break
	}

	// Find the associated sim account
	for _, acc := range accs {
		if acc.Address.String() == mainnetAccount.Address {
			return mainnetAccount.ProjectID, acc, mainnetAccount.Shares, true
		}
	}
	return 0, simtypes.Account{}, nil, false
}

// deliverSimTx delivers the tx for simulation from the provided message
func deliverSimTx(
	r *rand.Rand,
	app *baseapp.BaseApp,
	ctx sdk.Context,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	simAccount simtypes.Account,
	msg TypedMsg,
	coinsSpent sdk.Coins,
	txGen client.TxConfig,
) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	txCtx := simulation.OperationInput{
		R:               r,
		App:             app,
		TxGen:           txGen,
		Cdc:             nil,
		Msg:             msg,
		Context:         ctx,
		SimAccount:      simAccount,
		AccountKeeper:   ak,
		Bankkeeper:      bk,
		ModuleName:      types.ModuleName,
		CoinsSpentInMsg: coinsSpent,
	}
	return simulation.GenAndDeliverTxWithRandFees(txCtx)
}
