package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
)

// CreateNewChain creates a new chain in the store from the provided information
func (k Keeper) CreateNewChain(
	ctx context.Context,
	coordinatorID uint64,
	genesisChainID,
	sourceURL,
	sourceHash string,
	initialGenesis types.InitialGenesis,
	hasProject bool,
	projectID uint64,
	isMainnet bool,
	accountBalance sdk.Coins,
	metadata []byte,
) (uint64, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	coordinator, err := k.profileKeeper.GetCoordinator(ctx, coordinatorID)
	if err != nil {
		return 0, sdkerrors.Wrapf(err, "%d", coordinatorID)
	}

	chain := types.Chain{
		CoordinatorID:   coordinatorID,
		GenesisChainID:  genesisChainID,
		CreatedAt:       sdkCtx.BlockTime().Unix(),
		SourceURL:       sourceURL,
		SourceHash:      sourceHash,
		InitialGenesis:  initialGenesis,
		HasProject:      hasProject,
		ProjectID:       projectID,
		IsMainnet:       isMainnet,
		LaunchTriggered: false,
		LaunchTime:      time.Unix(0, 0).UTC(),
		AccountBalance:  accountBalance,
		Metadata:        metadata,
	}

	if err := chain.Validate(); err != nil {
		return 0, err
	}

	// If the chain is associated to a project, project existence and coordinator is checked
	if hasProject {
		project, err := k.projectKeeper.GetProject(ctx, projectID)
		if err != nil {
			return 0, sdkerrors.Wrapf(err, "%d", projectID)
		}
		if project.CoordinatorID != coordinatorID {
			return 0, fmt.Errorf(
				"chain coordinator %d and project coordinator %d don't match",
				coordinatorID,
				project.CoordinatorID,
			)
		}
	}

	// Append the chain to the store
	launchID, err := k.AppendChain(ctx, chain)

	// Register the chain to the project
	if hasProject {
		if err := k.projectKeeper.AddChainToProject(ctx, projectID, launchID); err != nil {
			return 0, err
		}
	}

	return launchID, sdkCtx.EventManager().EmitTypedEvent(&types.EventChainCreated{
		LaunchID:           launchID,
		CoordinatorAddress: coordinator.Address,
		CoordinatorID:      coordinatorID,
	})
}

// AppendChain appends a chain in the store with a new launch id and update the count
func (k Keeper) AppendChain(ctx context.Context, chain types.Chain) (uint64, error) {
	launchID, err := k.ChainSeq.Next(ctx)
	if err != nil {
		return 0, ignterrors.Criticalf("failed to get next chain sequence %s", err.Error())
	}
	chain.LaunchID = launchID
	if err := k.Chain.Set(ctx, launchID, chain); err != nil {
		return 0, ignterrors.Criticalf("chain not set %s", err.Error())
	}
	return launchID, nil
}

func (k Keeper) GetChain(ctx context.Context, launchID uint64) (types.Chain, error) {
	chain, err := k.Chain.Get(ctx, launchID)
	if errors.Is(err, collections.ErrNotFound) {
		return types.Chain{}, types.ErrChainNotFound
	}
	return chain, err
}

// Chains returns all Chain.
func (k Keeper) Chains(ctx context.Context) ([]types.Chain, error) {
	chains := make([]types.Chain, 0)
	err := k.Chain.Walk(ctx, nil, func(_ uint64, chain types.Chain) (bool, error) {
		chains = append(chains, chain)
		return false, nil
	})
	return chains, err
}
