package launch

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/network/api/network/launch/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListChain",
					Use:       "list-chain",
					Short:     "List all Chain",
				},
				{
					RpcMethod:      "GetChain",
					Use:            "get-chain [launch-id]",
					Short:          "Gets a Chain by id",
					Alias:          []string{"show-chain"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				{
					RpcMethod:      "ListGenesisAccount",
					Use:            "list-genesis-account",
					Short:          "List all GenesisAccount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				{
					RpcMethod:      "GetGenesisAccount",
					Use:            "get-genesis-account [id]",
					Short:          "Gets a GenesisAccount",
					Alias:          []string{"show-genesis-account"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListGenesisValidator",
					Use:            "list-genesis-validator",
					Short:          "List all GenesisValidator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				{
					RpcMethod:      "GetGenesisValidator",
					Use:            "get-genesis-validator [id]",
					Short:          "Gets a GenesisValidator",
					Alias:          []string{"show-genesis-validator"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListVestingAccount",
					Use:            "list-vesting-account",
					Short:          "List all VestingAccount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				{
					RpcMethod:      "GetVestingAccount",
					Use:            "get-vesting-account [id]",
					Short:          "Gets a VestingAccount",
					Alias:          []string{"show-vesting-account"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListRequest",
					Use:            "list-request",
					Short:          "List all Request",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				{
					RpcMethod:      "GetRequest",
					Use:            "get-request [id]",
					Short:          "Gets a Request by id",
					Alias:          []string{"show-request"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "requestID"}},
				},
				{
					RpcMethod: "ListParamChange",
					Use:       "list-param-change",
					Short:     "List all ParamChange",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateChain",
					Use:            "create-chain [genesis-chain-id] [source-url] [source-hash] [initial-genesis] [has-project] [project-id] [account-balance] [metadata]",
					Short:          "Send a CreateChain tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "genesisChainID"}, {ProtoField: "sourceURL"}, {ProtoField: "sourceHash"}, {ProtoField: "initialGenesis"}, {ProtoField: "hasProject"}, {ProtoField: "projectID"}, {ProtoField: "accountBalance"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "EditChain",
					Use:            "edit-chain [launch-id] [set-project-id] [project-id] [metadata]",
					Short:          "Send a EditChain tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "setProjectID"}, {ProtoField: "projectID"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateLaunchInformation",
					Use:            "update-launch-information [launch-id] [genesis-chain-id] [source-url] [source-hash] [initial-genesis]",
					Short:          "Send a UpdateLaunchInformation tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "genesisChainID"}, {ProtoField: "sourceURL"}, {ProtoField: "sourceHash"}, {ProtoField: "initialGenesis"}},
				},
				{
					RpcMethod:      "SendRequest",
					Use:            "send-request [launch-id] [content]",
					Short:          "Send a SendRequest tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "content"}},
				},
				{
					RpcMethod:      "SettleRequest",
					Use:            "settle-request [launch-id] [request-id] [approve]",
					Short:          "Send a SettleRequest tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "requestID"}, {ProtoField: "approve"}},
				},
				{
					RpcMethod:      "TriggerLaunch",
					Use:            "trigger-launch [launch-id] [launch-time]",
					Short:          "Send a TriggerLaunch tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "launchTime"}},
				},
				{
					RpcMethod:      "RevertLaunch",
					Use:            "revert-launch [launch-id]",
					Short:          "Send a RevertLaunch tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
