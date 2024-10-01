package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) SpecialAllocationsBalance(ctx context.Context, req *types.QuerySpecialAllocationsBalanceRequest) (*types.QuerySpecialAllocationsBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QuerySpecialAllocationsBalanceResponse{}, nil
}
