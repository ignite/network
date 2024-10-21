package keeper

import (
	"context"

	"cosmossdk.io/collections"
	"github.com/pkg/errors"

	"github.com/ignite/network/x/profile/types"
)

func (k Keeper) GetValidatorByOperatorAddress(ctx context.Context, operatorAddress string) (types.ValidatorByOperatorAddress, error) {
	val, err := k.ValidatorByOperatorAddress.Get(ctx, operatorAddress)
	if errors.Is(err, collections.ErrNotFound) {
		return types.ValidatorByOperatorAddress{}, types.ErrValidatorByOperatorAddressNotFound
	}
	return val, err
}

func (k Keeper) GetValidator(ctx context.Context, address string) (types.Validator, error) {
	acc, err := k.Validator.Get(ctx, address)
	if errors.Is(err, collections.ErrNotFound) {
		return types.Validator{}, types.ErrValidatorNotFound
	}
	return acc, err
}

// ListValidator returns all Validator.
func (k Keeper) ListValidator(ctx context.Context) ([]types.Validator, error) {
	validators := make([]types.Validator, 0)
	err := k.Validator.Walk(ctx, nil, func(_ string, validator types.Validator) (bool, error) {
		validators = append(validators, validator)
		return false, nil
	})
	return validators, err
}
