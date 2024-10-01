package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/testutil/sample"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNMainnetAccount(keeper keeper.Keeper, ctx context.Context, n int) []types.MainnetAccount {
	items := make([]types.MainnetAccount, n)
	for i := range items {
		address := sample.AccAddress(r)
		items[i].ProjectID = uint64(i)
		items[i].Address = address.String()

		_ = keeper.MainnetAccount.Set(ctx, collections.Join(items[i].ProjectID, address), items[i])
	}
	return items
}

func TestMainnetAccountQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ProjectKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNMainnetAccount(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMainnetAccountRequest
		response *types.QueryGetMainnetAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMainnetAccountRequest{
				ProjectID: msgs[0].ProjectID,
			},
			response: &types.QueryGetMainnetAccountResponse{MainnetAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMainnetAccountRequest{
				ProjectID: msgs[1].ProjectID,
			},
			response: &types.QueryGetMainnetAccountResponse{MainnetAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMainnetAccountRequest{
				ProjectID: 100000,
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
			response, err := qs.GetMainnetAccount(ctx, tc.request)
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

func TestMainnetAccountQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.ProjectKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNMainnetAccount(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMainnetAccountRequest {
		return &types.QueryAllMainnetAccountRequest{
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
			resp, err := qs.ListMainnetAccount(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MainnetAccount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListMainnetAccount(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MainnetAccount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListMainnetAccount(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MainnetAccount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListMainnetAccount(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
