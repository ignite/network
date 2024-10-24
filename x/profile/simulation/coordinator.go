package simulation

import (
	"errors"
	"math/rand"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/keeper"
	"github.com/ignite/network/x/profile/types"
)

// FindCoordinatorAccount find a sim account for a coordinator that exists or not
func FindCoordinatorAccount(
	r *rand.Rand,
	ctx sdk.Context,
	k keeper.Keeper,
	accs []simtypes.Account,
	exist bool,
) (simtypes.Account, bool) {
	// Randomize the set for coordinator operation entropy
	r.Shuffle(len(accs), func(i, j int) {
		accs[i], accs[j] = accs[j], accs[i]
	})

	for _, acc := range accs {
		coordByAddress, err := k.CoordinatorByAddress.Get(ctx, acc.Address)
		found := !errors.Is(err, collections.ErrNotFound)
		if found == exist {
			coord, err := k.Coordinator.Get(ctx, coordByAddress.CoordinatorId)
			if err == nil && !coord.Active {
				continue
			}
			return acc, true
		}
	}
	return simtypes.Account{}, false
}

func SimulateMsgCreateCoordinator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgCreateCoordinator{}

		// Find an account with no coordinator
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs, false)
		if !found {
			// No message if all coordinator created
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip coordinator creation"), nil, nil
		}

		msg = types.NewMsgCreateCoordinator(
			simAccount.Address.String(),
			sample.String(r, 30),
			sample.String(r, 30),
			sample.String(r, 30),
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateCoordinatorDescription(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgUpdateCoordinatorDescription{}

		// Find an account with coordinator associated
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs, true)
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update coordinator description"), nil, nil
		}

		desc := sample.CoordinatorDescription(r)
		msg = types.NewMsgUpdateCoordinatorDescription(
			simAccount.Address.String(),
			desc.Identity,
			desc.Website,
			desc.Details,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateCoordinatorAddress(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgUpdateCoordinatorAddress{}

		// Select a random account
		coord, found := FindCoordinatorAccount(r, ctx, k, accs, true)
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update coordinator address"), nil, nil
		}
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs, false)
		if !found && coord.Address.String() != simAccount.Address.String() {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update coordinator address"), nil, nil
		}
		msg = types.NewMsgUpdateCoordinatorAddress(coord.Address.String(), simAccount.Address.String())
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      coord,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDisableCoordinator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgDisableCoordinator{}
		// Find an account with coordinator associated
		// avoid delete coordinator associated a chain (id 0,1,2)
		simAccount, found := FindCoordinatorAccount(r, ctx, k, accs[3:], true)
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update coordinator delete"), nil, nil
		}

		msg = types.NewMsgDisableCoordinator(simAccount.Address.String())
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
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
