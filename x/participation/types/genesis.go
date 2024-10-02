package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AuctionUsedAllocationsList: []AuctionUsedAllocations{},
		UsedAllocationsList:        []UsedAllocations{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in auctionUsedAllocations
	auctionUsedAllocationsIndexMap := make(map[string]struct{})
	for _, elem := range gs.AuctionUsedAllocationsList {
		if _, ok := auctionUsedAllocationsIndexMap[elem.Address]; ok {
			return fmt.Errorf("duplicated index for auctionUsedAllocations")
		}
		auctionUsedAllocationsIndexMap[elem.Address] = struct{}{}
	}

	// Check for duplicated index in usedAllocations
	usedAllocationsIndexMap := make(map[string]struct{})
	for _, elem := range gs.UsedAllocationsList {
		if _, ok := usedAllocationsIndexMap[elem.Address]; ok {
			return fmt.Errorf("duplicated index for usedAllocations")
		}
		usedAllocationsIndexMap[elem.Address] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
