package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"

	"github.com/ignite/network/x/monitoringc/types"
)

// AllVerifiedClientID returns all VerifiedClientID.
func (k Keeper) AllVerifiedClientID(ctx context.Context) ([]types.VerifiedClientID, error) {
	verifiedClientIdList := make([]types.VerifiedClientID, 0)
	err := k.VerifiedClientID.Walk(ctx, nil, func(_ uint64, value types.VerifiedClientID) (bool, error) {
		verifiedClientIdList = append(verifiedClientIdList, value)
		return false, nil
	})
	return verifiedClientIdList, err
}

// AllLaunchIDFromVerifiedClientID returns all LaunchIDFromVerifiedClientID.
func (k Keeper) AllLaunchIDFromVerifiedClientID(ctx context.Context) ([]types.LaunchIDFromVerifiedClientID, error) {
	launchIDFromVerifiedClientIdList := make([]types.LaunchIDFromVerifiedClientID, 0)
	err := k.LaunchIDFromVerifiedClientID.Walk(ctx, nil, func(_ string, value types.LaunchIDFromVerifiedClientID) (bool, error) {
		launchIDFromVerifiedClientIdList = append(launchIDFromVerifiedClientIdList, value)
		return false, nil
	})
	return launchIDFromVerifiedClientIdList, err
}

// AddVerifiedClientID add a specific verifiedClientID without duplication in the store from its launch id
func (k Keeper) AddVerifiedClientID(ctx context.Context, launchID uint64, clientID string) error {
	verifiedClientID, err := k.VerifiedClientID.Get(ctx, launchID)
	if errors.Is(err, collections.ErrNotFound) {
		verifiedClientID = types.VerifiedClientID{LaunchId: launchID}
	} else if err != nil {
		return err
	}

	for _, cID := range verifiedClientID.ClientIdList {
		if clientID == cID {
			return nil
		}
	}
	verifiedClientID.ClientIdList = append(verifiedClientID.ClientIdList, clientID)
	return k.VerifiedClientID.Set(ctx, launchID, verifiedClientID)
}
