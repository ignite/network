package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "monitoringc"

	// FullModuleName defines the full module name used in interface like CLI to make it more explanatory
	FullModuleName = "monitoring-consumer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_monitoringc"

	// Version defines the current version the IBC module supports
	Version = "monitoringc-1"

	// PortID is the default port id that module binds to
	PortID = "monitoringc"
)

var (
	ParamsKey = collections.NewPrefix("p_monitoringc")

	// PortKey defines the key to store the port ID in store
	PortKey = collections.NewPrefix("monitoringc-port-")

	// LaunchIDFromChannelIDKey is the prefix to retrieve all LaunchIDFromChannelID
	LaunchIDFromChannelIDKey = collections.NewPrefix("LaunchIDFromChannelID/value/")

	// LaunchIDFromVerifiedClientIDKey is the prefix to retrieve all LaunchIDFromVerifiedClientID
	LaunchIDFromVerifiedClientIDKey = collections.NewPrefix("LaunchIDFromVerifiedClientID/value/")

	// MonitoringHistoryKey is the prefix to retrieve all MonitoringHistory
	MonitoringHistoryKey = collections.NewPrefix("MonitoringHistory/value/")

	// VerifiedClientIDKey is the prefix to retrieve all VerifiedClientID
	VerifiedClientIDKey = collections.NewPrefix("VerifiedClientID/value/")

	// ProviderClientIDKey is the prefix to retrieve all ProviderClientID
	ProviderClientIDKey = collections.NewPrefix("ProviderClientID/value/")
)
