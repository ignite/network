package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) GetProjectChains(ctx context.Context, req *types.QueryGetProjectChainsRequest) (*types.QueryGetProjectChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.ProjectChains.Get(ctx, req.ProjectId)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetProjectChainsResponse{ProjectChains: val}, nil
}
