package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/types"
)

// generate a Tx with 2 signatures if validator account is not equal to operator address account
func genAndDeliverTxWithRandFeesAddOpAddr(txCtx simulation.OperationInput, opSimAcc simtypes.Account) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	spendable := txCtx.Bankkeeper.SpendableCoins(txCtx.Context, account.GetAddress())
	opAccout := txCtx.AccountKeeper.GetAccount(txCtx.Context, opSimAcc.Address)

	accNumbers := []uint64{account.GetAccountNumber()}
	accSequences := []uint64{account.GetSequence()}
	privs := []cryptotypes.PrivKey{txCtx.SimAccount.PrivKey}
	if account != opAccout {
		accNumbers = append(accNumbers, opAccout.GetAccountNumber())
		accSequences = append(accSequences, opAccout.GetSequence())
		privs = append(privs, opSimAcc.PrivKey)
	}

	var fees sdk.Coins
	var err error

	coins, hasNeg := spendable.SafeSub(txCtx.CoinsSpentInMsg...)
	if hasNeg {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "message doesn't leave room for fees"), nil, err
	}

	fees, err = sample.Fees(txCtx.R, coins)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "unable to generate fees"), nil, err
	}

	tx, err := simtestutil.GenSignedMockTx(
		txCtx.R,
		txCtx.TxGen,
		[]sdk.Msg{txCtx.Msg},
		fees,
		simtestutil.DefaultGenTxGas,
		txCtx.Context.ChainID(),
		accNumbers,
		accSequences,
		privs...,
	)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "unable to generate mock tx"), nil, err
	}

	_, _, err = txCtx.App.SimDeliver(txCtx.TxGen.TxEncoder(), tx)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, ""), nil, nil
}

func SimulateMsgAddValidatorOperatorAddress(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		// choose two addresses that are not equal
		opAccount, _ := simtypes.RandomAcc(r, accs)
		for simAccount.Address.Equals(opAccount.Address) {
			opAccount, _ = simtypes.RandomAcc(r, accs)
		}

		msg := &types.MsgAddValidatorOperatorAddress{
			ValidatorAddress: simAccount.Address.String(),
			OperatorAddress:  opAccount.Address.String(),
		}

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
		return genAndDeliverTxWithRandFeesAddOpAddr(txCtx, opAccount)
	}
}

func SimulateMsgUpdateValidatorDescription(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		desc := sample.ValidatorDescription(sample.String(r, 50))
		msg := types.NewMsgUpdateValidatorDescription(
			simAccount.Address.String(),
			desc.Identity,
			desc.Moniker,
			desc.Website,
			desc.SecurityContact,
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
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
