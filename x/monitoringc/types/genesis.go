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
		index := fmt.Sprint(elem.ChannelID)
		if _, ok := launchIDFromChannelIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for launchIDFromChannelID")
		}
		launchIDFromChannelIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in launchIDFromVerifiedClientID
	launchIDFromVerifiedClientIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.LaunchIDFromVerifiedClientIDList {
		index := fmt.Sprint(elem.ClientID)
		if _, ok := launchIDFromVerifiedClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for launchIDFromVerifiedClientID")
		}
		launchIDFromVerifiedClientIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in monitoringHistory
	monitoringHistoryIndexMap := make(map[string]struct{})

	for _, elem := range gs.MonitoringHistoryList {
		index := fmt.Sprint(elem.LaunchID)
		if _, ok := monitoringHistoryIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for monitoringHistory")
		}
		monitoringHistoryIndexMap[index] = struct{}{}
	}

	// Check for duplicated index in verifiedClientID
	verifiedClientIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.VerifiedClientIDList {
		index := fmt.Sprint(elem.LaunchID)
		if _, ok := verifiedClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for verifiedClientID")
		}
		verifiedClientIDIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in providerClientID
	providerClientIDIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProviderClientIDList {
		index := fmt.Sprint(elem.LaunchID)
		if _, ok := providerClientIDIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for providerClientID")
		}
		providerClientIDIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
