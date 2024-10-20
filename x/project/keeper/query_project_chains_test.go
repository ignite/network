package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func createNProjectChains(keeper keeper.Keeper, ctx context.Context, n int) []types.ProjectChains {
	items := make([]types.ProjectChains, n)
	for i := range items {
		items[i].ProjectId = uint64(i)

		_ = keeper.ProjectChains.Set(ctx, items[i].ProjectId, items[i])
	}
	return items
}

func TestProjectChainsQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ProjectKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNProjectChains(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetProjectChainsRequest
		response *types.QueryGetProjectChainsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProjectChainsRequest{
				ProjectId: msgs[0].ProjectId,
			},
			response: &types.QueryGetProjectChainsResponse{ProjectChains: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProjectChainsRequest{
				ProjectId: msgs[1].ProjectId,
			},
			response: &types.QueryGetProjectChainsResponse{ProjectChains: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProjectChainsRequest{
				ProjectId: 100000,
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
			response, err := qs.GetProjectChains(ctx, tc.request)
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
