package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/monitoringc/types"
)

func (q queryServer) ListLaunchIDFromChannelID(ctx context.Context, req *types.QueryAllLaunchIDFromChannelIDRequest) (*types.QueryAllLaunchIDFromChannelIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	launchIDFromChannelIDs, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.LaunchIDFromChannelID,
		req.Pagination,
		func(_ string, value types.LaunchIDFromChannelID) (types.LaunchIDFromChannelID, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLaunchIDFromChannelIDResponse{LaunchIDFromChannelID: launchIDFromChannelIDs, Pagination: pageRes}, nil
}

func (q queryServer) GetLaunchIDFromChannelID(ctx context.Context, req *types.QueryGetLaunchIDFromChannelIDRequest) (*types.QueryGetLaunchIDFromChannelIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.LaunchIDFromChannelID.Get(ctx, req.ChannelID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetLaunchIDFromChannelIDResponse{LaunchIDFromChannelID: val}, nil
}
