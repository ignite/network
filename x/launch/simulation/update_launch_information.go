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

func SimulateMsgUpdateLaunchInformation(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgUpdateLaunchInformation{}

		// Select a chain with a valid coordinator account
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "chain not found"), nil, nil
		}

		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.LaunchId)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "coordinator account not found"), nil, nil
		}

		modify := r.Intn(100) < 50
		msg = sample.MsgUpdateLaunchInformation(r,
			simAccount.Address.String(),
			chain.LaunchId,
			modify,
			!modify,
			modify,
			!modify && r.Intn(100) < 50,
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
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
