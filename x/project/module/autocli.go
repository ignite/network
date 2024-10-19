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
					Short:          "List all MainnetAccount",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}},
				},
				{
					RpcMethod:      "GetMainnetAccount",
					Use:            "get-mainnet-account [project-id] [address]",
					Short:          "Gets a MainnetAccount",
					Alias:          []string{"show-mainnet-account"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}, {ProtoField: "address"}},
				},
				{
					RpcMethod: "ListProject",
					Use:       "list-project",
					Short:     "List all Project",
				},
				{
					RpcMethod:      "GetProject",
					Use:            "get-project [project-id]",
					Short:          "Gets a Project by id",
					Alias:          []string{"show-project"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}},
				},
				{
					RpcMethod:      "GetProjectChains",
					Use:            "get-project-chains [project-id]",
					Short:          "Gets a ProjectChains",
					Alias:          []string{"show-project-chains"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}},
				},
				{
					RpcMethod:      "SpecialAllocationsBalance",
					Use:            "special-allocations-balance [project-id]",
					Short:          "Query SpecialAllocationsBalance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}},
				},
				{
					RpcMethod:      "MainnetAccountBalance",
					Use:            "mainnet-account-balance [project-id] [address]",
					Short:          "Query MainnetAccountBalance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "ListMainnetAccountBalance",
					Use:            "list-mainnet-account-balance [project-id]",
					Short:          "Query ListMainnetAccountBalance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}},
				},
				{
					RpcMethod:      "TotalShares",
					Use:            "total-shares",
					Short:          "Query TotalShares",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					Use:            "create-project [project-name] [total-supply] [metadata]",
					Short:          "Send a CreateProject tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectName"}, {ProtoField: "totalSupply"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "EditProject",
					Use:            "edit-project [project-id] [name] [metadata]",
					Short:          "Send a EditProject tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectId"}, {ProtoField: "name"}, {ProtoField: "metadata"}},
				},
				{
					RpcMethod:      "UpdateTotalSupply",
					Use:            "update-total-supply [project-id] [total-supply-update]",
					Short:          "Send a UpdateTotalSupply tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectId"}, {ProtoField: "totalSupplyUpdate"}},
				},
				{
					RpcMethod:      "UpdateSpecialAllocations",
					Use:            "update-special-allocations [project-id] [special-allocations]",
					Short:          "Send a UpdateSpecialAllocations tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectId"}, {ProtoField: "specialAllocations"}},
				},
				{
					RpcMethod:      "InitializeMainnet",
					Use:            "initialize-mainnet [project-id] [source-url] [source-hash] [mainnet-chain-id]",
					Short:          "Send a InitializeMainnet tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}, {ProtoField: "sourceURL"}, {ProtoField: "sourceHash"}, {ProtoField: "mainnetChainID"}},
				},
				{
					RpcMethod:      "MintVouchers",
					Use:            "mint-vouchers [project-id] [shares]",
					Short:          "Send a MintVouchers tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}, {ProtoField: "shares"}},
				},
				{
					RpcMethod:      "BurnVouchers",
					Use:            "burn-vouchers [project-id] [vouchers]",
					Short:          "Send a BurnVouchers tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}, {ProtoField: "vouchers"}},
				},
				{
					RpcMethod:      "RedeemVouchers",
					Use:            "redeem-vouchers [project-id] [account] [vouchers]",
					Short:          "Send a RedeemVouchers tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}, {ProtoField: "account"}, {ProtoField: "vouchers"}},
				},
				{
					RpcMethod:      "UnredeemVouchers",
					Use:            "unredeem-vouchers [project-id] [shares]",
					Short:          "Send a UnredeemVouchers tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "projectID"}, {ProtoField: "shares"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
