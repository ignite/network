package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/x/participation/keeper"
	"github.com/ignite/network/x/participation/types"
)

func SimulateMsgParticipate(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	fk types.FundraisingKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgParticipate{}
		auction, found := RandomAuctionParticipationEnabled(ctx, r, fk, k)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no valid auction found"), nil, nil
		}

		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}

		tierList := params.ParticipationTierList
		tier, found := RandomTierFromList(r, tierList)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no valid tiers"), nil, nil
		}

		simAccount, _, found := RandomAccWithAvailableAllocations(ctx, r, k, accs, tier.RequiredAllocations, auction.GetId())
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no account with allocations"), nil, nil
		}

		msg = types.NewMsgParticipate(
			simAccount.Address.String(),
			auction.GetId(),
			tier.TierId,
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
