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

func SimulateMsgRedeemVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgRedeemVouchers{}
		prjtID, simAccount, vouchers, found := GetAccountWithVouchers(r, ctx, bk, k, accs, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip redeem vouchers"), nil, nil
		}

		// Select a random account to redeem vouchers into
		accountNb := r.Intn(len(accs))

		msg = types.NewMsgRedeemVouchers(
			simAccount.Address.String(),
			prjtID,
			accs[accountNb].Address.String(),
			vouchers,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, vouchers, txGen)
	}
}
