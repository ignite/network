package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) ListMainnetAccount(ctx context.Context, req *types.QueryAllMainnetAccountRequest) (*types.QueryAllMainnetAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	mainnetAccounts, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.MainnetAccount,
		req.Pagination,
		func(_ collections.Pair[uint64, sdk.AccAddress], value types.MainnetAccount) (types.MainnetAccount, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMainnetAccountResponse{MainnetAccount: mainnetAccounts, Pagination: pageRes}, nil
}

func (q queryServer) GetMainnetAccount(ctx context.Context, req *types.QueryGetMainnetAccountRequest) (*types.QueryGetMainnetAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	address, err := q.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	val, err := q.k.MainnetAccount.Get(ctx, collections.Join(req.ProjectId, sdk.AccAddress(address)))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetMainnetAccountResponse{MainnetAccount: val}, nil
}
