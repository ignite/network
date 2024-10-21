package simulation

import (
	"math/rand"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

// SimulateMsgRequestAddGenesisAccount simulates a MsgRequestAddGenesisAccount message
func SimulateMsgRequestAddGenesisAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgSendRequest{}

		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}
		fee := params.RequestFee

		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, true)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "non-triggered chain not found"), nil, nil
		}

		// Select a random account no in genesis
		r.Shuffle(len(accs), func(i, j int) {
			accs[i], accs[j] = accs[j], accs[i]
		})
		var simAccount simtypes.Account
		var availableAccount bool
		for _, acc := range accs {
			_, err := k.GenesisAccount.Get(ctx, collections.Join(chain.LaunchId, acc.Address))
			if err != nil {
				continue
			}
			_, err = k.VestingAccount.Get(ctx, collections.Join(chain.LaunchId, acc.Address))
			if err != nil {
				continue
			}
			simAccount = acc
			availableAccount = true
			break
		}
		if !availableAccount {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no available account"), nil, nil
		}

		msg = sample.MsgSendRequestWithAddAccount(r,
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchId,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgRequestAddVestingAccount simulates a MsgRequestAddVestingAccount message
func SimulateMsgRequestAddVestingAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgSendRequest{}
		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}
		fee := params.RequestFee

		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, true)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "non-triggered chain not found"), nil, nil
		}

		// Select a random account no in genesis
		r.Shuffle(len(accs), func(i, j int) {
			accs[i], accs[j] = accs[j], accs[i]
		})
		var simAccount simtypes.Account
		var availableAccount bool
		for _, acc := range accs {
			_, err := k.GenesisAccount.Get(ctx, collections.Join(chain.LaunchId, acc.Address))
			if err != nil {
				continue
			}
			_, err = k.VestingAccount.Get(ctx, collections.Join(chain.LaunchId, acc.Address))
			if err != nil {
				continue
			}
			simAccount = acc
			availableAccount = true
			break
		}
		if !availableAccount {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "no available account"), nil, nil
		}

		msg = sample.MsgSendRequestWithAddVestingAccount(r,
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchId,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgRequestRemoveAccount simulates a MsgRequestRemoveAccount message
func SimulateMsgRequestRemoveAccount(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgSendRequest{}
		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}
		fee := params.RequestFee

		type accChain struct {
			address  string
			launchID uint64
		}

		// build list of genesis and vesting accounts
		accChains := make([]accChain, 0)
		genAccs, err := k.AllGenesisAccount(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "all genesis account error"), nil, err
		}
		for _, acc := range genAccs {
			accChains = append(accChains, accChain{
				address:  acc.Address,
				launchID: acc.LaunchId,
			})
		}
		vestAccs, err := k.AllVestingAccount(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "all vesting account error"), nil, err
		}
		for _, acc := range vestAccs {
			accChains = append(accChains, accChain{
				address:  acc.Address,
				launchID: acc.LaunchId,
			})
		}

		// add entropy
		r.Shuffle(len(accChains), func(i, j int) {
			accChains[i], accChains[j] = accChains[j], accChains[i]
		})

		var (
			simAccount simtypes.Account
			accAddr    string
			accChainID uint64
		)
		found := false
		for _, accChain := range accChains {
			if IsLaunchTriggeredChain(ctx, k, accChain.launchID) {
				continue
			}
			// get coordinator account
			var err error
			simAccount, err = FindChainCoordinatorAccount(ctx, k, accs, accChain.launchID)
			if err != nil {
				continue
			}
			accAddr = accChain.address
			accChainID = accChain.launchID
			found = true
			break
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "genesis account not found"), nil, nil
		}

		msg = types.NewMsgSendRequest(
			simAccount.Address.String(),
			accChainID,
			types.NewAccountRemoval(accAddr),
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgRequestAddValidator simulates a MsgRequestAddValidator message
func SimulateMsgRequestAddValidator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgSendRequest{}
		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}
		fee := params.RequestFee

		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "non-triggered chain not found"), nil, nil
		}
		// Select a random account
		simAccount, _ := simtypes.RandomAcc(r, accs)
		// Select between new address or coordinator address randomly
		msg = sample.MsgSendRequestWithAddValidator(r,
			simAccount.Address.String(),
			simAccount.Address.String(),
			chain.LaunchId,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgRequestRemoveValidator simulates a MsgRequestRemoveValidator message
func SimulateMsgRequestRemoveValidator(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgSendRequest{}
		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}
		fee := params.RequestFee

		// Select a validator
		simAccount, valAcc, found := FindRandomValidator(r, ctx, k, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "validator not found"), nil, nil
		}

		msg = sample.MsgSendRequestWithRemoveValidator(
			simAccount.Address.String(),
			valAcc.Address,
			valAcc.LaunchId,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgRequestParamChange simulates a MsgSendRequest message with ParamChange content
func SimulateMsgRequestParamChange(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := types.MsgSendRequest{}
		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}
		fee := params.RequestFee

		// Select a chain without launch triggered
		chain, found := FindRandomChain(r, ctx, k, false, false)
		if !found {
			// No message if no non-triggered chain
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "non-triggered chain not found"), nil, nil
		}
		simAccount, _ := simtypes.RandomAcc(r, accs)

		msg = sample.MsgSendRequestWithParamChange(
			r,
			simAccount.Address.String(),
			chain.LaunchId,
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
			CoinsSpentInMsg: fee,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
