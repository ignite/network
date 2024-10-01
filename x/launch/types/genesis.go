package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ChainList:            []Chain{},
		GenesisAccountList:   []GenesisAccount{},
		GenesisValidatorList: []GenesisValidator{},
		VestingAccountList:   []VestingAccount{},
		RequestList:          []Request{},
		RequestCounters:      []RequestCounter{},
		ParamChangeList:      []ParamChange{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	launchIDMap, err := validateChains(gs)
	if err != nil {
		return err
	}

	if err := validateRequests(gs, launchIDMap); err != nil {
		return err
	}

	if err := validateAccounts(gs, launchIDMap); err != nil {
		return err
	}

	// Check for duplicated index in paramChange
	paramChangeIndexMap := make(map[string]struct{})

	for _, elem := range gs.ParamChangeList {
		index := fmt.Sprint(elem.LaunchID)
		if _, ok := paramChangeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for paramChange")
		}
		paramChangeIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

func validateChains(gs GenesisState) (map[uint64]struct{}, error) {
	// Check for duplicated index in chain
	counter := gs.GetChainCount()
	launchIDMap := make(map[uint64]struct{})
	for _, elem := range gs.ChainList {
		if err := elem.Validate(); err != nil {
			return nil, fmt.Errorf("invalid chain %d: %s", elem.LaunchID, err.Error())
		}

		launchID := elem.LaunchID
		if _, ok := launchIDMap[launchID]; ok {
			return nil, fmt.Errorf("duplicated launch ID for chain")
		}
		launchIDMap[launchID] = struct{}{}

		if elem.LaunchID >= counter {
			return nil, fmt.Errorf("launch id %d should be lower or equal than the last id %d",
				elem.LaunchID, counter)
		}
	}

	return launchIDMap, nil
}

func validateRequests(gs GenesisState, launchIDMap map[uint64]struct{}) error {
	// We checkout request counts to perform verification
	requestCounterMap := make(map[uint64]uint64)
	for _, elem := range gs.RequestCounters {
		if _, ok := requestCounterMap[elem.LaunchID]; ok {
			return fmt.Errorf("duplicated request counter")
		}
		requestCounterMap[elem.LaunchID] = elem.Counter

		// Each genesis account must be associated with an existing chain
		if _, ok := launchIDMap[elem.LaunchID]; !ok {
			return fmt.Errorf("request counter to a non-existing chain: %d",
				elem.LaunchID,
			)
		}
	}

	// Check for duplicated index in request
	requestIndexMap := make(map[string]struct{})
	for _, elem := range gs.RequestList {
		index := fmt.Sprint(elem.LaunchID, elem.RequestID)
		if _, ok := requestIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for request")
		}
		requestIndexMap[index] = struct{}{}

		// Each request pool must be associated with an existing chain
		if _, ok := launchIDMap[elem.LaunchID]; !ok {
			return fmt.Errorf("a request pool is associated to a non-existing chain: %d",
				elem.LaunchID,
			)
		}

		// Check the request counter of the associated chain is not below the request ID
		requestCounter, ok := requestCounterMap[elem.LaunchID]
		if !ok {
			return fmt.Errorf("chain %d has requests but no request counter",
				elem.LaunchID,
			)
		}
		if elem.RequestID >= requestCounter {
			return fmt.Errorf("chain %d contains a request with an ID above the request counter: %d >= %d",
				elem.LaunchID,
				elem.RequestID,
				requestCounter,
			)
		}
	}

	return nil
}

func validateAccounts(gs GenesisState, launchIDMap map[uint64]struct{}) error {
	// Check for duplicated index in genesisAccount
	genesisAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.GenesisAccountList {
		index := fmt.Sprint(elem.LaunchID, elem.Address)
		if _, ok := genesisAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for genesisAccount")
		}
		genesisAccountIndexMap[index] = struct{}{}

		// Each genesis account must be associated with an existing chain
		if _, ok := launchIDMap[elem.LaunchID]; !ok {
			return fmt.Errorf("account %s is associated to a non-existing chain: %d",
				elem.Address,
				elem.LaunchID,
			)
		}
	}

	// Check for duplicated index in vestingAccount
	vestingAccountIndexMap := make(map[string]struct{})
	for _, elem := range gs.VestingAccountList {
		index := fmt.Sprint(elem.LaunchID, elem.Address)
		if _, ok := vestingAccountIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for vestingAccount")
		}
		vestingAccountIndexMap[index] = struct{}{}

		// Each vesting account must be associated with an existing chain
		if _, ok := launchIDMap[elem.LaunchID]; !ok {
			return fmt.Errorf("account %s is associated to a non-existing chain: %d",
				elem.Address,
				elem.LaunchID,
			)
		}

		// An address cannot be defined as a genesis account and a vesting account for the same chain
		accountIndex := fmt.Sprint(elem.LaunchID, elem.Address)
		if _, ok := genesisAccountIndexMap[accountIndex]; ok {
			return fmt.Errorf("account %s can't be a genesis account and a vesting account at the same time for the chain: %d",
				elem.Address,
				elem.LaunchID,
			)
		}
	}

	// Check for duplicated index in genesisValidator
	genesisValidatorIndexMap := make(map[string]struct{})
	for _, elem := range gs.GenesisValidatorList {
		index := fmt.Sprint(elem.LaunchID, elem.Address)
		if _, ok := genesisValidatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for genesisValidator")
		}
		genesisValidatorIndexMap[index] = struct{}{}

		// Each genesis validator must be associated with an existing chain
		if _, ok := launchIDMap[elem.LaunchID]; !ok {
			return fmt.Errorf("validator %s is associated to a non-existing chain: %d",
				elem.Address,
				elem.LaunchID,
			)
		}
	}

	return nil
}
