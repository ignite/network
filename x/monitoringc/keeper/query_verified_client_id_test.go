package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/x/monitoringc/keeper"
	"github.com/ignite/network/x/monitoringc/types"
)

func createNVerifiedClientID(ctx context.Context, keeper *keeper.Keeper, n int) []types.VerifiedClientID {
	items := make([]types.VerifiedClientID, n)
	for i := range items {
		items[i].LaunchId = uint64(i)
		items[i].ClientIdList = []string{strconv.Itoa(i)}

		_ = keeper.VerifiedClientID.Set(ctx, items[i].LaunchId, items[i])
	}
	return items
}

func TestVerifiedClientIDQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.MonitoringcKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNVerifiedClientID(ctx, k, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVerifiedClientIDRequest
		response *types.QueryGetVerifiedClientIDResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVerifiedClientIDRequest{
				LaunchId: msgs[0].LaunchId,
			},
			response: &types.QueryGetVerifiedClientIDResponse{VerifiedClientId: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVerifiedClientIDRequest{
				LaunchId: msgs[1].LaunchId,
			},
			response: &types.QueryGetVerifiedClientIDResponse{VerifiedClientId: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVerifiedClientIDRequest{
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
			response, err := qs.GetVerifiedClientID(ctx, tc.request)
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
