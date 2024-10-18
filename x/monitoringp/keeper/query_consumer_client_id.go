package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/monitoringp/types"
)

func (q queryServer) GetConsumerClientID(ctx context.Context, req *types.QueryGetConsumerClientIDRequest) (*types.QueryGetConsumerClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.ConsumerClientID.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetConsumerClientIDResponse{ConsumerClientID: val}, nil
}
