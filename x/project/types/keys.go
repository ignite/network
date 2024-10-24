package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "project"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_project"
)

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_project")

	// ProjectKey is the prefix to retrieve all Project
	ProjectKey = collections.NewPrefix("project/value/")
	// ProjectCountKey is the prefix to retrieve all ProjectCount
	ProjectCountKey = collections.NewPrefix("project/count/")

	// ProjectChainsKey is the prefix to retrieve all ProjectChains
	ProjectChainsKey = collections.NewPrefix("ProjectChains/value/")

	// MainnetAccountKey is the prefix to retrieve all MainnetAccount
	MainnetAccountKey = collections.NewPrefix("MainnetAccount/value/")

	// TotalSharesKey is the prefix to retrieve all TotalShares
	TotalSharesKey = collections.NewPrefix("TotalShares/value/")
)
