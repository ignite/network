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

func SimulateMsgUnredeemVouchers(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgUnredeemVouchers{}
		// Find a random account from a random project that contains shares
		prjtID, simAccount, shares, found := GetAccountWithShares(r, ctx, k, accs, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip unredeem vouchers"), nil, nil
		}

		msg = types.NewMsgUnredeemVouchers(
			simAccount.Address.String(),
			prjtID,
			shares,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins(), txGen)
	}
}
