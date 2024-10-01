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
	"github.com/ignite/network/x/reward/keeper"
	"github.com/ignite/network/x/reward/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRewardPool(keeper keeper.Keeper, ctx context.Context, n int) []types.RewardPool {
	items := make([]types.RewardPool, n)
	for i := range items {
		items[i].LaunchID = uint64(i)

		_ = keeper.RewardPool.Set(ctx, items[i].LaunchID, items[i])
	}
	return items
}

func TestRewardPoolQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.RewardKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNRewardPool(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetRewardPoolRequest
		response *types.QueryGetRewardPoolResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRewardPoolRequest{
				LaunchID: msgs[0].LaunchID,
			},
			response: &types.QueryGetRewardPoolResponse{RewardPool: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRewardPoolRequest{
				LaunchID: msgs[1].LaunchID,
			},
			response: &types.QueryGetRewardPoolResponse{RewardPool: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRewardPoolRequest{
				LaunchID: 100000,
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
			response, err := qs.GetRewardPool(ctx, tc.request)
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

func TestRewardPoolQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.RewardKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNRewardPool(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRewardPoolRequest {
		return &types.QueryAllRewardPoolRequest{
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
			resp, err := qs.ListRewardPool(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RewardPool), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RewardPool),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListRewardPool(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RewardPool), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RewardPool),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListRewardPool(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RewardPool),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListRewardPool(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
