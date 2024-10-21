package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "profile"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_profile"
)

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_profile")

	// CoordinatorKey is the prefix to retrieve all Coordinator
	CoordinatorKey = collections.NewPrefix("coordinator/value/")
	// CoordinatorCountKey is the prefix to retrieve all CoordinatorCount
	CoordinatorCountKey = collections.NewPrefix("coordinator/count/")
	// CoordinatorByAddressKey is the prefix to retrieve all CoordinatorByAddress
	CoordinatorByAddressKey = collections.NewPrefix("coordinatorByAddress/value/")

	// ValidatorKey is the prefix to retrieve all Validator
	ValidatorKey = collections.NewPrefix("Validator/value/")

	// ValidatorByOperatorAddressKey is the prefix to retrieve all ValidatorByOperatorAddress
	ValidatorByOperatorAddressKey = collections.NewPrefix("ValidatorByOperatorAddress/value/")
)
