package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/reward/types"
)

func (q queryServer) ListRewardPool(ctx context.Context, req *types.QueryAllRewardPoolRequest) (*types.QueryAllRewardPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	rewardPools, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.RewardPool,
		req.Pagination,
		func(_ uint64, value types.RewardPool) (types.RewardPool, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRewardPoolResponse{RewardPool: rewardPools, Pagination: pageRes}, nil
}

func (q queryServer) GetRewardPool(ctx context.Context, req *types.QueryGetRewardPoolRequest) (*types.QueryGetRewardPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.RewardPool.Get(ctx, req.LaunchId)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetRewardPoolResponse{RewardPool: val}, nil
}
