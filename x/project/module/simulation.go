package project

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ignite/network/testutil/sample"
	projectsimulation "github.com/ignite/network/x/project/simulation"
	"github.com/ignite/network/x/project/types"
)

// avoid unused import issue
var _ = sample.AccAddress

const (
	defaultWeightMsgCreateProject            = 25
	defaultWeightMsgEditProject              = 20
	defaultWeightMsgUpdateTotalSupply        = 20
	defaultWeightMsgInitializeMainnet        = 15
	defaultWeightMsgUpdateSpecialAllocations = 20
	defaultWeightMsgMintVouchers             = 20
	defaultWeightMsgBurnVouchers             = 20
	defaultWeightMsgRedeemVouchers           = 20
	defaultWeightMsgUnredeemVouchers         = 20

	opWeightMsgCreateProject            = "op_weight_msg_create_project"
	opWeightMsgEditProject              = "op_weight_msg_edit_project"
	opWeightMsgUpdateTotalSupply        = "op_weight_msg_update_total_supply"
	opWeightMsgInitializeMainnet        = "op_weight_msg_initialize_mainnet"
	opWeightMsgUpdateSpecialAllocations = "op_weight_msg_update_special_allocations"
	opWeightMsgMintVouchers             = "op_weight_msg_mint_vouchers"
	opWeightMsgBurnVouchers             = "op_weight_msg_burn_vouchers"
	opWeightMsgRedeemVouchers           = "op_weight_msg_redeem_vouchers"
	opWeightMsgUnredeemVouchers         = "op_weight_msg_unredeem_vouchers"

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	projectGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&projectGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProject int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = defaultWeightMsgCreateProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProject,
		projectsimulation.SimulateMsgCreateProject(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgEditProject int
	simState.AppParams.GetOrGenerate(opWeightMsgEditProject, &weightMsgEditProject, nil,
		func(_ *rand.Rand) {
			weightMsgEditProject = defaultWeightMsgEditProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEditProject,
		projectsimulation.SimulateMsgEditProject(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgUpdateTotalSupply int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTotalSupply, &weightMsgUpdateTotalSupply, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTotalSupply = defaultWeightMsgUpdateTotalSupply
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTotalSupply,
		projectsimulation.SimulateMsgUpdateTotalSupply(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgUpdateSpecialAllocations int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateSpecialAllocations, &weightMsgUpdateSpecialAllocations, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateSpecialAllocations = defaultWeightMsgUpdateSpecialAllocations
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateSpecialAllocations,
		projectsimulation.SimulateMsgUpdateSpecialAllocations(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgInitializeMainnet int
	simState.AppParams.GetOrGenerate(opWeightMsgInitializeMainnet, &weightMsgInitializeMainnet, nil,
		func(_ *rand.Rand) {
			weightMsgInitializeMainnet = defaultWeightMsgInitializeMainnet
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitializeMainnet,
		projectsimulation.SimulateMsgInitializeMainnet(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgMintVouchers int
	simState.AppParams.GetOrGenerate(opWeightMsgMintVouchers, &weightMsgMintVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgMintVouchers = defaultWeightMsgMintVouchers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintVouchers,
		projectsimulation.SimulateMsgMintVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgBurnVouchers int
	simState.AppParams.GetOrGenerate(opWeightMsgBurnVouchers, &weightMsgBurnVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgBurnVouchers = defaultWeightMsgBurnVouchers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnVouchers,
		projectsimulation.SimulateMsgBurnVouchers(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgRedeemVouchers int
	simState.AppParams.GetOrGenerate(opWeightMsgRedeemVouchers, &weightMsgRedeemVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgRedeemVouchers = defaultWeightMsgRedeemVouchers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRedeemVouchers,
		projectsimulation.SimulateMsgRedeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	var weightMsgUnredeemVouchers int
	simState.AppParams.GetOrGenerate(opWeightMsgUnredeemVouchers, &weightMsgUnredeemVouchers, nil,
		func(_ *rand.Rand) {
			weightMsgUnredeemVouchers = defaultWeightMsgUnredeemVouchers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnredeemVouchers,
		projectsimulation.SimulateMsgUnredeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateProject,
			defaultWeightMsgCreateProject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgCreateProject(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgEditProject,
			defaultWeightMsgEditProject,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgEditProject(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTotalSupply,
			defaultWeightMsgUpdateTotalSupply,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgUpdateTotalSupply(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateSpecialAllocations,
			defaultWeightMsgUpdateSpecialAllocations,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgUpdateSpecialAllocations(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitializeMainnet,
			defaultWeightMsgInitializeMainnet,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgInitializeMainnet(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgMintVouchers,
			defaultWeightMsgMintVouchers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgMintVouchers(am.accountKeeper, am.bankKeeper, am.profileKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBurnVouchers,
			defaultWeightMsgBurnVouchers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgBurnVouchers(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRedeemVouchers,
			defaultWeightMsgRedeemVouchers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgRedeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUnredeemVouchers,
			defaultWeightMsgUnredeemVouchers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				projectsimulation.SimulateMsgUnredeemVouchers(am.accountKeeper, am.bankKeeper, am.keeper, simState.TxConfig)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
