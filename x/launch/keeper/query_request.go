package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/launch/types"
)

func (q queryServer) ListRequest(ctx context.Context, req *types.QueryAllRequestRequest) (*types.QueryAllRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	requests, pageRes, err := query.CollectionFilteredPaginate(
		ctx,
		q.k.Request,
		req.Pagination,
		func(_ collections.Pair[uint64, uint64], value types.Request) (bool, error) {
			return value.LaunchID == req.LaunchID, nil
		},
		func(_ collections.Pair[uint64, uint64], value types.Request) (types.Request, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRequestResponse{Request: requests, Pagination: pageRes}, nil
}

func (q queryServer) GetRequest(ctx context.Context, req *types.QueryGetRequestRequest) (*types.QueryGetRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	request, err := q.k.Request.Get(ctx, collections.Join(req.LaunchID, req.RequestID))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetRequestResponse{Request: request}, nil
}
