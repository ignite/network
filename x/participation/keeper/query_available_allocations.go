package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/participation/types"
)

func (q queryServer) AvailableAllocations(ctx context.Context, req *types.QueryAvailableAllocationsRequest) (*types.QueryAvailableAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	availableAlloc, err := q.k.GetAvailableAllocations(ctx, req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &types.QueryAvailableAllocationsResponse{AvailableAllocations: availableAlloc}, nil
}
