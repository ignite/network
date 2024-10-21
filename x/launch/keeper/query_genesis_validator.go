package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/launch/types"
)

func (q queryServer) ListGenesisValidator(ctx context.Context, req *types.QueryAllGenesisValidatorRequest) (*types.QueryAllGenesisValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	genesisValidators, pageRes, err := query.CollectionFilteredPaginate(
		ctx,
		q.k.GenesisValidator,
		req.Pagination,
		func(_ collections.Pair[uint64, sdk.AccAddress], value types.GenesisValidator) (bool, error) {
			return req.LaunchId == value.LaunchId, nil
		},
		func(_ collections.Pair[uint64, sdk.AccAddress], value types.GenesisValidator) (types.GenesisValidator, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGenesisValidatorResponse{GenesisValidator: genesisValidators, Pagination: pageRes}, nil
}

func (q queryServer) GetGenesisValidator(ctx context.Context, req *types.QueryGetGenesisValidatorRequest) (*types.QueryGetGenesisValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	address, err := q.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	val, err := q.k.GenesisValidator.Get(ctx, collections.Join(req.LaunchId, sdk.AccAddress(address)))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetGenesisValidatorResponse{GenesisValidator: val}, nil
}
