package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/profile/types"
)

// AppendCoordinator appends a coordinator in the store with a new coordinator id and update the count
func (k Keeper) AppendCoordinator(ctx context.Context, coordinator types.Coordinator) (uint64, error) {
	cordinatorID, err := k.CoordinatorSeq.Next(ctx)
	if err != nil {
		return 0, ignterrors.Criticalf("failed to get next coordinator sequence %s", err.Error())
	}
	coordinator.CoordinatorId = cordinatorID
	if err := k.Coordinator.Set(ctx, cordinatorID, coordinator); err != nil {
		return 0, ignterrors.Criticalf("coordinator not set %s", err.Error())
	}
	return cordinatorID, nil
}

func (k Keeper) CoordinatorIDFromAddress(ctx context.Context, address sdk.AccAddress) (uint64, error) {
	coordinatorByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if errors.Is(err, collections.ErrNotFound) {
		return 0, types.ErrCoordinatorAddressNotFound
	}
	return coordinatorByAddress.CoordinatorId, err
}

func (k Keeper) GetCoordinator(ctx context.Context, coordinatorID uint64) (types.Coordinator, error) {
	acc, err := k.Coordinator.Get(ctx, coordinatorID)
	if errors.Is(err, collections.ErrNotFound) {
		return types.Coordinator{}, types.ErrCoordinatorNotFound
	}
	return acc, err
}

// Coordinators returns all Coordinator.
func (k Keeper) Coordinators(ctx context.Context) ([]types.Coordinator, error) {
	coordinators := make([]types.Coordinator, 0)
	err := k.Coordinator.Walk(ctx, nil, func(_ uint64, coordinator types.Coordinator) (bool, error) {
		coordinators = append(coordinators, coordinator)
		return false, nil
	})
	return coordinators, err
}

// GetCoordinatorByAddress returns the CoordinatorByAddress associated to an address
// returns ErrCoordAddressNotFound if not found in the store
// if the corresponding Coordinator is not found or is inactive, returns ErrCritical
func (k Keeper) GetCoordinatorByAddress(ctx context.Context, address sdk.AccAddress) (types.CoordinatorByAddress, error) {
	coordinatorByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if errors.Is(err, collections.ErrNotFound) {
		return types.CoordinatorByAddress{}, sdkerrors.Wrapf(types.ErrCoordinatorAddressNotFound, "address: %s", address)
	} else if err != nil {
		return types.CoordinatorByAddress{}, err
	}

	coordinator, err := k.GetCoordinator(ctx, coordinatorByAddress.CoordinatorId)
	if errors.Is(err, types.ErrCoordinatorNotFound) {
		// return critical error
		return types.CoordinatorByAddress{}, ignterrors.Criticalf("a coordinator address is associated to a non-existent coordinator ID: %d",
			coordinatorByAddress.CoordinatorId)
	} else if err != nil {
		return types.CoordinatorByAddress{}, err
	}

	if !coordinator.Active {
		// return critical error
		return types.CoordinatorByAddress{}, ignterrors.Criticalf("a coordinator address is inactive and should not exist in the store: ID: %d",
			coordinatorByAddress.CoordinatorId)
	}

	return coordinatorByAddress, nil
}

// CoordinatorByAddresses returns all CoordinatorByAddress.
func (k Keeper) CoordinatorByAddresses(ctx context.Context) ([]types.CoordinatorByAddress, error) {
	coordinatorByAddresses := make([]types.CoordinatorByAddress, 0)
	err := k.CoordinatorByAddress.Walk(ctx, nil, func(_ sdk.AccAddress, coordinatorByAddress types.CoordinatorByAddress) (bool, error) {
		coordinator, err := k.GetCoordinator(ctx, coordinatorByAddress.CoordinatorId)
		if errors.Is(err, types.ErrCoordinatorNotFound) {
			// return critical error
			return true, ignterrors.Criticalf("a coordinator address is associated to a non-existent coordinator ID: %d",
				coordinatorByAddress.CoordinatorId)
		} else if err != nil {
			return true, err
		}

		if !coordinator.Active {
			// return critical error
			return true, ignterrors.Criticalf("a coordinator address is inactive and should not exist in the store: ID: %d",
				coordinatorByAddress.CoordinatorId)
		}

		coordinatorByAddresses = append(coordinatorByAddresses, coordinatorByAddress)
		return false, nil
	})
	return coordinatorByAddresses, err
}
