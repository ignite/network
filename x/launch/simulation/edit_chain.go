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

func SimulateMsgEditChain(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgEditChain{}

		// Select a chain with a valid coordinator account
		chain, found := FindRandomChain(r, ctx, k, r.Intn(100) < 50, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "chain not found"), nil, nil
		}

		simAccount, err := FindChainCoordinatorAccount(ctx, k, accs, chain.LaunchId)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "coordinator account not found"), nil, nil
		}

		modify := r.Intn(100) < 50
		setProjectID := r.Intn(100) < 50
		// do not set projectID if already set
		if chain.HasProject {
			setProjectID = false
		}
		// ensure there is always a value to edit
		if !modify && !setProjectID {
			modify = true
		}

		projectID := uint64(0)
		ok := false
		if setProjectID {
			projectID, ok = FindCoordinatorProject(r, ctx, k.GetProjectKeeper(), chain.CoordinatorId, chain.LaunchId)
			if !ok {
				modify = true
				setProjectID = false
			}
		}

		msg = sample.MsgEditChain(r,
			simAccount.Address.String(),
			chain.LaunchId,
			setProjectID,
			projectID,
			modify,
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
