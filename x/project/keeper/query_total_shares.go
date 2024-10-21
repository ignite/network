package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) TotalShares(ctx context.Context, req *types.QueryTotalSharesRequest) (*types.QueryTotalSharesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	shares, err := q.k.TotalShares.Get(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can't get total shares %s", err.Error())
	}

	return &types.QueryTotalSharesResponse{TotalShares: shares}, nil
}
