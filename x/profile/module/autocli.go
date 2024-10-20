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
					Short:     "List all coordinators",
				},
				{
					RpcMethod:      "GetCoordinator",
					Use:            "get-coordinator [coordinator-id]",
					Short:          "Get a Coordinator by id",
					Alias:          []string{"show-coordinator"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "coordinator_id"}},
				},
				{
					RpcMethod: "ListValidator",
					Use:       "list-validator",
					Short:     "List all validators",
				},
				{
					RpcMethod:      "GetValidator",
					Use:            "get-validator [address]",
					Short:          "Get a validator",
					Alias:          []string{"show-validator"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "GetCoordinatorByAddress",
					Use:            "get-coordinator-by-address [address]",
					Short:          "Get a coordinator by its address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "GetValidatorByOperatorAddress",
					Use:            "get-validator-by-operator-address [operatorAddress]",
					Short:          "Get a validator address by an associated operator address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "operator_address"}},
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
					RpcMethod: "CreateCoordinator",
					Use:       "create-coordinator",
					Short:     "Create a new coordinator profile",
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"identity": {Name: "identity", Usage: "Coordinator identity"},
						"website":  {Name: "website", Usage: "Coordinator website URL"},
						"details":  {Name: "details", Usage: "Coordinator details"},
					},
				},
				{
					RpcMethod: "UpdateCoordinatorDescription",
					Use:       "update-coordinator-description",
					Short:     "Update a coordinator description",
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"identity": {Name: "identity", Usage: "Coordinator identity"},
						"website":  {Name: "website", Usage: "Coordinator website URL"},
						"details":  {Name: "details", Usage: "Coordinator details"},
					},
				},
				{
					RpcMethod: "DisableCoordinator",
					Use:       "disable-coordinator",
					Short:     "Disable the coordinator profile associated to the sender address",
				},
				{
					RpcMethod:      "UpdateCoordinatorAddress",
					Use:            "update-coordinator-address [new-address]",
					Short:          "Update a coordinator address",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "new_address"}},
				},
				{
					RpcMethod: "UpdateValidatorDescription",
					Use:       "update-validator-description",
					Short:     "Update a validator description",
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"identity":         {Name: "identity", Usage: "Validator identity"},
						"moniker":          {Name: "moniker", Usage: "Validator moniker"},
						"website":          {Name: "website", Usage: "Validator website URL"},
						"security_contact": {Name: "security-contact", Usage: "Validator security contact"},
						"details":          {Name: "details", Usage: "Validator details"},
					},
				},
				{
					RpcMethod:      "AddValidatorOperatorAddress",
					Use:            "add-validator-operator-address [operator-address]",
					Short:          "Associate an validator operator address to a validator on SPN",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "operator_address"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
