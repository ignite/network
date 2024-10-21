package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/participation/types"
)

func (q queryServer) ListUsedAllocations(ctx context.Context, req *types.QueryAllUsedAllocationsRequest) (*types.QueryAllUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	usedAllocationss, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.UsedAllocations,
		req.Pagination,
		func(_ string, value types.UsedAllocations) (types.UsedAllocations, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUsedAllocationsResponse{UsedAllocations: usedAllocationss, Pagination: pageRes}, nil
}

func (q queryServer) GetUsedAllocations(ctx context.Context, req *types.QueryGetUsedAllocationsRequest) (*types.QueryGetUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.UsedAllocations.Get(ctx, req.Address)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetUsedAllocationsResponse{UsedAllocations: val}, nil
}
