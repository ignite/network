package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	networktypes "github.com/ignite/network/pkg/types"
	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func TestTotalSharesQuery(t *testing.T) {
	var (
		k, ctx, _ = keepertest.ProjectKeeper(t)
		qs        = keeper.NewQueryServerImpl(k)
	)
	err := k.TotalShares.Set(ctx, networktypes.TotalShareNumber)
	require.NoError(t, err)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryTotalSharesRequest
		response *types.QueryTotalSharesResponse
		err      error
	}{
		{
			desc:     "should allow valid query",
			request:  &types.QueryTotalSharesRequest{},
			response: &types.QueryTotalSharesResponse{TotalShares: networktypes.TotalShareNumber},
		},
		{
			desc: "should return InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.TotalShares(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}
