package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "reward"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_reward"
)

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_reward")

	// RewardPoolKey is the prefix to retrieve all RewardPool
	RewardPoolKey = collections.NewPrefix("RewardPool/value/")
)
