package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func SimulateMsgUpdateTotalSupply(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgUpdateTotalSupply{}
		simAccount, prjtID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, true, true)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip update total supply"), nil, nil
		}

		msg = types.NewMsgUpdateTotalSupply(
			simAccount.Address.String(),
			prjtID,
			sample.TotalSupply(r),
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins(), txGen)
	}
}
