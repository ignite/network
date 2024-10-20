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
					Short:     "List all chains",
				},
				{
					RpcMethod:      "GetChain",
					Use:            "get-chain [launch-id]",
					Short:          "Get a Chain by id",
					Alias:          []string{"show-chain"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
				},
				{
					RpcMethod:      "ListGenesisAccount",
					Use:            "list-genesis-account [launch-id]",
					Short:          "List all genesis accounts for a launch",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
				},
				{
					RpcMethod:      "GetGenesisAccount",
					Use:            "get-genesis-account [launch-id] [address]",
					Short:          "Get the genesis account for a launch",
					Alias:          []string{"show-genesis-account"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListGenesisValidator",
					Use:            "list-genesis-validator [launch-id]",
					Short:          "List all genesis validators",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
				},
				{
					RpcMethod:      "GetGenesisValidator",
					Use:            "get-genesis-validator [launch-id] [address]",
					Short:          "Get a genesis validator",
					Alias:          []string{"show-genesis-validator"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListVestingAccount",
					Use:            "list-vesting-account [launch-id]",
					Short:          "list all vesting accounts for a launch",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
				},
				{
					RpcMethod:      "GetVestingAccount",
					Use:            "get-vesting-account [launch-id] [address]",
					Short:          "shows the vesting account for a launch",
					Alias:          []string{"show-vesting-account"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListRequest",
					Use:            "list-request [launch-id]",
					Short:          "List all requests",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
				},
				{
					RpcMethod:      "GetRequest",
					Use:            "get-request [launch-id] [request-id]",
					Short:          "Get a request",
					Alias:          []string{"show-request"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}, {ProtoField: "request_id"}},
				},
				{
					RpcMethod: "ListParamChange",
					Use:       "list-param-change",
					Short:     "List all param changes",
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
					RpcMethod:      "EditChain",
					Use:            "edit-chain [launch-id] [set-project-id] [project-id] [metadata]",
					Short:          "Edit chain information",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"project_id": {Name: "project_id", Usage: "Set the project ID if the chain is not associated with a project"},
						"metadata":   {Name: "metadata", Usage: "Set metadata field for the chain"},
					},
				},
				{
					RpcMethod:      "SettleRequest",
					Use:            "settle-request [approve|reject] [launch-id] [request-id]",
					Short:          "Approve or reject a pending request",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "approve"}, {ProtoField: "launch_id"}, {ProtoField: "request_id"}},
				},
				{
					RpcMethod:      "RevertLaunch",
					Use:            "revert-launch [launch-id]",
					Short:          "Revert the launch of a chain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
