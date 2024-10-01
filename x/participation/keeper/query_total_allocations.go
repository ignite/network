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

	// TODO: Process the query

	return &types.QueryTotalAllocationsResponse{}, nil
}
