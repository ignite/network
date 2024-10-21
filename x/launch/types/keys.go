package types

import (
	"path/filepath"

	"cosmossdk.io/collections"
)

const (
	// ModuleName defines the module name
	ModuleName = "launch"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_launch"
)

func ParamChangeSubKey(module, param string) string {
	return filepath.Join(module, param)
}

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("p_launch")

	// ChainKey is the prefix to retrieve all Chain
	ChainKey = collections.NewPrefix("chain/value/")
	// ChainCountKey is the prefix to retrieve all ChainCount
	ChainCountKey = collections.NewPrefix("chain/count/")

	// RequestKey is the prefix to retrieve all Request
	RequestKey = collections.NewPrefix("request/value/")
	// RequestCountKey is the prefix to retrieve all RequestCount
	RequestCountKey = collections.NewPrefix("request/count/")

	// GenesisAccountKey is the prefix to retrieve all GenesisAccount
	GenesisAccountKey = collections.NewPrefix("GenesisAccount/value/")

	// VestingAccountKey is the prefix to retrieve all VestingAccount
	VestingAccountKey = collections.NewPrefix("VestingAccount/value/")

	// ParamChangeKey is the prefix to retrieve all ParamChange
	ParamChangeKey = collections.NewPrefix("ParamChange/value/")

	// GenesisValidatorKey is the prefix to retrieve all GenesisValidator
	GenesisValidatorKey = collections.NewPrefix("GenesisValidator/value/")
)
