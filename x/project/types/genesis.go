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
		if err := elem.Validate(gs.TotalShares); err != nil {
			return fmt.Errorf("invalid project %d: %s", elem.ProjectID, err.Error())
		}
		projectIDMap[elem.ProjectID] = true
	}

	// Check for duplicated index in projectChains
	projectChainsIndexMap := make(map[uint64]struct{})
	for _, elem := range gs.ProjectChainsList {
		if _, ok := projectIDMap[elem.ProjectID]; !ok {
			return fmt.Errorf("project id %d doesn't exist for chains", elem.ProjectID)
		}
		if _, ok := projectChainsIndexMap[elem.ProjectID]; ok {
			return fmt.Errorf("duplicated index for projectChains")
		}
		projectChainsIndexMap[elem.ProjectID] = struct{}{}
	}

	// Check for duplicated index in mainnetAccount
	mainnetAccountIndexMap := make(map[uint64]struct{})
	for _, elem := range gs.MainnetAccountList {
		if _, ok := projectIDMap[elem.ProjectID]; !ok {
			return fmt.Errorf("project id %d doesn't exist for mainnet account %s",
				elem.ProjectID, elem.Address)
		}
		if _, ok := mainnetAccountIndexMap[elem.ProjectID]; ok {
			return fmt.Errorf("duplicated index for mainnetAccount")
		}
		mainnetAccountIndexMap[elem.ProjectID] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
