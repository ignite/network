package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/participation/types"
)

func (q queryServer) TotalAllocations(ctx context.Context, req *types.QueryTotalAllocationsRequest) (*types.QueryTotalAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	numAlloc, err := q.k.GetTotalAllocations(ctx, req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryTotalAllocationsResponse{TotalAllocations: numAlloc}, nil
}
