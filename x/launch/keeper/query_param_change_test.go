package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

func createNParamChange(keeper *keeper.Keeper, ctx context.Context, n int) []types.ParamChange {
	items := make([]types.ParamChange, n)
	launchID := uint64(0)
	for i := range items {
		items[i].LaunchID = launchID
		items[i].Module = fmt.Sprintf("module_%d", i)
		items[i].Param = fmt.Sprintf("param_%d", i)

		_ = keeper.ParamChange.Set(ctx, collections.Join(items[i].LaunchID, types.ParamChangeSubKey(items[i].Module, items[i].Param)), items[i])
	}
	return items
}

func TestParamChangeQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNParamChange(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllParamChangeRequest {
		return &types.QueryAllParamChangeRequest{
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
			resp, err := qs.ListParamChange(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ParamChange), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ParamChange),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListParamChange(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ParamChange), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ParamChange),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListParamChange(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ParamChange),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListParamChange(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
