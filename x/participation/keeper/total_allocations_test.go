package keeper_test

import (
	"strconv"
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/participation/types"
)

func TestTotalAllocationsGet(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	invalidAddress := strconv.Itoa(1)
	params := types.DefaultParams()
	params.AllocationPrice = types.AllocationPrice{Bonded: sdkmath.NewInt(100)}

	err := tk.ParticipationKeeper.Params.Set(ctx, params)
	require.NoError(t, err)
	validAddress := sample.Address(r)
	addressNegativeDelegations := sample.Address(r)

	_, _, err = tk.DelegateN(ctx, r, validAddress, 100, 10)
	require.NoError(t, err)
	_, _, err = tk.DelegateN(ctx, r, addressNegativeDelegations, -100, 10)
	require.NoError(t, err)

	for _, tc := range []struct {
		name       string
		address    string
		allocation sdkmath.Int
		wantError  bool
	}{
		{
			name:       "should allow valid address",
			address:    validAddress,
			allocation: sdkmath.NewInt(10), // 100 * 10 / 100 = 10
		},
		{
			name:      "should prevent invalid address",
			address:   invalidAddress,
			wantError: true,
		},
		{
			name:      "should prevent negative delegations",
			address:   addressNegativeDelegations,
			wantError: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			alloc, err := tk.ParticipationKeeper.GetTotalAllocations(ctx, tc.address)
			if tc.wantError {
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.allocation, alloc)
			}
		})
	}
}
