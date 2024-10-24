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
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/keeper"
	"github.com/ignite/network/x/profile/types"
)

func createNCoordinator(keeper keeper.Keeper, ctx context.Context, n int) []types.Coordinator {
	items := make([]types.Coordinator, n)
	for i := range items {
		iu := uint64(i)
		items[i].CoordinatorId = iu
		_ = keeper.Coordinator.Set(ctx, iu, items[i])
		_ = keeper.CoordinatorSeq.Set(ctx, iu)
	}
	return items
}

func TestCoordinatorQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ProfileKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNCoordinator(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetCoordinatorRequest
		response *types.QueryGetCoordinatorResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCoordinatorRequest{CoordinatorId: msgs[0].CoordinatorId},
			response: &types.QueryGetCoordinatorResponse{Coordinator: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCoordinatorRequest{CoordinatorId: msgs[1].CoordinatorId},
			response: &types.QueryGetCoordinatorResponse{Coordinator: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCoordinatorRequest{CoordinatorId: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetCoordinator(ctx, tc.request)
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

func TestCoordinatorQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.ProfileKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNCoordinator(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCoordinatorRequest {
		return &types.QueryAllCoordinatorRequest{
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
			resp, err := qs.ListCoordinator(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Coordinator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Coordinator),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListCoordinator(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Coordinator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Coordinator),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListCoordinator(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Coordinator),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListCoordinator(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func createNCoordinatorByAddress(keeper keeper.Keeper, ctx context.Context, n int) []types.Coordinator {
	coordByAddress := make([]types.CoordinatorByAddress, n)
	coordinators := make([]types.Coordinator, n)
	for i := range coordByAddress {
		address := sample.AccAddress(r)
		iu := uint64(i)
		coordinators[i].CoordinatorId = iu
		coordinators[i].Address = address.String()
		_ = keeper.Coordinator.Set(ctx, iu, coordinators[i])
		_ = keeper.CoordinatorSeq.Set(ctx, iu)

		coordByAddress[i].CoordinatorId = iu
		coordByAddress[i].Address = address.String()
		_ = keeper.CoordinatorByAddress.Set(ctx, address, coordByAddress[i])

	}
	return coordinators
}

func TestCoordinatorByAddressQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ProfileKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNCoordinatorByAddress(k, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCoordinatorByAddressRequest
		response *types.QueryGetCoordinatorByAddressResponse
		err      error
	}{
		{
			desc:     "should allow querying first coordinator by address",
			request:  &types.QueryGetCoordinatorByAddressRequest{Address: msgs[0].Address},
			response: &types.QueryGetCoordinatorByAddressResponse{Coordinator: msgs[0]},
		},
		{
			desc:     "should allow querying second coordinator by address",
			request:  &types.QueryGetCoordinatorByAddressRequest{Address: msgs[1].Address},
			response: &types.QueryGetCoordinatorByAddressResponse{Coordinator: msgs[1]},
		},
		{
			desc:    "should prevent querying non existing coordinator by address",
			request: &types.QueryGetCoordinatorByAddressRequest{Address: sample.Address(r)},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "should prevent querying invalid coordinator by address",
			request: &types.QueryGetCoordinatorByAddressRequest{Address: "invalid"},
			err:     status.Error(codes.InvalidArgument, "invalid address"),
		},
		{
			desc: "should prevent querying with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetCoordinatorByAddress(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
