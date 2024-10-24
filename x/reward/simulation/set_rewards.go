package simulation

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/x/reward/keeper"
	"github.com/ignite/network/x/reward/types"
)

const (
	ProbabilityCreateRewardPool = 70
	ProbabilityCloseRewardPool  = 70

	ActionCreate = iota
	ActionEdit
	ActionClose
)

func SimulateMsgSetRewards(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgSetRewards{}
		createRewardPool := r.Int63n(100) < ProbabilityCreateRewardPool
		closeRewardPool := r.Int63n(100) < ProbabilityCloseRewardPool
		checkBalance := true

		// choose action to be taken
		action := ActionCreate
		if !createRewardPool && closeRewardPool {
			action = ActionClose
			checkBalance = false
		} else if !createRewardPool && !closeRewardPool {
			action = ActionEdit
		}

		wantCoin := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(r.Int63n(1_000_000))))
		chain, found := FindRandomChainWithCoordBalance(r, ctx, k, bk, createRewardPool, checkBalance, wantCoin)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no viable chain to be found"), nil, nil
		}

		coordinator, err := k.GetProfileKeeper().GetCoordinator(ctx, chain.CoordinatorId)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "chain does not have a coordinator"), nil, nil
		}

		coordAccAddr, err := k.AddressCodec().StringToBytes(coordinator.Address)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, err
		}

		simAccount, found := simtypes.FindAccount(accs, sdk.AccAddress(coordAccAddr))
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "simulation account not found for chain coordinator"), nil, nil
		}

		// initialize basic message
		msg = &types.MsgSetRewards{
			Provider: simAccount.Address.String(),
			LaunchId: chain.LaunchId,
		}

		// set message based on action to be taken
		switch action {
		case ActionCreate:
			msg.LastRewardHeight = ctx.BlockHeight() + r.Int63n(1000)
			msg.Coins = wantCoin
		case ActionEdit:
			pool, err := k.RewardPool.Get(ctx, chain.LaunchId)
			if err != nil {
				return simtypes.NoOpMsg(types.ModuleName, msg.Type(), err.Error()), nil, err
			}
			msg.LastRewardHeight = pool.LastRewardHeight + r.Int63n(1000)
			msg.Coins = wantCoin
		case ActionClose:
			msg.LastRewardHeight = 0
			msg.Coins = sdk.NewCoins()
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
			CoinsSpentInMsg: wantCoin,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
