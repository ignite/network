package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:                           PortID,
		LaunchIdFromChannelIdList:        []LaunchIDFromChannelID{},
		LaunchIdFromVerifiedClientIdList: []LaunchIDFromVerifiedClientID{},
		MonitoringHistoryList:            []MonitoringHistory{},
		VerifiedClientIdList:             []VerifiedClientID{},
		ProviderClientIdList:             []ProviderClientID{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// Check for duplicated index in verifiedClientID
	verifiedClientIDIndexMap := make(map[uint64]struct{})
	clientIDMap := make(map[string]struct{})
	for _, elem := range gs.VerifiedClientIdList {
		if _, ok := verifiedClientIDIndexMap[elem.LaunchId]; ok {
			return fmt.Errorf("duplicated index for verifiedClientID")
		}
		verifiedClientIDIndexMap[elem.LaunchId] = struct{}{}

		// Check for duplicated client id
		for _, clientID := range elem.ClientIdList {
			key := clientIDKey(elem.LaunchId, clientID)
			if _, ok := clientIDMap[key]; ok {
				return fmt.Errorf("duplicated client id")
			}
			clientIDMap[key] = struct{}{}
		}
	}

	// Check for duplicated index in providerClientID
	providerClientIDIndexMap := make(map[uint64]struct{})
	for _, elem := range gs.ProviderClientIdList {
		if _, ok := providerClientIDIndexMap[elem.LaunchId]; ok {
			return fmt.Errorf("duplicated index for providerClientID")
		}
		providerClientIDIndexMap[elem.LaunchId] = struct{}{}

		// Check if the client id exist
		key := clientIDKey(elem.LaunchId, elem.ClientId)
		if _, ok := clientIDMap[key]; !ok {
			return fmt.Errorf("client id from providerClientID list not found")
		}
	}

	// Check for duplicated index in launchIDFromVerifiedClientID
	launchIDFromVerifiedClientIDIndexMap := make(map[string]struct{})
	for _, elem := range gs.LaunchIdFromVerifiedClientIdList {
		if _, ok := launchIDFromVerifiedClientIDIndexMap[elem.ClientId]; ok {
			return fmt.Errorf("duplicated index for launchIDFromVerifiedClientID")
		}
		launchIDFromVerifiedClientIDIndexMap[elem.ClientId] = struct{}{}

		// Check if the client id exist
		key := clientIDKey(elem.LaunchId, elem.ClientId)
		if _, ok := clientIDMap[key]; !ok {
			return fmt.Errorf("client id from launchIDFromVerifiedClientID list not found")
		}
	}

	// Check for duplicated index in launchIDFromChannelID
	launchIDFromChannelIDIndexMap := make(map[string]struct{})
	for _, elem := range gs.LaunchIdFromChannelIdList {
		if _, ok := launchIDFromChannelIDIndexMap[elem.ChannelId]; ok {
			return fmt.Errorf("duplicated index for launchIDFromChannelID")
		}
		launchIDFromChannelIDIndexMap[elem.ChannelId] = struct{}{}
	}

	// Check for duplicated index in monitoringHistory
	monitoringHistoryIndexMap := make(map[uint64]struct{})
	for _, elem := range gs.MonitoringHistoryList {
		if _, ok := monitoringHistoryIndexMap[elem.LaunchId]; ok {
			return fmt.Errorf("duplicated index for monitoringHistory")
		}
		monitoringHistoryIndexMap[elem.LaunchId] = struct{}{}
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// clientIDKey creates a string key for launch id and client id
func clientIDKey(launchID uint64, clientID string) string {
	return fmt.Sprintf("%d-%s", launchID, clientID)
}
