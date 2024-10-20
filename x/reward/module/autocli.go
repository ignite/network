package reward

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/network/api/network/reward/v1"
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
					RpcMethod: "ListRewardPool",
					Use:       "list-reward-pool",
					Short:     "List all reward pools",
				},
				{
					RpcMethod:      "GetRewardPool",
					Use:            "get-reward-pool [launch-id]",
					Short:          "Get the reward pool for a launch",
					Alias:          []string{"show-reward-pool"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}},
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
					RpcMethod:      "SetRewards",
					Use:            "set-rewards [launch-id] [coins] [last-reward-height]",
					Short:          "Set rewards for being validator of a chain",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launch_id"}, {ProtoField: "coins"}, {ProtoField: "last_reward_height"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
