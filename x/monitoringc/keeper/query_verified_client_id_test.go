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

func createNVerifiedClientID(ctx context.Context, keeper *keeper.Keeper, n int) []types.VerifiedClientID {
	items := make([]types.VerifiedClientID, n)
	for i := range items {
		items[i].LaunchID = uint64(i)

		_ = keeper.VerifiedClientID.Set(ctx, items[i].LaunchID, items[i])
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
				LaunchID: msgs[0].LaunchID,
			},
			response: &types.QueryGetVerifiedClientIDResponse{VerifiedClientID: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVerifiedClientIDRequest{
				LaunchID: msgs[1].LaunchID,
			},
			response: &types.QueryGetVerifiedClientIDResponse{VerifiedClientID: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVerifiedClientIDRequest{
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
