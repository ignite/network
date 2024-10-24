package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"github.com/pkg/errors"

	"github.com/ignite/network/x/project/types"
)

// AppendProject appends a project in the store with a new id and update the count
func (k Keeper) AppendProject(ctx context.Context, project types.Project) (uint64, error) {
	projectID, err := k.ProjectSeq.Next(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get next project sequence %s", err.Error())
	}
	project.ProjectId = projectID
	if err := k.Project.Set(ctx, projectID, project); err != nil {
		return 0, fmt.Errorf("project not set %s", err.Error())
	}
	return projectID, nil
}

func (k Keeper) GetProject(ctx context.Context, projectID uint64) (types.Project, error) {
	project, err := k.Project.Get(ctx, projectID)
	if errors.Is(err, collections.ErrNotFound) {
		return types.Project{}, types.ErrProjectNotFound
	}
	return project, err
}

// Projects returns all Project.
func (k Keeper) Projects(ctx context.Context) ([]types.Project, error) {
	projects := make([]types.Project, 0)
	err := k.Project.Walk(ctx, nil, func(_ uint64, project types.Project) (bool, error) {
		projects = append(projects, project)
		return false, nil
	})
	return projects, err
}
