package types

import (
	"errors"
	"fmt"
	"time"

	sdkmath "cosmossdk.io/math"
)

var (
	DefaultAllocationPrice = AllocationPrice{
		Bonded: sdkmath.NewInt(1000),
	}
	DefaultParticipationTierList = []Tier{
		{
			TierId:              1,
			RequiredAllocations: sdkmath.OneInt(),
			Benefits: TierBenefits{
				MaxBidAmount: sdkmath.NewInt(1000),
			},
		},
		{
			TierId:              2,
			RequiredAllocations: sdkmath.NewInt(2),
			Benefits: TierBenefits{
				MaxBidAmount: sdkmath.NewInt(2000),
			},
		},
		{
			TierId:              3,
			RequiredAllocations: sdkmath.NewInt(5),
			Benefits: TierBenefits{
				MaxBidAmount: sdkmath.NewInt(10000),
			},
		},
		{
			TierId:              4,
			RequiredAllocations: sdkmath.NewInt(10),
			Benefits: TierBenefits{
				MaxBidAmount: sdkmath.NewInt(30000),
			},
		},
	}

	// DefaultRegistrationPeriod is set to be 1/3 of the default staking UnbondingTime of 21 days.
	DefaultRegistrationPeriod = time.Hour * 24 * 7 // One week
	// DefaultWithdrawalDelay is set to be 2/3 of the default staking UnbondingTime of 21 days. Together with
	// DefaultRegistrationPeriod they sum up to the total default UnbondingTime
	DefaultWithdrawalDelay = time.Hour * 24 * 14 // Two weeks
)

// NewParams creates a new Params instance
func NewParams(
	allocationPrice AllocationPrice,
	participationTierList []Tier,
	registrationPeriod,
	withdrawalDelay time.Duration,
) Params {
	return Params{
		AllocationPrice:       allocationPrice,
		ParticipationTierList: participationTierList,
		RegistrationPeriod:    registrationPeriod,
		WithdrawalDelay:       withdrawalDelay,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultAllocationPrice,
		DefaultParticipationTierList,
		DefaultRegistrationPeriod,
		DefaultWithdrawalDelay,
	)
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateAllocationPrice(p.AllocationPrice); err != nil {
		return err
	}

	if err := validateParticipationTierList(p.ParticipationTierList); err != nil {
		return err
	}

	if err := validateTimeDuration(p.RegistrationPeriod); err != nil {
		return err
	}

	return validateTimeDuration(p.WithdrawalDelay)
}

// validateAllocationPrice validates the AllocationPrice param
func validateAllocationPrice(allocationPrice AllocationPrice) error {
	if allocationPrice.Bonded.IsNil() {
		return errors.New("value for 'bonded' should be set")
	}

	if !allocationPrice.Bonded.IsPositive() {
		return errors.New("value for 'bonded' must be greater than zero")
	}

	return nil
}

// validateParticipationTierList validates the ParticipationTierList param
func validateParticipationTierList(participationTierList []Tier) error {
	tiersIndexMap := make(map[uint64]struct{})
	for _, tier := range participationTierList {
		// check IDs are unique
		if _, ok := tiersIndexMap[tier.TierId]; ok {
			return fmt.Errorf("duplicated tier ID: %v", tier.TierId)
		}
		tiersIndexMap[tier.TierId] = struct{}{}

		if tier.RequiredAllocations.LTE(sdkmath.ZeroInt()) {
			return errors.New("required allocations must be greater than zero")
		}

		if err := validateTierBenefits(tier.Benefits); err != nil {
			return err
		}
	}

	return nil
}

func validateTierBenefits(b TierBenefits) error {
	if b.MaxBidAmount.IsNil() {
		return errors.New("max bid amount should be set")
	}

	if !b.MaxBidAmount.IsPositive() {
		return fmt.Errorf("max bid amount must be greater than zero")
	}

	return nil
}

// validateTimeDuration validates a time.Duration parameter
func validateTimeDuration(v time.Duration) error {
	if v <= 0 {
		return fmt.Errorf("time frame must be positive")
	}

	return nil
}

func (a *AllocationPrice) Equal(cmp *AllocationPrice) bool {
	return a.Bonded.Equal(cmp.Bonded)
}

func (t *Tier) Equal(cmp *Tier) bool {
	switch {
	case t.TierId != cmp.TierId:
		return false
	case !t.RequiredAllocations.Equal(cmp.RequiredAllocations):
		return false
	case !t.Benefits.Equal(&cmp.Benefits):
		return false
	}
	return true
}

func (t *TierBenefits) Equal(cmp *TierBenefits) bool {
	return t.MaxBidAmount.Equal(cmp.MaxBidAmount)
}
