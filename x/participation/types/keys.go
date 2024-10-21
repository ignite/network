package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "participation"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_participation"
)

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_participation")

	// AuctionUsedAllocationsKey is the prefix to retrieve all AuctionUsedAllocations
	AuctionUsedAllocationsKey = collections.NewPrefix("AuctionUsedAllocations/value/")

	// UsedAllocationsKey is the prefix to retrieve all UsedAllocations
	UsedAllocationsKey = collections.NewPrefix("UsedAllocations/value/")
)
