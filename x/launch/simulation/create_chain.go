package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

func SimulateMsgCreateChain(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgCreateChain{}

		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}

		// Check if the coordinator address is already in the store
		var found bool
		var simAccount simtypes.Account
		for _, acc := range accs {
			_, err := k.GetProfileKeeper().CoordinatorIDFromAddress(ctx, acc.Address)
			if err == nil {
				simAccount = acc
				found = true
				break
			}
		}
		if !found {
			// No message if no coordinator
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip create a new chain"), nil, nil
		}

		// skip if account cannot cover creation fee
		creationFee := params.ChainCreationFee
		msg = sample.MsgCreateChain(r,
			simAccount.Address.String(),
			"",
			false,
			0,
		)

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             &msg,
			Context:         ctx,
			SimAccount:      simAccount,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: creationFee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
