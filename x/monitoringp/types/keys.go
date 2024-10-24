package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "monitoringp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_monitoringp"

	// Version defines the current version the IBC module supports
	Version = "monitoring-1"

	// PortID is the default port id that module binds to
	PortID = "monitoringp"
)

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_monitoringp")

	// PortKey defines the key to store the port ID in store
	PortKey = collections.NewPrefix("monitoringp-port-")

	// MonitoringInfoKey is the prefix to retrieve all MonitoringInfo
	MonitoringInfoKey = collections.NewPrefix("monitoringInfo/value/")

	// ConnectionChannelIDKey is the prefix to retrieve all ConnectionChannelID
	ConnectionChannelIDKey = collections.NewPrefix("connectionChannelID/value/")

	// ConsumerClientIDKey is the prefix to retrieve all ConsumerClientID
	ConsumerClientIDKey = collections.NewPrefix("consumerClientID/value/")
)
