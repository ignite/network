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

func SimulateMsgSettleRequest(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgSettleRequest{}

		// Select a random request without launch triggered
		request, found := FindRandomRequest(r, ctx, k)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "request for non-triggered chain not found"), nil, nil
		}

		// Find coordinator account
		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, request.LaunchId)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, nil
		}

		approve := r.Intn(100) < 50
		msg = sample.MsgSettleRequest(
			simAccount.Address.String(),
			request.LaunchId,
			request.RequestId,
			approve,
		)

		// if we cannot check the request, reject
		if err := keeper.CheckRequest(ctx, k, request.LaunchId, request); err != nil {
			msg.Approve = false
		}

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
