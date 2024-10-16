package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/x/monitoringp/keeper"
	"github.com/ignite/network/x/monitoringp/types"
)

func TestMsgUpdateParams(t *testing.T) {
	k, ctx, _ := keepertest.MonitoringpKeeper(t)
	ms := keeper.NewMsgServerImpl(k)

	params := types.DefaultParams()
	require.NoError(t, k.Params.Set(ctx, params))

	// default params
	testCases := []struct {
		name  string
		input *types.MsgUpdateParams
		err   error
	}{
		{
			name: "invalid authority",
			input: &types.MsgUpdateParams{
				Authority: "invalid",
				Params:    params,
			},
			err: types.ErrInvalidSigner,
		},
		{
			name: "send enabled param",
			input: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params: types.Params{
					LastBlockHeight:         types.DefaultLastBlockHeight,
					ConsumerChainID:         types.DefaultConsumerChainID,
					ConsumerUnbondingPeriod: networktypes.DefaultUnbondingPeriod,
					ConsumerRevisionHeight:  networktypes.DefaultRevisionHeight,
				},
			},
		},
		{
			name: "all good",
			input: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params:    params,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.UpdateParams(ctx, tc.input)

			if tc.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
