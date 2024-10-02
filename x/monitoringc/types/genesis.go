package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortID:                           PortID,
		LaunchIDFromChannelIDList:        []LaunchIDFromChannelID{},
		LaunchIDFromVerifiedClientIDList: []LaunchIDFromVerifiedClientID{},
		MonitoringHistoryList:            []MonitoringHistory{},
		VerifiedClientIDList:             []VerifiedClientID{},
		ProviderClientIDList:             []ProviderClientID{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortID); err != nil {
		return err
	}
	// Check for duplicated index in launchIDFromChannelID
	launchIDFromChannelIDIndexMap := make(map[string]struct{})
	for _, elem := range gs.LaunchIDFromChannelIDList {
		if _, ok := launchIDFromChannelIDIndexMap[elem.ChannelID]; ok {
			return fmt.Errorf("duplicated index for launchIDFromChannelID")
		}
		launchIDFromChannelIDIndexMap[elem.ChannelID] = struct{}{}
	}

	// Check for duplicated index in launchIDFromVerifiedClientID
	launchIDFromVerifiedClientIDIndexMap := make(map[string]struct{})
	for _, elem := range gs.LaunchIDFromVerifiedClientIDList {
		if _, ok := launchIDFromVerifiedClientIDIndexMap[elem.ClientID]; ok {
			return fmt.Errorf("duplicated index for launchIDFromVerifiedClientID")
		}
		launchIDFromVerifiedClientIDIndexMap[elem.ClientID] = struct{}{}
	}

	// Check for duplicated index in monitoringHistory
	monitoringHistoryIndexMap := make(map[uint64]struct{})
	for _, elem := range gs.MonitoringHistoryList {
		if _, ok := monitoringHistoryIndexMap[elem.LaunchID]; ok {
			return fmt.Errorf("duplicated index for monitoringHistory")
		}
		monitoringHistoryIndexMap[elem.LaunchID] = struct{}{}
	}

	// Check for duplicated index in verifiedClientID
	verifiedClientIDIndexMap := make(map[uint64]struct{})
	for _, elem := range gs.VerifiedClientIDList {
		if _, ok := verifiedClientIDIndexMap[elem.LaunchID]; ok {
			return fmt.Errorf("duplicated index for verifiedClientID")
		}
		verifiedClientIDIndexMap[elem.LaunchID] = struct{}{}
	}

	// Check for duplicated index in providerClientID
	providerClientIDIndexMap := make(map[uint64]struct{})
	for _, elem := range gs.ProviderClientIDList {
		if _, ok := providerClientIDIndexMap[elem.LaunchID]; ok {
			return fmt.Errorf("duplicated index for providerClientID")
		}
		providerClientIDIndexMap[elem.LaunchID] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
