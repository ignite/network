package simulation

import (
	"fmt"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/x/profile/keeper"
)

// FindAccount find a specific address from an account list
func FindAccount(k *keeper.Keeper, accs []simtypes.Account, address string) (simtypes.Account, error) {
	creator, err := k.AddressCodec().StringToBytes(address)
	if err != nil {
		return simtypes.Account{}, err
	}
	simAccount, found := simtypes.FindAccount(accs, sdk.AccAddress(creator))
	if !found {
		return simAccount, fmt.Errorf("address %s not found in the sim accounts", address)
	}
	return simAccount, nil
}

// generate a Tx with more than one signature
func genAndDeliverTxWithSigners(txCtx simulation.OperationInput, accounts ...simtypes.Account) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	spendable := txCtx.Bankkeeper.SpendableCoins(txCtx.Context, account.GetAddress())

	var (
		accMap       = make(map[string]struct{})
		accNumbers   = []uint64{account.GetAccountNumber()}
		accSequences = []uint64{account.GetSequence()}
		privs        = []cryptotypes.PrivKey{txCtx.SimAccount.PrivKey}
	)
	accMap[txCtx.SimAccount.Address.String()] = struct{}{}
	for _, acc := range accounts {
		if _, ok := accMap[acc.Address.String()]; ok {
			continue
		}
		accMap[acc.Address.String()] = struct{}{}
		account := txCtx.AccountKeeper.GetAccount(txCtx.Context, acc.Address)
		accNumbers = append(accNumbers, account.GetAccountNumber())
		accSequences = append(accSequences, account.GetSequence())
		privs = append(privs, acc.PrivKey)
	}

	// feePayer is the first signer
	coins, hasNeg := spendable.SafeSub(txCtx.CoinsSpentInMsg...)
	if hasNeg {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "message doesn't leave room for fees"), nil, nil
	}

	fees, err := simtypes.RandomFees(txCtx.R, txCtx.Context, coins)
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

	_, _, err = txCtx.App.SimTxFinalizeBlock(txCtx.TxGen.TxEncoder(), tx)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, sdk.MsgTypeURL(txCtx.Msg), "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, ""), nil, nil
}
