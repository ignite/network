package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/profile/types"
)

func (q queryServer) ListCoordinator(ctx context.Context, req *types.QueryAllCoordinatorRequest) (*types.QueryAllCoordinatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	coordinators, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Coordinator,
		req.Pagination,
		func(_ uint64, value types.Coordinator) (types.Coordinator, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCoordinatorResponse{Coordinator: coordinators, Pagination: pageRes}, nil
}

func (q queryServer) GetCoordinator(ctx context.Context, req *types.QueryGetCoordinatorRequest) (*types.QueryGetCoordinatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	coordinator, err := q.k.Coordinator.Get(ctx, req.ID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetCoordinatorResponse{Coordinator: coordinator}, nil
}

func (q queryServer) GetCoordinatorByAddress(ctx context.Context, req *types.QueryGetCoordinatorByAddressRequest) (*types.QueryGetCoordinatorByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	address, err := q.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	coordinatorByAddress, err := q.k.CoordinatorByAddress.Get(ctx, address)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	coordinator, err := q.k.Coordinator.Get(ctx, coordinatorByAddress.CoordinatorID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, sdkerrors.ErrKeyNotFound
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetCoordinatorByAddressResponse{Coordinator: coordinator}, nil
}
