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

func (q queryServer) ListChain(ctx context.Context, req *types.QueryAllChainRequest) (*types.QueryAllChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	chains, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Chain,
		req.Pagination,
		func(_ uint64, value types.Chain) (types.Chain, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllChainResponse{Chain: chains, Pagination: pageRes}, nil
}

func (q queryServer) GetChain(ctx context.Context, req *types.QueryGetChainRequest) (*types.QueryGetChainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	chain, err := q.k.Chain.Get(ctx, req.LaunchId)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetChainResponse{Chain: chain}, nil
}
