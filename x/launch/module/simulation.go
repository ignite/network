package launch

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	launchsimulation "github.com/ignite/network/x/launch/simulation"
	"github.com/ignite/network/x/launch/types"
)

// avoid unused import issue
var _ = sample.AccAddress

const (
	defaultWeightMsgCreateChain              int = 50
	defaultWeightMsgEditChain                int = 20
	defaultWeightMsgRequestAddGenesisAccount int = 50
	defaultWeightMsgRequestAddVestingAccount int = 50
	defaultWeightMsgRequestRemoveAccount     int = 15
	defaultWeightMsgRequestAddValidator      int = 50
	defaultWeightMsgRequestRemoveValidator   int = 15
	defaultWeightMsgRequestParamChange       int = 15
	defaultWeightMsgSettleRequest            int = 50
	defaultWeightMsgTriggerLaunch            int = 15
	defaultWeightMsgRevertLaunch             int = 0
	defaultWeightMsgUpdateLaunchInformation  int = 20

	opWeightMsgCreateChain              = "op_weight_msg_create_chain"
	opWeightMsgEditChain                = "op_weight_msg_edit_chain"
	opWeightMsgRequestAddGenesisAccount = "op_weight_msg_request_add_genesis_account"
	opWeightMsgRequestAddVestingAccount = "op_weight_msg_request_add_vesting_account"
	opWeightMsgRequestRemoveAccount     = "op_weight_msg_request_remove_account"
	opWeightMsgRequestAddValidator      = "op_weight_msg_request_add_validator"
	opWeightMsgRequestRemoveValidator   = "op_weight_msg_request_remove_validator"
	opWeightMsgRequestParamChange       = "op_weight_msg_request_change_param"
	opWeightMsgTriggerLaunch            = "op_weight_msg_trigger_launch"
	opWeightMsgRevertLaunch             = "op_weight_msg_revert_launch"
	opWeightMsgSettleRequest            = "op_weight_msg_settle_request"
	opWeightMsgUpdateLaunchInformation  = "op_weight_msg_update_launch_information"

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	launchGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&launchGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateChain int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateChain, &weightMsgCreateChain, nil,
		func(_ *rand.Rand) {
			weightMsgCreateChain = defaultWeightMsgCreateChain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateChain,
		launchsimulation.SimulateMsgCreateChain(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgEditChain int
	simState.AppParams.GetOrGenerate(opWeightMsgEditChain, &weightMsgEditChain, nil,
		func(_ *rand.Rand) {
			weightMsgEditChain = defaultWeightMsgEditChain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEditChain,
		launchsimulation.SimulateMsgEditChain(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgUpdateLaunchInformation int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateLaunchInformation, &weightMsgUpdateLaunchInformation, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateLaunchInformation = defaultWeightMsgUpdateLaunchInformation
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateLaunchInformation,
		launchsimulation.SimulateMsgUpdateLaunchInformation(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgRequestAddGenesisAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestAddGenesisAccount, &weightMsgRequestAddGenesisAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddGenesisAccount = defaultWeightMsgRequestAddGenesisAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestAddGenesisAccount,
		launchsimulation.SimulateMsgRequestAddGenesisAccount(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgRequestAddVestingAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestAddVestingAccount, &weightMsgRequestAddVestingAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddVestingAccount = defaultWeightMsgRequestAddVestingAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestAddVestingAccount,
		launchsimulation.SimulateMsgRequestAddVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgRequestRemoveAccount int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestRemoveAccount, &weightMsgRequestRemoveAccount, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRemoveAccount = defaultWeightMsgRequestRemoveAccount
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestRemoveAccount,
		launchsimulation.SimulateMsgRequestRemoveAccount(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgRequestAddValidator int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestAddValidator, &weightMsgRequestAddValidator, nil,
		func(_ *rand.Rand) {
			weightMsgRequestAddValidator = defaultWeightMsgRequestAddValidator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestAddValidator,
		launchsimulation.SimulateMsgRequestAddValidator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgRequestRemoveValidator int
	simState.AppParams.GetOrGenerate(opWeightMsgRequestRemoveValidator, &weightMsgRequestRemoveValidator, nil,
		func(_ *rand.Rand) {
			weightMsgRequestRemoveValidator = defaultWeightMsgRequestRemoveValidator
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestRemoveValidator,
		launchsimulation.SimulateMsgRequestRemoveValidator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgSettleRequest int
	simState.AppParams.GetOrGenerate(opWeightMsgSettleRequest, &weightMsgSettleRequest, nil,
		func(_ *rand.Rand) {
			weightMsgSettleRequest = defaultWeightMsgSettleRequest
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSettleRequest,
		launchsimulation.SimulateMsgSettleRequest(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgTriggerLaunch int
	simState.AppParams.GetOrGenerate(opWeightMsgTriggerLaunch, &weightMsgTriggerLaunch, nil,
		func(_ *rand.Rand) {
			weightMsgTriggerLaunch = defaultWeightMsgTriggerLaunch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTriggerLaunch,
		launchsimulation.SimulateMsgTriggerLaunch(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgRevertLaunch int
	simState.AppParams.GetOrGenerate(opWeightMsgRevertLaunch, &weightMsgRevertLaunch, nil,
		func(_ *rand.Rand) {
			weightMsgRevertLaunch = defaultWeightMsgRevertLaunch
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevertLaunch,
		launchsimulation.SimulateMsgRevertLaunch(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateChain,
			defaultWeightMsgCreateChain,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgCreateChain(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgEditChain,
			defaultWeightMsgEditChain,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgEditChain(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateLaunchInformation,
			defaultWeightMsgUpdateLaunchInformation,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgUpdateLaunchInformation(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),

		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestAddGenesisAccount,
			defaultWeightMsgRequestAddGenesisAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgRequestAddGenesisAccount(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestAddVestingAccount,
			defaultWeightMsgRequestAddVestingAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgRequestAddVestingAccount(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestRemoveAccount,
			defaultWeightMsgRequestRemoveAccount,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgRequestRemoveAccount(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestAddValidator,
			defaultWeightMsgRequestAddValidator,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgRequestAddValidator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestRemoveValidator,
			defaultWeightMsgRequestRemoveValidator,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgRequestRemoveValidator(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestParamChange,
			defaultWeightMsgRequestParamChange,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgRequestParamChange(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSettleRequest,
			defaultWeightMsgSettleRequest,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgSettleRequest(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTriggerLaunch,
			defaultWeightMsgTriggerLaunch,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgTriggerLaunch(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRevertLaunch,
			defaultWeightMsgRevertLaunch,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				launchsimulation.SimulateMsgRevertLaunch(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
