package profile

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/network/api/network/profile/v1"
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
					RpcMethod: "ListCoordinator",
					Use:       "list-coordinator",
					Short:     "List all Coordinator",
				},
				{
					RpcMethod:      "GetCoordinator",
					Use:            "get-coordinator [id]",
					Short:          "Gets a Coordinator by id",
					Alias:          []string{"show-coordinator"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ID"}},
				},
				{
					RpcMethod: "ListValidator",
					Use:       "list-validator",
					Short:     "List all Validator",
				},
				{
					RpcMethod:      "GetValidator",
					Use:            "get-validator [id]",
					Short:          "Gets a Validator",
					Alias:          []string{"show-validator"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "GetCoordinatorByAddress",
					Use:            "get-coordinator-by-address [address]",
					Short:          "Query GetCoordinatorByAddress",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "GetValidatorByOperatorAddress",
					Use:            "get-validator-by-operator-address [address]",
					Short:          "Query GetValidatorByOperatorAddress",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
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
					RpcMethod:      "CreateCoordinator",
					Use:            "create-coordinator [description] [active]",
					Short:          "Create Coordinator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "description"}, {ProtoField: "active"}},
				},
				{
					RpcMethod:      "UpdateCoordinatorDescription",
					Use:            "update-coordinator [id] [description] [active]",
					Short:          "Update Coordinator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ID"}, {ProtoField: "description"}, {ProtoField: "active"}},
				},
				{
					RpcMethod:      "DisableCoordinator",
					Use:            "delete-coordinator [id]",
					Short:          "Delete Coordinator",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ID"}},
				},
				{
					RpcMethod:      "UpdateCoordinatorAddress",
					Use:            "update-coordinator-address [new-address]",
					Short:          "Send a UpdateCoordinatorAddress tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "newAddress"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
