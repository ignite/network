package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/collections"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

func createNRequest(keeper *keeper.Keeper, ctx context.Context, n int) []types.Request {
	items := make([]types.Request, n)
	launchID := uint64(0)
	for i := range items {
		iu := uint64(i)
		items[i].LaunchId = launchID
		items[i].RequestId = iu
		_ = keeper.Request.Set(ctx, collections.Join(launchID, iu), items[i])
		_ = keeper.RequestSeq.Set(ctx, launchID, iu)
	}
	return items
}

func TestRequestQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNRequest(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetRequestRequest
		response *types.QueryGetRequestResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetRequestRequest{LaunchId: msgs[0].LaunchId, RequestId: msgs[0].RequestId},
			response: &types.QueryGetRequestResponse{Request: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetRequestRequest{LaunchId: msgs[1].LaunchId, RequestId: msgs[1].RequestId},
			response: &types.QueryGetRequestResponse{Request: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetRequestRequest{LaunchId: uint64(len(msgs)), RequestId: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetRequest(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRequestQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNRequest(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRequestRequest {
		return &types.QueryAllRequestRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListRequest(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Request), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Request),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListRequest(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Request), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Request),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListRequest(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Request),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListRequest(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
