package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) MainnetAccountBalance(ctx context.Context, req *types.QueryMainnetAccountBalanceRequest) (*types.QueryMainnetAccountBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryMainnetAccountBalanceResponse{}, nil
}

func (q queryServer) ListMainnetAccountBalance(ctx context.Context, req *types.QueryListMainnetAccountBalanceRequest) (*types.QueryListMainnetAccountBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryListMainnetAccountBalanceResponse{}, nil
}
