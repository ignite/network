package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
)

// IsProjectMainnetLaunchTriggered returns true if the provided project has an associated mainnet chain whose launch
// has been already triggered
func (k Keeper) IsProjectMainnetLaunchTriggered(ctx context.Context, projectID uint64) (bool, error) {
	project, err := k.GetProject(ctx, projectID)
	if err != nil {
		return false, sdkerrors.Wrapf(err, "%d", projectID)
	}

	if project.MainnetInitialized {
		chain, err := k.launchKeeper.GetChain(ctx, project.MainnetID)
		if err != nil {
			return false, fmt.Errorf("mainnet chain %d for project %d not found", project.MainnetID, projectID)
		}
		if !chain.IsMainnet {
			return false, fmt.Errorf("chain %d for project %d is not a mainnet chain", project.MainnetID, projectID)
		}
		if chain.LaunchTriggered {
			return true, nil
		}
	}
	return false, nil
}
