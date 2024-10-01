package types

import (
	"fmt"

	networktypes "github.com/ignite/network/pkg/types"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MainnetAccountList: []MainnetAccount{},
		ProjectList:        []Project{},
		ProjectChainsList:  []ProjectChains{},
		TotalShares:        networktypes.TotalShareNumber,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in mainnetAccount
	mainnetAccountIndexMap := make(map[string]struct{})

	for _, elem := range gs.MainnetAccountList {
		index := fmt.Sprint(elem.ProjectID)
		if _, ok := mainnetAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mainnetAccount")
		}
		mainnetAccountIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in project
	projectIDMap := make(map[uint64]bool)
	projectCount := gs.GetProjectCount()
	for _, elem := range gs.ProjectList {
		if _, ok := projectIDMap[elem.ProjectID]; ok {
			return fmt.Errorf("duplicated id for project")
		}
		if elem.ProjectID >= projectCount {
			return fmt.Errorf("project id should be lower or equal than the last id")
		}
		projectIDMap[elem.ProjectID] = true
	}
	// Check for duplicated index in projectChains
	projectChainsIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProjectChainsList {
		index := fmt.Sprint(elem.ProjectID)
		if _, ok := projectChainsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for projectChains")
		}
		projectChainsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
