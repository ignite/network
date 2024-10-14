package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/participation/types"
)

func (q queryServer) ListAuctionUsedAllocations(ctx context.Context, req *types.QueryAllAuctionUsedAllocationsRequest) (*types.QueryAllAuctionUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	auctionUsedAllocations, pageRes, err := query.CollectionFilteredPaginate(
		ctx,
		q.k.AuctionUsedAllocations,
		req.Pagination,
		func(key collections.Pair[sdk.AccAddress, uint64], value types.AuctionUsedAllocations) (bool, error) {
			return req.Address == value.Address, nil
		},
		func(_ collections.Pair[sdk.AccAddress, uint64], value types.AuctionUsedAllocations) (types.AuctionUsedAllocations, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAuctionUsedAllocationsResponse{AuctionUsedAllocations: auctionUsedAllocations, Pagination: pageRes}, nil
}

func (q queryServer) GetAuctionUsedAllocations(ctx context.Context, req *types.QueryGetAuctionUsedAllocationsRequest) (*types.QueryGetAuctionUsedAllocationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	address, err := q.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	val, err := q.k.AuctionUsedAllocations.Get(ctx, collections.Join(sdk.AccAddress(address), req.AuctionID))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetAuctionUsedAllocationsResponse{AuctionUsedAllocations: val}, nil
}
