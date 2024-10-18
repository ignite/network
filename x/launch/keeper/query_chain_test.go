package keeper_test

import (
	"context"
	"testing"

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

func createNChain(keeper *keeper.Keeper, ctx context.Context, n int) []types.Chain {
	items := make([]types.Chain, n)
	for i := range items {
		iu := uint64(i)
		items[i].LaunchID = iu
		_ = keeper.Chain.Set(ctx, iu, items[i])
		_ = keeper.ChainSeq.Set(ctx, iu)
	}
	return items
}

func TestChainQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNChain(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetChainRequest
		response *types.QueryGetChainResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetChainRequest{LaunchID: msgs[0].LaunchID},
			response: &types.QueryGetChainResponse{Chain: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetChainRequest{LaunchID: msgs[1].LaunchID},
			response: &types.QueryGetChainResponse{Chain: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetChainRequest{LaunchID: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetChain(ctx, tc.request)
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

func TestChainQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNChain(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllChainRequest {
		return &types.QueryAllChainRequest{
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
			resp, err := qs.ListChain(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Chain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Chain),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListChain(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Chain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Chain),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListChain(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Chain),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListChain(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
