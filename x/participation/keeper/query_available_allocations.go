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

	// TODO: Process the query

	return &types.QueryAvailableAllocationsResponse{}, nil
}
