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
					Use:       "list-auction-used-allocations [address]",
					Short:     "List all used allocations for auctions for an address",
				},
				{
					RpcMethod:      "GetAuctionUsedAllocations",
					Use:            "get-auction-used-allocations [address] [auction-id]",
					Short:          "Get used allocations for an auction",
					Alias:          []string{"show-auction-used-allocations"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "auction_id"}},
				},
				{
					RpcMethod: "ListUsedAllocations",
					Use:       "list-used-allocations",
					Short:     "List all used allocations",
				},
				{
					RpcMethod:      "GetUsedAllocations",
					Use:            "get-used-allocations [address]",
					Short:          "Get total used allocations for an address",
					Alias:          []string{"show-used-allocations"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "TotalAllocations",
					Use:            "total-allocations [address]",
					Short:          "Get the total allocations available for an account",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				{
					RpcMethod:      "AvailableAllocations",
					Use:            "available-allocations [address]",
					Short:          "Get the available, unused allocations for an account",
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
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}, {ProtoField: "tier_id"}},
				},
				{
					RpcMethod:      "WithdrawAllocations",
					Use:            "withdraw-allocations [auction-id]",
					Short:          "Send a WithdrawAllocations tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auction_id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
