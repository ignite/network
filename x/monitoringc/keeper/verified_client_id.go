package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"

	"github.com/ignite/network/x/monitoringc/types"
)

// AllVerifiedClientID returns all VerifiedClientID.
func (k Keeper) AllVerifiedClientID(ctx context.Context) ([]types.VerifiedClientID, error) {
	verifiedClientIDs := make([]types.VerifiedClientID, 0)
	err := k.VerifiedClientID.Walk(ctx, nil, func(_ uint64, value types.VerifiedClientID) (bool, error) {
		verifiedClientIDs = append(verifiedClientIDs, value)
		return false, nil
	})
	return verifiedClientIDs, err
}

// AllLaunchIDFromVerifiedClientID returns all LaunchIDFromVerifiedClientID.
func (k Keeper) AllLaunchIDFromVerifiedClientID(ctx context.Context) ([]types.LaunchIDFromVerifiedClientID, error) {
	launchIDFromVerifiedClientIDs := make([]types.LaunchIDFromVerifiedClientID, 0)
	err := k.LaunchIDFromVerifiedClientID.Walk(ctx, nil, func(_ string, value types.LaunchIDFromVerifiedClientID) (bool, error) {
		launchIDFromVerifiedClientIDs = append(launchIDFromVerifiedClientIDs, value)
		return false, nil
	})
	return launchIDFromVerifiedClientIDs, err
}

// AddVerifiedClientID add a specific verifiedClientID without duplication in the store from its launch id
func (k Keeper) AddVerifiedClientID(ctx context.Context, launchID uint64, clientID string) error {
	verifiedClientID, err := k.VerifiedClientID.Get(ctx, launchID)
	if errors.Is(err, collections.ErrNotFound) {
		verifiedClientID = types.VerifiedClientID{LaunchID: launchID}
	} else if err != nil {
		return err
	}

	for _, cID := range verifiedClientID.ClientIDs {
		if clientID == cID {
			return nil
		}
	}
	verifiedClientID.ClientIDs = append(verifiedClientID.ClientIDs, clientID)
	return k.VerifiedClientID.Set(ctx, launchID, verifiedClientID)
}
