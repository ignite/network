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

// SimulateMsgCreateProject simulates a MsgCreateProject message
func SimulateMsgCreateProject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgCreateProject{}

		simAccount, _, found := GetCoordSimAccount(r, ctx, pk, accs)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip project creation"), nil, nil
		}

		params, err := k.Params.Get(ctx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "params not found"), nil, err
		}
		creationFee := params.ProjectCreationFee

		msg = types.NewMsgCreateProject(
			simAccount.Address.String(),
			sample.ProjectName(r),
			sample.TotalSupply(r),
			sample.Metadata(r, 20),
		)

		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, creationFee, txGen)
	}
}
