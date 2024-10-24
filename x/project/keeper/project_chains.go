package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	"github.com/ignite/network/x/project/types"
)

// AddChainToProject adds a new chain into an existing project.
func (k Keeper) AddChainToProject(ctx context.Context, projectID, launchID uint64) error {
	// Check project exist
	if _, err := k.GetProject(ctx, projectID); err != nil {
		return fmt.Errorf("project %d not found", projectID)
	}

	projectChains, err := k.GetProjectChains(ctx, projectID)
	if errors.Is(err, types.ErrProjectChainsNotFound) {
		projectChains = types.ProjectChains{
			ProjectId: projectID,
			Chains:    []uint64{launchID},
		}
	} else if err != nil {
		return err
	} else {
		// Ensure no duplicated chain ID
		for _, existingChainID := range projectChains.Chains {
			if existingChainID == launchID {
				return fmt.Errorf("chain %d already associated to project %d", launchID, projectID)
			}
		}
		projectChains.Chains = append(projectChains.Chains, launchID)
	}
	if err := k.ProjectChains.Set(ctx, projectChains.ProjectId, projectChains); err != nil {
		return err
	}
	return sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventProjectChainAdded{
		ProjectId: projectID,
		LaunchId:  launchID,
	})
}

func (k Keeper) GetProjectChains(ctx context.Context, projectID uint64) (types.ProjectChains, error) {
	projectChains, err := k.ProjectChains.Get(ctx, projectID)
	if errors.Is(err, collections.ErrNotFound) {
		return types.ProjectChains{}, types.ErrProjectChainsNotFound
	}
	return projectChains, err
}
