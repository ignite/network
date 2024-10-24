package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func SimulateMsgUpdateSpecialAllocations(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgUpdateSpecialAllocations{}
		simAccount, prjtID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, false, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update special allocations"), nil, nil
		}

		// get shares for both genesis distribution and claimable airdrop
		genesisDistribution, getShares := GetSharesFromProject(r, ctx, k, prjtID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update special allocations"), nil, nil
		}
		claimableAirdrop, getShares := GetSharesFromProject(r, ctx, k, prjtID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update special allocations"), nil, nil
		}

		// GetSharesFromProject returns a number of available shares for a project
		// potentially genesisDistribution + claimableAirdrop can overflow the available shares
		// we divide by two all amounts to avoid overflowing available shares
		for i, s := range genesisDistribution {
			genesisDistribution[i].Amount = s.Amount.QuoRaw(2)
		}
		for i, s := range claimableAirdrop {
			claimableAirdrop[i].Amount = s.Amount.QuoRaw(2)
		}

		msg = types.NewMsgUpdateSpecialAllocations(
			simAccount.Address.String(),
			prjtID,
			types.NewSpecialAllocations(genesisDistribution, claimableAirdrop),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins(), txGen)
	}
}
