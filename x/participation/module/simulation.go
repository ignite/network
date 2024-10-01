package participation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	participationsimulation "github.com/ignite/network/x/participation/simulation"
	"github.com/ignite/network/x/participation/types"
)

// avoid unused import issue
var _ = sample.AccAddress

const (
	defaultWeightMsgParticipate         int = 50
	defaultWeightMsgWithdrawAllocations int = 50

	opWeightMsgParticipate         = "op_weight_msg_participate"
	opWeightMsgWithdrawAllocations = "op_weight_withdraw_allocations"

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	participationGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&participationGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgParticipate int
	simState.AppParams.GetOrGenerate(opWeightMsgParticipate, &weightMsgParticipate, nil,
		func(_ *rand.Rand) {
			weightMsgParticipate = defaultWeightMsgParticipate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgParticipate,
		participationsimulation.SimulateMsgParticipate(am.accountKeeper, am.bankKeeper, am.fundraisingKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgWithdrawAllocations int
	simState.AppParams.GetOrGenerate(opWeightMsgWithdrawAllocations, &weightMsgWithdrawAllocations, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawAllocations = defaultWeightMsgWithdrawAllocations
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawAllocations,
		participationsimulation.SimulateMsgWithdrawAllocations(am.accountKeeper, am.bankKeeper, am.fundraisingKeeper, am.keeper, simState.TxConfig),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgParticipate,
			defaultWeightMsgParticipate,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				participationsimulation.SimulateMsgParticipate(am.accountKeeper, am.bankKeeper, am.fundraisingKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdrawAllocations,
			defaultWeightMsgWithdrawAllocations,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				participationsimulation.SimulateMsgWithdrawAllocations(am.accountKeeper, am.bankKeeper, am.fundraisingKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
