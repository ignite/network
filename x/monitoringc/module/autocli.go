package monitoringc

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/network/api/network/monitoringc/v1"
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
					RpcMethod: "ListLaunchIDFromChannelID",
					Use:       "list-launch-id-from-channel-id",
					Short:     "List all LaunchIDFromChannelID",
				},
				{
					RpcMethod:      "GetLaunchIDFromChannelID",
					Use:            "get-launch-id-from-channel-id [channel-id]",
					Short:          "Gets a LaunchIDFromChannelID",
					Alias:          []string{"show-launch-id-from-channel-id"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "channelID"}},
				},
				{
					RpcMethod:      "GetMonitoringHistory",
					Use:            "get-monitoring-history [launch-id]",
					Short:          "Gets a MonitoringHistory",
					Alias:          []string{"show-monitoring-history"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				{
					RpcMethod:      "GetVerifiedClientID",
					Use:            "get-verified-client-id [launch-id]",
					Short:          "Gets a VerifiedClientID",
					Alias:          []string{"show-verified-client-id"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
				},
				{
					RpcMethod: "ListProviderClientID",
					Use:       "list-provider-client-id",
					Short:     "List all ProviderClientID",
				},
				{
					RpcMethod:      "GetProviderClientID",
					Use:            "get-provider-client-id [launch-id]",
					Short:          "Gets a ProviderClientID",
					Alias:          []string{"show-provider-client-id"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}},
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
					RpcMethod:      "CreateClient",
					Use:            "create-client [launch-id] [consensus-state] [validator-set] [unbonding-period] [revision-height]",
					Short:          "Send a CreateClient tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "launchID"}, {ProtoField: "consensusState"}, {ProtoField: "validatorSet"}, {ProtoField: "unbondingPeriod"}, {ProtoField: "revisionHeight"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
