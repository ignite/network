package reward

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	rewardsimulation "github.com/ignite/network/x/reward/simulation"
	"github.com/ignite/network/x/reward/types"
)

// avoid unused import issue
var _ = sample.AccAddress

const (
	opWeightMsgSetRewards          = "op_weight_msg_set_rewards"
	defaultWeightMsgSetRewards int = 50

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rewardGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rewardGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSetRewards int
	simState.AppParams.GetOrGenerate(opWeightMsgSetRewards, &weightMsgSetRewards, nil,
		func(_ *rand.Rand) {
			weightMsgSetRewards = defaultWeightMsgSetRewards
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetRewards,
		rewardsimulation.SimulateMsgSetRewards(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetRewards,
			defaultWeightMsgSetRewards,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rewardsimulation.SimulateMsgSetRewards(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
