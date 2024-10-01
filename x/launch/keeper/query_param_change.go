package keeper

import (
	"context"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/launch/types"
)

func (q queryServer) ListParamChange(ctx context.Context, req *types.QueryAllParamChangeRequest) (*types.QueryAllParamChangeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	paramChanges, pageRes, err := query.CollectionFilteredPaginate(
		ctx,
		q.k.ParamChange,
		req.Pagination,
		func(_ collections.Pair[uint64, string], value types.ParamChange) (bool, error) {
			return value.LaunchID == req.LaunchID, nil
		},
		func(_ collections.Pair[uint64, string], value types.ParamChange) (types.ParamChange, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllParamChangeResponse{ParamChange: paramChanges, Pagination: pageRes}, nil
}
