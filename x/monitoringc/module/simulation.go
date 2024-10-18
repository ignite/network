package monitoringc

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	monitoringcsimulation "github.com/ignite/network/x/monitoringc/simulation"
	"github.com/ignite/network/x/monitoringc/types"
)

// avoid unused import issue
var _ = sample.AccAddress

const (
	opWeightMsgCreateClient          = "op_weight_msg_create_client"
	defaultWeightMsgCreateClient int = 50

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	monitoringcGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortID: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&monitoringcGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateClient int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateClient, &weightMsgCreateClient, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClient = defaultWeightMsgCreateClient
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateClient,
		monitoringcsimulation.SimulateMsgCreateClient(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateClient,
			defaultWeightMsgCreateClient,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				monitoringcsimulation.SimulateMsgCreateClient(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
