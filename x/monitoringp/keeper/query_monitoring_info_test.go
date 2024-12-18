package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/monitoringp/keeper"
	"github.com/ignite/network/x/monitoringp/types"
)

func TestMonitoringInfoQuery(t *testing.T) {
	k, ctx, _ := keepertest.MonitoringpKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	item := types.MonitoringInfo{}
	err := k.MonitoringInfo.Set(ctx, item)
	require.NoError(t, err)

	tests := []struct {
		desc     string
		request  *types.QueryGetMonitoringInfoRequest
		response *types.QueryGetMonitoringInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMonitoringInfoRequest{},
			response: &types.QueryGetMonitoringInfoResponse{MonitoringInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetMonitoringInfo(ctx, tc.request)
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
