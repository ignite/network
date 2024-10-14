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

func (q queryServer) ListVestingAccount(ctx context.Context, req *types.QueryAllVestingAccountRequest) (*types.QueryAllVestingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	vestingAccounts, pageRes, err := query.CollectionFilteredPaginate(
		ctx,
		q.k.VestingAccount,
		req.Pagination,
		func(_ collections.Pair[uint64, sdk.AccAddress], value types.VestingAccount) (bool, error) {
			return req.LaunchID == value.LaunchID, nil
		},
		func(_ collections.Pair[uint64, sdk.AccAddress], value types.VestingAccount) (types.VestingAccount, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVestingAccountResponse{VestingAccount: vestingAccounts, Pagination: pageRes}, nil
}

func (q queryServer) GetVestingAccount(ctx context.Context, req *types.QueryGetVestingAccountRequest) (*types.QueryGetVestingAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	address, err := q.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	val, err := q.k.VestingAccount.Get(ctx, collections.Join(req.LaunchID, sdk.AccAddress(address)))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetVestingAccountResponse{VestingAccount: val}, nil
}
