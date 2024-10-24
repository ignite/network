package project

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/network/api/network/project/v1"
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
					RpcMethod:      "ListMainnetAccount",
					Use:            "list-mainnet-account",
					Short:          "List all mainnet accounts for a project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}},
				},
				{
					RpcMethod:      "GetMainnetAccount",
					Use:            "get-mainnet-account [project-id] [address]",
					Short:          "Get the mainnet account for a project",
					Alias:          []string{"show-mainnet-account"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "address"}},
				},
				{
					RpcMethod: "ListProject",
					Use:       "list-project",
					Short:     "List all projects",
				},
				{
					RpcMethod:      "GetProject",
					Use:            "get-project [project-id]",
					Short:          "Get a Project by id",
					Alias:          []string{"show-project"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}},
				},
				{
					RpcMethod:      "GetProjectChains",
					Use:            "get-project-chains [project-id]",
					Short:          "Query the coin balance for special allocations",
					Alias:          []string{"show-project-chains"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}},
				},
				{
					RpcMethod:      "SpecialAllocationsBalance",
					Use:            "special-allocations-balance [project-id]",
					Short:          "Query SpecialAllocationsBalance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}},
				},
				{
					RpcMethod:      "MainnetAccountBalance",
					Use:            "mainnet-account-balance [project-id] [address]",
					Short:          "Get the mainnet account balance for a project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListMainnetAccountBalance",
					Use:            "list-mainnet-account-balance [project-id]",
					Short:          "List all mainnet account balances for a project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}},
				},
				{
					RpcMethod: "TotalShares",
					Use:       "total-shares",
					Short:     "Get the total-shares value of projects",
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
					RpcMethod:      "CreateProject",
					Use:            "create-project [project-name] [total-supply]",
					Short:          "Create a new project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_name"}, {ProtoField: "total_supply"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"metadata": {Name: "metadata", Usage: "Set metadata field for the project"},
					},
				},
				{
					RpcMethod:      "EditProject",
					Use:            "edit-project [project-id]",
					Short:          "Edit the project name or metadata",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"name":     {Name: "name", Usage: "Set name for the project"},
						"metadata": {Name: "metadata", Usage: "Set metadata field for the project"},
					},
				},
				{
					RpcMethod:      "UpdateTotalSupply",
					Use:            "update-total-supply [project-id] [total-supply]",
					Short:          "Update the total supply of the mainnet of a project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "total_supply_update"}},
				},
				{
					RpcMethod:      "InitializeMainnet",
					Use:            "initialize-mainnet [project-id] [source-url] [source-hash] [mainnet-chain-id]",
					Short:          "Initialize the mainnet of the project to open gentxs submissions and fix total supply",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "source_url"}, {ProtoField: "source_hash"}, {ProtoField: "mainnet_chain_id"}},
				},
				{
					RpcMethod:      "MintVouchers",
					Use:            "mint-vouchers [project-id] [shares]",
					Short:          "Mint vouchers from project shares",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "shares"}},
				},
				{
					RpcMethod:      "BurnVouchers",
					Use:            "burn-vouchers [project-id] [vouchers]",
					Short:          "Burn vouchers and decrease allocated shares of the project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "vouchers"}},
				},
				{
					RpcMethod:      "RedeemVouchers",
					Use:            "redeem-vouchers [project-id] [vouchers]",
					Short:          "Redeem vouchers and allocate shares for an account in the mainnet of the project",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "vouchers"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"account": {Name: "account", Usage: "Account address that receives shares allocation from redeemed vouchers"},
					},
				},
				{
					RpcMethod:      "UnredeemVouchers",
					Use:            "unredeem-vouchers [project-id] [shares]",
					Short:          "Unredeem vouchers that have been redeemed into an account and get vouchers back",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "project_id"}, {ProtoField: "shares"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
