package keeper_test

import (
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/participation/keeper"
	"github.com/ignite/network/x/participation/types"
)

func TestShowAvailableAllocationsQuery(t *testing.T) {
	ctx, tk, _ := keepertest.NewTestSetup(t)
	qs := keeper.NewQueryServerImpl(tk.ParticipationKeeper)

	allocationPrice := types.AllocationPrice{Bonded: sdkmath.NewInt(100)}
	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	err := tk.ParticipationKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	addr := sample.Address(r)
	dels, _, err := tk.DelegateN(ctx, r, addr, 100, 10)
	require.NoError(t, err)

	for _, tc := range []struct {
		name     string
		request  *types.QueryAvailableAllocationsRequest
		response *types.QueryAvailableAllocationsResponse
		err      error
	}{
		{
			name: "should allow valid case",
			request: &types.QueryAvailableAllocationsRequest{
				Address: dels[0].DelegatorAddress,
			},
			response: &types.QueryAvailableAllocationsResponse{AvailableAllocations: sdkmath.NewInt(10)},
		},

		{
			name: "should prevent invalid address",
			request: &types.QueryAvailableAllocationsRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "decoding bech32 failed: invalid bech32 string length 6: invalid participant address"),
		},
		{
			name: "should return invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response, err := qs.AvailableAllocations(ctx, tc.request)
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
