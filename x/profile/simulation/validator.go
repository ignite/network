package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/types"
)

func SimulateMsgAddValidatorOperatorAddress(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		// choose two addresses that are not equal
		simAccount, _ := simtypes.RandomAcc(r, accs)
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
		return genAndDeliverTxWithSigners(txCtx, opAccount)
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
