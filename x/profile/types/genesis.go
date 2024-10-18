package types

import (
	"errors"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                      DefaultParams(),
		ValidatorList:               []Validator{},
		ValidatorsByOperatorAddress: []ValidatorByOperatorAddress{},
		CoordinatorList:             []Coordinator{},
		CoordinatorsByAddress:       []CoordinatorByAddress{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := gs.ValidateValidators(); err != nil {
		return err
	}

	if err := gs.ValidateCoordinators(); err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

func (gs GenesisState) ValidateValidators() error {
	// Check for duplicated index in validator
	validatorIndexMap := make(map[string]Validator)
	for _, elem := range gs.ValidatorList {
		valIndex := elem.Address
		if _, ok := validatorIndexMap[valIndex]; ok {
			return errors.New("duplicated index for validator")
		}
		validatorIndexMap[valIndex] = elem
	}

	// Check for duplicated index in validatorByOperatorAddress
	validatorByOperatorAddressIndexMap := make(map[string]struct{})
	for _, elem := range gs.ValidatorsByOperatorAddress {
		index := elem.OperatorAddress
		if _, ok := validatorByOperatorAddressIndexMap[index]; ok {
			return errors.New("duplicated index for validatorByOperatorAddress")
		}
		valIndex := elem.ValidatorAddress
		validator, ok := validatorIndexMap[valIndex]
		if !ok {
			return errors.New("validator operator address not found for Validator")
		}
		if !validator.HasOperatorAddress(elem.OperatorAddress) {
			return errors.New("operator address not found in the Validator operator address list")
		}
		validatorByOperatorAddressIndexMap[index] = struct{}{}
	}

	return nil
}

func (gs GenesisState) ValidateCoordinators() error {
	// Check for duplicated index in coordinatorByAddress
	coordinatorByAddressIndexMap := make(map[string]uint64)
	for _, elem := range gs.CoordinatorsByAddress {
		index := elem.Address
		if _, ok := coordinatorByAddressIndexMap[index]; ok {
			return errors.New("duplicated index for coordinatorByAddress")
		}
		coordinatorByAddressIndexMap[index] = elem.CoordinatorID
	}

	// Check for duplicated ID in coordinator or if coordinator is inactive
	coordinatorIDMap := make(map[uint64]bool)
	counter := gs.CoordinatorCount
	for _, elem := range gs.CoordinatorList {
		if _, ok := coordinatorIDMap[elem.CoordinatorID]; ok {
			return errors.New("duplicated id for coordinator")
		}
		if elem.CoordinatorID >= counter {
			return errors.New("coordinator id should be lower or equal than the last id")
		}
		index := elem.Address
		_, found := coordinatorByAddressIndexMap[index]

		switch {
		case !found && elem.Active:
			return errors.New("coordinator address not found for CoordinatorByAddress")
		case found && !elem.Active:
			return errors.New("coordinator found by CoordinatorByAddress should not be inactive")
		}

		coordinatorIDMap[elem.CoordinatorID] = true

		// Remove to check if all coordinator by address exist
		delete(coordinatorByAddressIndexMap, index)
	}
	// Check if all coordinator by address exist
	if len(coordinatorByAddressIndexMap) > 0 {
		return errors.New("coordinator address not found for coordinatorID")
	}
	return nil
}
