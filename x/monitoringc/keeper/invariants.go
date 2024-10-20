package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/monitoringc/types"
)

const (
	missingVerifiedClientIDRoute = "missing-verified-client-id"
)

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k *Keeper) {
	ir.RegisterRoute(types.ModuleName, missingVerifiedClientIDRoute,
		MissingVerifiedClientIDInvariant(k))
}

// AllInvariants runs all invariants of the module.
func AllInvariants(k *Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		return MissingVerifiedClientIDInvariant(k)(ctx)
	}
}

// MissingVerifiedClientIDInvariant checks if any of the clientIDs in `VerifiedClientID` does not have a corresponding
// entry in `LaunchIDFromVerifiedClientID`
func MissingVerifiedClientIDInvariant(k *Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		allVerifiedClientID, err := k.AllVerifiedClientID(ctx)
		if err != nil {
			return "", false
		}
		allLaunchIDFromVerifiedClientID, err := k.AllLaunchIDFromVerifiedClientID(ctx)
		if err != nil {
			return "", false
		}
		clientIDMap := make(map[string]struct{})
		for _, launchIDFromVerifiedClientID := range allLaunchIDFromVerifiedClientID {
			clientIDMap[clientIDKey(launchIDFromVerifiedClientID.LaunchId, launchIDFromVerifiedClientID.ClientId)] = struct{}{}
		}
		for _, verifiedClientID := range allVerifiedClientID {
			for _, clientID := range verifiedClientID.ClientIdList {
				if _, ok := clientIDMap[clientIDKey(verifiedClientID.LaunchId, clientID)]; !ok {
					return sdk.FormatInvariant(
						types.ModuleName, missingVerifiedClientIDRoute,
						"client id from verifiedClient list not found in launchIDFromVerifiedClientID list",
					), true
				}
			}
		}
		return "", false
	}
}

// clientIDKey creates a string key for launch id and client id
func clientIDKey(launchID uint64, clientID string) string {
	return fmt.Sprintf("%d-%s", launchID, clientID)
}
