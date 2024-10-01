package types

import (
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// DefaultMinLaunchTime ...
	// TODO: set back this value to the default one
	// time.Hour * 24
	DefaultMinLaunchTime = time.Hour * 5
	DefaultMaxLaunchTime = time.Hour * 24 * 7

	// DefaultRevertDelay is the delay after the launch time when it is possible to revert the launch of the chain
	// launch can be reverted on-chain when the actual chain launch failed (incorrect gentx, etc...)
	// This delay must be small be big enough to ensure nodes had the time to bootstrap\
	// This currently corresponds to 1 hour
	DefaultRevertDelay = time.Hour

	DefaultFee = sdk.Coins(nil) // EmptyCoins

	MaxParametrableLaunchTime  = time.Hour * 24 * 31
	MaxParametrableRevertDelay = time.Hour * 24

	DefaultMaxMetadataLength uint64 = 2000
)

// NewLaunchTimeRange creates a new LaunchTimeRange instance
func NewLaunchTimeRange(minLaunchTime, maxLaunchTime time.Duration) LaunchTimeRange {
	return LaunchTimeRange{
		MinLaunchTime: minLaunchTime,
		MaxLaunchTime: maxLaunchTime,
	}
}

// NewParams creates a new Params instance
func NewParams(
	minLaunchTime,
	maxLaunchTime,
	revertDelay time.Duration,
	chainCreationFee,
	requestFee sdk.Coins,
	maxMetadataLength uint64,
) Params {
	return Params{
		LaunchTimeRange:   NewLaunchTimeRange(minLaunchTime, maxLaunchTime),
		RevertDelay:       revertDelay,
		ChainCreationFee:  chainCreationFee,
		RequestFee:        requestFee,
		MaxMetadataLength: maxMetadataLength,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinLaunchTime,
		DefaultMaxLaunchTime,
		DefaultRevertDelay,
		DefaultFee,
		DefaultFee,
		DefaultMaxMetadataLength,
	)
}

// Validate validates the set of params.
func (p Params) Validate() error {
	if err := validateLaunchTimeRange(p.LaunchTimeRange); err != nil {
		return err
	}
	if err := validateRevertDelay(p.RevertDelay); err != nil {
		return err
	}
	if err := validateMaxMetadataLength(p.MaxMetadataLength); err != nil {
		return err
	}
	if err := p.ChainCreationFee.Validate(); err != nil {
		return err
	}
	return p.RequestFee.Validate()
}

func validateLaunchTimeRange(i interface{}) error {
	v, ok := i.(LaunchTimeRange)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// it is enough to check that minLaunchTime is positive since it must be that minLaunchTime < maxLaunchTime
	if v.MinLaunchTime < 0 {
		return errors.New("MinLaunchTime can't be negative")
	}

	if v.MinLaunchTime > v.MaxLaunchTime {
		return errors.New("MinLaunchTime can't be higher than MaxLaunchTime")
	}

	// just need to check max launch time due to check above that guarantees correctness of the range
	if v.MaxLaunchTime > MaxParametrableLaunchTime {
		return errors.New("max parametrable launch time reached")
	}
	return nil
}

func validateRevertDelay(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v > MaxParametrableRevertDelay {
		return errors.New("max parametrable revert delay reached")
	}

	if v <= 0 {
		return errors.New("revert delay parameter must be positive")
	}

	return nil
}

func validateRequestFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return v.Validate()
}

func validateMaxMetadataLength(i interface{}) error {
	if _, ok := i.(uint64); !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func (l *LaunchTimeRange) Equal(cmp *LaunchTimeRange) bool {
	switch {
	case l.GetMinLaunchTime().Nanoseconds() != cmp.GetMinLaunchTime().Nanoseconds():
		return false
	case l.GetMaxLaunchTime().Nanoseconds() != cmp.GetMaxLaunchTime().Nanoseconds():
		return false
	}
	return true
}
