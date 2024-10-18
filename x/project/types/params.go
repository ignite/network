package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultMinTotalSupply            = sdkmath.NewInt(100)                   // One hundred
	DefaultMaxTotalSupply            = sdkmath.NewInt(1_000_000_000_000_000) // One Quadrillion
	DefaultProjectCreationFee        = sdk.Coins(nil)                        // EmptyCoins
	DefaultMaxMetadataLength  uint64 = 2000
)

// NewTotalSupplyRange creates a new TotalSupplyRange instance
func NewTotalSupplyRange(minTotalSupply, maxTotalSupply sdkmath.Int) TotalSupplyRange {
	return TotalSupplyRange{
		MinTotalSupply: minTotalSupply,
		MaxTotalSupply: maxTotalSupply,
	}
}

// DefaultTotalSupplyRange returns default total supply range.
func DefaultTotalSupplyRange() TotalSupplyRange {
	return NewTotalSupplyRange(
		DefaultMinTotalSupply,
		DefaultMaxTotalSupply,
	)
}

// NewParams creates a new Params instance
func NewParams(
	minTotalSupply,
	maxTotalSupply sdkmath.Int,
	projectCreationFee sdk.Coins,
	maxMetadataLength uint64,
) Params {
	return Params{
		TotalSupplyRange:   NewTotalSupplyRange(minTotalSupply, maxTotalSupply),
		ProjectCreationFee: projectCreationFee,
		MaxMetadataLength:  maxMetadataLength,
	}
}

// DefaultParams returns default project parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinTotalSupply,
		DefaultMaxTotalSupply,
		DefaultProjectCreationFee,
		DefaultMaxMetadataLength,
	)
}

// Validate validates the set of params.
func (p Params) Validate() error {
	if err := validateTotalSupplyRange(p.TotalSupplyRange); err != nil {
		return err
	}
	if err := validateProjectCreationFee(p.ProjectCreationFee); err != nil {
		return err
	}
	return p.ProjectCreationFee.Validate()
}

func validateTotalSupplyRange(v TotalSupplyRange) error {
	if err := v.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

func validateProjectCreationFee(v sdk.Coins) error {
	return v.Validate()
}
