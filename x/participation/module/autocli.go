package participation

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/network/api/network/participation/v1"
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
					RpcMethod: "ListAuctionUsedAllocations",
					Use:       "list-auction-used-allocations",
					Short:     "List all AuctionUsedAllocations",
				},
				{
					RpcMethod:      "GetAuctionUsedAllocations",
					Use:            "get-auction-used-allocations [id]",
					Short:          "Gets a AuctionUsedAllocations",
					Alias:          []string{"show-auction-used-allocations"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod: "ListUsedAllocations",
					Use:       "list-used-allocations",
					Short:     "List all UsedAllocations",
				},
				{
					RpcMethod:      "GetUsedAllocations",
					Use:            "get-used-allocations [id]",
					Short:          "Gets a UsedAllocations",
					Alias:          []string{"show-used-allocations"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "TotalAllocations",
					Use:            "total-allocations [address]",
					Short:          "Query TotalAllocations",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},

				{
					RpcMethod:      "AvailableAllocations",
					Use:            "available-allocations [address]",
					Short:          "Query AvailableAllocations",
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
					RpcMethod:      "Participate",
					Use:            "participate [auction-id] [tier-id]",
					Short:          "Send a Participate tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auctionID"}, {ProtoField: "tierID"}},
				},
				{
					RpcMethod:      "WithdrawAllocations",
					Use:            "withdraw-allocations [auction-id]",
					Short:          "Send a WithdrawAllocations tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auctionID"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
