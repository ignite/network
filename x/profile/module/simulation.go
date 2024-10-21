package profile

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	profilesimulation "github.com/ignite/network/x/profile/simulation"
	"github.com/ignite/network/x/profile/types"
)

const (
	defaultWeightMsgUpdateValidatorDescription   = 50
	defaultWeightMsgAddValidatorOperatorAddress  = 50
	defaultWeightMsgCreateCoordinator            = 50
	defaultWeightMsgUpdateCoordinatorDescription = 20
	defaultWeightMsgUpdateCoordinatorAddress     = 20
	defaultWeightMsgDisableCoordinator           = 5

	opWeightMsgUpdateValidatorDescription   = "op_weight_msg_update_validator_description"
	opWeightMsgAddValidatorOperatorAddress  = "op_weight_msg_add_vaildator_operator_address"
	opWeightMsgCreateCoordinator            = "op_weight_msg_create_coordinator"
	opWeightMsgUpdateCoordinatorDescription = "op_weight_msg_update_coordinator_description"
	opWeightMsgUpdateCoordinatorAddress     = "op_weight_msg_update_coordinator_address"
	opWeightMsgDisableCoordinator           = "op_weight_msg_disable_coordinator"

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	profileGenesis := sample.ProfileGenesisState(simState.Rand, accs...)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&profileGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCoordinator int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateCoordinator, &weightMsgCreateCoordinator, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCoordinator = defaultWeightMsgCreateCoordinator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCoordinator,
		profilesimulation.SimulateMsgCreateCoordinator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgUpdateCoordinatorDescription int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateCoordinatorDescription, &weightMsgUpdateCoordinatorDescription, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCoordinatorDescription = defaultWeightMsgUpdateCoordinatorDescription
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCoordinatorDescription,
		profilesimulation.SimulateMsgUpdateCoordinatorDescription(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgDisableCoordinator int
	simState.AppParams.GetOrGenerate(opWeightMsgDisableCoordinator, &weightMsgDisableCoordinator, nil,
		func(_ *rand.Rand) {
			weightMsgDisableCoordinator = defaultWeightMsgDisableCoordinator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDisableCoordinator,
		profilesimulation.SimulateMsgDisableCoordinator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgUpdateCoordinatorAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateCoordinatorAddress, &weightMsgUpdateCoordinatorAddress, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCoordinatorAddress = defaultWeightMsgUpdateCoordinatorAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCoordinatorAddress,
		profilesimulation.SimulateMsgUpdateCoordinatorAddress(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgUpdateValidatorDescription int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateValidatorDescription, &weightMsgUpdateValidatorDescription, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateValidatorDescription = defaultWeightMsgUpdateValidatorDescription
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateValidatorDescription,
		profilesimulation.SimulateMsgUpdateValidatorDescription(am.accountKeeper, am.bankKeeper, simState.TxConfig),
	))

	var weightMsgAddValidatorOperatorAddress int
	simState.AppParams.GetOrGenerate(opWeightMsgAddValidatorOperatorAddress, &weightMsgAddValidatorOperatorAddress, nil,
		func(_ *rand.Rand) {
			weightMsgAddValidatorOperatorAddress = defaultWeightMsgAddValidatorOperatorAddress
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddValidatorOperatorAddress,
		profilesimulation.SimulateMsgAddValidatorOperatorAddress(am.accountKeeper, am.bankKeeper, simState.TxConfig),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateCoordinator,
			defaultWeightMsgCreateCoordinator,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgCreateCoordinator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateCoordinatorDescription,
			defaultWeightMsgUpdateCoordinatorDescription,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgUpdateCoordinatorDescription(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDisableCoordinator,
			defaultWeightMsgDisableCoordinator,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgDisableCoordinator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateCoordinatorAddress,
			defaultWeightMsgUpdateCoordinatorAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgUpdateCoordinatorAddress(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateValidatorDescription,
			defaultWeightMsgUpdateValidatorDescription,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgUpdateValidatorDescription(am.accountKeeper, am.bankKeeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddValidatorOperatorAddress,
			defaultWeightMsgAddValidatorOperatorAddress,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				profilesimulation.SimulateMsgAddValidatorOperatorAddress(am.accountKeeper, am.bankKeeper, simState.TxConfig)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
