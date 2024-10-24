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
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func createNProject(keeper keeper.Keeper, ctx context.Context, n int) []types.Project {
	items := make([]types.Project, n)
	for i := range items {
		iu := uint64(i)
		items[i].ProjectId = iu
		_ = keeper.Project.Set(ctx, iu, items[i])
		_ = keeper.ProjectSeq.Set(ctx, iu)
	}
	return items
}

func TestProjectQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ProjectKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNProject(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetProjectRequest
		response *types.QueryGetProjectResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetProjectRequest{ProjectId: msgs[0].ProjectId},
			response: &types.QueryGetProjectResponse{Project: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetProjectRequest{ProjectId: msgs[1].ProjectId},
			response: &types.QueryGetProjectResponse{Project: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetProjectRequest{ProjectId: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetProject(ctx, tc.request)
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

func TestProjectQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.ProjectKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNProject(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProjectRequest {
		return &types.QueryAllProjectRequest{
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
			resp, err := qs.ListProject(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Project),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListProject(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Project), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Project),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListProject(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Project),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListProject(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
