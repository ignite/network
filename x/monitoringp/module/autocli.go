package monitoringp

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/ignite/network/api/network/monitoringp/v1"
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
					RpcMethod: "GetMonitoringInfo",
					Use:       "get-monitoring-info",
					Short:     "Get the monitoring information",
					Alias:     []string{"show-monitoring-info"},
				},
				{
					RpcMethod: "GetConnectionChannelID",
					Use:       "get-connection-channel-id",
					Short:     "get the connection channel ID used for the IBC connection",
					Alias:     []string{"show-connection-channel-id"},
				},
				{
					RpcMethod: "GetConsumerClientID",
					Use:       "get-consumer-client-id",
					Short:     "Get the consumer client ID used for the IBC connection",
					Alias:     []string{"show-consumer-client-id"},
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
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
