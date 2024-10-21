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

func SimulateMsgBurnVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgBurnVouchers{}
		prjtID, simAccount, vouchers, found := GetAccountWithVouchers(r, ctx, bk, k, accs, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip burn vouchers"), nil, nil
		}

		msg = types.NewMsgBurnVouchers(
			simAccount.Address.String(),
			prjtID,
			vouchers,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, vouchers, txGen)
	}
}
