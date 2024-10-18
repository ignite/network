package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) ListProject(ctx context.Context, req *types.QueryAllProjectRequest) (*types.QueryAllProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	projects, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Project,
		req.Pagination,
		func(_ uint64, value types.Project) (types.Project, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProjectResponse{Project: projects, Pagination: pageRes}, nil
}

func (q queryServer) GetProject(ctx context.Context, req *types.QueryGetProjectRequest) (*types.QueryGetProjectResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	project, err := q.k.Project.Get(ctx, req.ProjectID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetProjectResponse{Project: project}, nil
}
