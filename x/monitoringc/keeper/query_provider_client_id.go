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

func (q queryServer) ListProviderClientID(ctx context.Context, req *types.QueryAllProviderClientIDRequest) (*types.QueryAllProviderClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	providerClientIDs, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.ProviderClientID,
		req.Pagination,
		func(_ uint64, value types.ProviderClientID) (types.ProviderClientID, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllProviderClientIDResponse{ProviderClientID: providerClientIDs, Pagination: pageRes}, nil
}

func (q queryServer) GetProviderClientID(ctx context.Context, req *types.QueryGetProviderClientIDRequest) (*types.QueryGetProviderClientIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.ProviderClientID.Get(ctx, req.LaunchID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetProviderClientIDResponse{ProviderClientID: val}, nil
}
