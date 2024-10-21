package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/profile/types"
)

const coordinatorIDNotFoundRoute = "coordinator-id-not-found"

// RegisterInvariants registers all module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, coordinatorIDNotFoundRoute,
		CoordinatorAddrNotFoundInvariant(k))
}

// CoordinatorAddrNotFoundInvariant invariant that checks if
// the `CoordinateByAddress` is associated with a coordinator
func CoordinatorAddrNotFoundInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		all, err := k.CoordinatorByAddresses(ctx)
		if err != nil {
			return sdk.FormatInvariant(types.ModuleName, coordinatorIDNotFoundRoute, err.Error()), true
		}
		for _, coordinatorByAddres := range all {
			if _, err := k.GetCoordinator(ctx, coordinatorByAddres.CoordinatorId); err != nil {
				return sdk.FormatInvariant(
					types.ModuleName, coordinatorIDNotFoundRoute,
					fmt.Sprintf("%s: %d", err, coordinatorByAddres.CoordinatorId),
				), true
			}
		}
		return "", false
	}
}
