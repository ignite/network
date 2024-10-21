package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/participation/keeper"
	"github.com/ignite/network/x/participation/types"
)

func createNUsedAllocations(keeper keeper.Keeper, ctx context.Context, n int) []types.UsedAllocations {
	items := make([]types.UsedAllocations, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)
		items[i].NumAllocations = sample.Int(r)

		_ = keeper.UsedAllocations.Set(ctx, items[i].Address, items[i])
	}
	return items
}

func TestUsedAllocationsQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ParticipationKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNUsedAllocations(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetUsedAllocationsRequest
		response *types.QueryGetUsedAllocationsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUsedAllocationsRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetUsedAllocationsResponse{UsedAllocations: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUsedAllocationsRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetUsedAllocationsResponse{UsedAllocations: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUsedAllocationsRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetUsedAllocations(ctx, tc.request)
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

func TestUsedAllocationsQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.ParticipationKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNUsedAllocations(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUsedAllocationsRequest {
		return &types.QueryAllUsedAllocationsRequest{
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
			resp, err := qs.ListUsedAllocations(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UsedAllocations),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListUsedAllocations(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.UsedAllocations),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListUsedAllocations(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.UsedAllocations),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListUsedAllocations(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
