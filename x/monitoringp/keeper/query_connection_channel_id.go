package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/monitoringp/types"
)

func (q queryServer) GetConnectionChannelID(ctx context.Context, req *types.QueryGetConnectionChannelIDRequest) (*types.QueryGetConnectionChannelIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.ConnectionChannelID.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetConnectionChannelIDResponse{ConnectionChannelID: val}, nil
}
