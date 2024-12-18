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
	"github.com/ignite/network/x/participation/keeper"
	"github.com/ignite/network/x/participation/types"
)

func createNAuctionUsedAllocations(keeper keeper.Keeper, ctx context.Context, n int) []types.AuctionUsedAllocations {
	items := make([]types.AuctionUsedAllocations, n)
	for i := range items {
		auctionID := uint64(i)
		address := sample.AccAddress(r)
		items[i].Address = address.String()
		items[i].AuctionId = auctionID
		items[i].NumAllocations = sample.Int(r)

		_ = keeper.AuctionUsedAllocations.Set(ctx, collections.Join(address, auctionID), items[i])
	}
	return items
}

func createNAuctionUsedAllocationsWithSameAddress(keeper keeper.Keeper, ctx context.Context, n int) []types.AuctionUsedAllocations {
	items := make([]types.AuctionUsedAllocations, n)
	address := sample.AccAddress(r)
	for i := range items {
		auctionID := uint64(i)
		items[i].Address = address.String()
		items[i].NumAllocations = sample.Int(r)

		_ = keeper.AuctionUsedAllocations.Set(ctx, collections.Join(address, auctionID), items[i])
	}
	return items
}

func TestAuctionUsedAllocationsQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ParticipationKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNAuctionUsedAllocations(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAuctionUsedAllocationsRequest
		response *types.QueryGetAuctionUsedAllocationsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAuctionUsedAllocationsRequest{
				AuctionId: msgs[0].AuctionId,
				Address:   msgs[0].Address,
			},
			response: &types.QueryGetAuctionUsedAllocationsResponse{AuctionUsedAllocations: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAuctionUsedAllocationsRequest{
				AuctionId: msgs[1].AuctionId,
				Address:   msgs[1].Address,
			},
			response: &types.QueryGetAuctionUsedAllocationsResponse{AuctionUsedAllocations: msgs[1]},
		},
		{
			desc: "NotFound",
			request: &types.QueryGetAuctionUsedAllocationsRequest{
				AuctionId: 100,
				Address:   sample.Address(r),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidAddress",
			request: &types.QueryGetAuctionUsedAllocationsRequest{
				AuctionId: 100,
				Address:   strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "invalid address"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetAuctionUsedAllocations(ctx, tc.request)
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

func TestAuctionUsedAllocationsQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.ParticipationKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNAuctionUsedAllocationsWithSameAddress(k, ctx, 5)
	address := msgs[0].Address

	request := func(addr string, next []byte, offset, limit uint64, total bool) *types.QueryAllAuctionUsedAllocationsRequest {
		return &types.QueryAllAuctionUsedAllocationsRequest{
			Address: addr,
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("should return empty set", func(t *testing.T) {
		resp, err := qs.ListAuctionUsedAllocations(ctx, request("invalid", nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, 0, int(resp.Pagination.Total))
		require.Nil(t, resp.AuctionUsedAllocations)
	})
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListAuctionUsedAllocations(ctx, request(address, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AuctionUsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AuctionUsedAllocations),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListAuctionUsedAllocations(ctx, request(address, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AuctionUsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AuctionUsedAllocations),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListAuctionUsedAllocations(ctx, request(address, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.AuctionUsedAllocations),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListAuctionUsedAllocations(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
