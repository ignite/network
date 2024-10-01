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

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGenesisValidator(keeper keeper.Keeper, ctx context.Context, n int) []types.GenesisValidator {
	items := make([]types.GenesisValidator, n)
	for i := range items {
		address := sample.AccAddress(r)
		items[i].LaunchID = uint64(i)
		items[i].Address = address.String()

		_ = keeper.GenesisValidator.Set(ctx, collections.Join(items[i].LaunchID, address), items[i])
	}
	return items
}

func TestGenesisValidatorQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNGenesisValidator(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetGenesisValidatorRequest
		response *types.QueryGetGenesisValidatorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGenesisValidatorRequest{
				LaunchID: msgs[0].LaunchID,
			},
			response: &types.QueryGetGenesisValidatorResponse{GenesisValidator: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGenesisValidatorRequest{
				LaunchID: msgs[1].LaunchID,
			},
			response: &types.QueryGetGenesisValidatorResponse{GenesisValidator: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGenesisValidatorRequest{
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
			response, err := qs.GetGenesisValidator(ctx, tc.request)
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

func TestGenesisValidatorQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.LaunchKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNGenesisValidator(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGenesisValidatorRequest {
		return &types.QueryAllGenesisValidatorRequest{
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
			resp, err := qs.ListGenesisValidator(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GenesisValidator),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListGenesisValidator(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.GenesisValidator),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListGenesisValidator(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.GenesisValidator),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListGenesisValidator(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
