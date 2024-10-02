package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

func createNVestingAccount(keeper keeper.Keeper, ctx context.Context, n int) []types.VestingAccount {
	items := make([]types.VestingAccount, n)
	for i := range items {
		address := sample.AccAddress(r)
		items[i].LaunchID = uint64(i)
		items[i].Address = address.String()

		_ = keeper.VestingAccount.Set(ctx, collections.Join(items[i].LaunchID, address), items[i])
	}
	return items
}

func TestVestingAccountQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNVestingAccount(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVestingAccountRequest
		response *types.QueryGetVestingAccountResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVestingAccountRequest{
				LaunchID: msgs[0].LaunchID,
			},
			response: &types.QueryGetVestingAccountResponse{VestingAccount: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVestingAccountRequest{
				LaunchID: msgs[1].LaunchID,
			},
			response: &types.QueryGetVestingAccountResponse{VestingAccount: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVestingAccountRequest{
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
			response, err := qs.GetVestingAccount(ctx, tc.request)
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

func TestVestingAccountQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNVestingAccount(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVestingAccountRequest {
		return &types.QueryAllVestingAccountRequest{
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
			resp, err := qs.ListVestingAccount(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VestingAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VestingAccount),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListVestingAccount(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.VestingAccount), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.VestingAccount),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListVestingAccount(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.VestingAccount),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListVestingAccount(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
