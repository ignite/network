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

func SimulateMsgMintVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgMintVouchers{}
		simAccount, prjtID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip mint vouchers"), nil, nil
		}

		shares, getShares := GetSharesFromProject(r, ctx, k, prjtID)
		if !getShares {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip mint vouchers"), nil, nil
		}

		msg = types.NewMsgMintVouchers(
			simAccount.Address.String(),
			prjtID,
			shares,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins(), txGen)
	}
}
