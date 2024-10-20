package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/monitoringc/keeper"
	"github.com/ignite/network/x/monitoringc/types"
)

func createNMonitoringHistory(keeper *keeper.Keeper, ctx context.Context, n int) []types.MonitoringHistory {
	items := make([]types.MonitoringHistory, n)
	for i := range items {
		items[i].LaunchId = uint64(i)

		_ = keeper.MonitoringHistory.Set(ctx, items[i].LaunchId, items[i])
	}
	return items
}

func TestMonitoringHistoryQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.MonitoringcKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNMonitoringHistory(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMonitoringHistoryRequest
		response *types.QueryGetMonitoringHistoryResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchId: msgs[0].LaunchId,
			},
			response: &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchId: msgs[1].LaunchId,
			},
			response: &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetMonitoringHistoryRequest{
				LaunchId: 100000,
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
			response, err := qs.GetMonitoringHistory(ctx, tc.request)
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
