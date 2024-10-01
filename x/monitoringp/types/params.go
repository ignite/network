package types

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/ignite/network/pkg/chainid"
	networktypes "github.com/ignite/network/pkg/types"
)

var (
	DefaultLastBlockHeight int64 = 1
	DefaultConsumerChainID       = "spn-1"
)

// NewParams creates a new Params instance.
func NewParams(
	lastBlockHeight int64,
	consumerChainID string,
	ccs networktypes.ConsensusState,
	consumerUnbondingpPeriod int64,
	consumerRevisionHeight uint64,
) Params {
	return Params{
		LastBlockHeight:         lastBlockHeight,
		ConsumerConsensusState:  ccs,
		ConsumerChainID:         consumerChainID,
		ConsumerUnbondingPeriod: consumerUnbondingpPeriod,
		ConsumerRevisionHeight:  consumerRevisionHeight,
	}
}

// DefaultParams returns a default set of parameters.
// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultLastBlockHeight,
		DefaultConsumerChainID,
		networktypes.ConsensusState{},
		networktypes.DefaultUnbondingPeriod,
		networktypes.DefaultRevisionHeight,
	)
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateLastBlockHeight(p.LastBlockHeight); err != nil {
		return err
	}
	if err := validateConsumerChainID(p.ConsumerChainID); err != nil {
		return err
	}
	if err := validateConsumerConsensusState(p.ConsumerConsensusState); err != nil {
		return err
	}
	if err := validateConsumerUnbondingPeriod(p.ConsumerUnbondingPeriod); err != nil {
		return err
	}
	return validateConsumerRevisionHeight(p.ConsumerRevisionHeight)
}

// validateLastBlockHeight validates last block height
func validateLastBlockHeight(i interface{}) error {
	lastBlockHeight, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if lastBlockHeight <= 0 {
		return errors.New("last block height can't be 0 or negative")
	}

	return nil
}

// validateConsumerConsensusState validates consumer consensus state
func validateConsumerConsensusState(i interface{}) error {
	ccs, ok := i.(networktypes.ConsensusState)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// perform the verification only if the Consumer Consensus State is defined
	// TODO: remove this check and set an official SPN mainnet consensus state as default
	if ccs.Timestamp != "" {
		tmConsensusState, err := ccs.ToTendermintConsensusState()
		if err != nil {
			return errors.Wrap(err, "consumer consensus state can't be converted")
		}
		if err := tmConsensusState.ValidateBasic(); err != nil {
			return errors.Wrap(err, "invalid consumer consensus state")
		}
	}
	return nil
}

// validateConsumerChainID validates consumer chain ID
func validateConsumerChainID(i interface{}) error {
	chainID, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, _, err := chainid.ParseGenesisChainID(chainID)
	if err != nil {
		return errors.Wrap(err, "invalid chain ID param")
	}
	return nil
}

// validateConsumerUnbondingPeriod validates consumer unbonding period
func validateConsumerUnbondingPeriod(i interface{}) error {
	unbondingPeriod, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if unbondingPeriod < networktypes.MinimalUnbondingPeriod {
		return fmt.Errorf("minimal unbonding period is %d", networktypes.MinimalUnbondingPeriod)
	}

	return nil
}

// validateConsumerRevisionHeight validates consumer revision height
func validateConsumerRevisionHeight(i interface{}) error {
	revisionHeight, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if revisionHeight == 0 {
		return fmt.Errorf("minimal revision height is %d", 1)
	}

	return nil
}
