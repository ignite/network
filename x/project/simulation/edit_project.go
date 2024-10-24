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

func SimulateMsgEditProject(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	pk types.ProfileKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msg := &types.MsgEditProject{}

		simAccount, prjtID, found := GetCoordSimAccountWithProjectID(r, ctx, pk, k, accs, false, false)
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "skip edit project"), nil, nil
		}

		var newName string
		var newMetadata []byte

		name := r.Intn(100) < 50
		metadata := r.Intn(100) < 50
		// ensure there is always a value to edit
		if !name && !metadata {
			metadata = true
		}

		if name {
			newName = sample.ProjectName(r)
		}
		if metadata {
			newMetadata = sample.Metadata(r, 20)
		}

		msg = types.NewMsgEditProject(
			simAccount.Address.String(),
			prjtID,
			newName,
			newMetadata,
		)
		return deliverSimTx(r, app, ctx, ak, bk, simAccount, msg, sdk.NewCoins(), txGen)
	}
}
