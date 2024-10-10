package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/x/profile/keeper"
	"github.com/ignite/network/x/profile/types"
)

func TestCoordinatorMsgServerCreate(t *testing.T) {
	k, ctx, addressCodec := keepertest.ProfileKeeper(t)
	srv := keeper.NewMsgServerImpl(k)

	address, err := addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		resp, err := srv.CreateCoordinator(ctx, &types.MsgCreateCoordinator{Address: address})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.CoordinatorID))
	}
}

func TestCoordinatorMsgServerUpdate(t *testing.T) {
	k, ctx, addressCodec := keepertest.ProfileKeeper(t)
	srv := keeper.NewMsgServerImpl(k)

	address, err := addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateCoordinator(ctx, &types.MsgCreateCoordinator{Address: address})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdateCoordinatorDescription
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgUpdateCoordinatorDescription{Address: "invalid", Description: types.CoordinatorDescription{}},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgUpdateCoordinatorDescription{Address: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgUpdateCoordinatorDescription{Address: address, Description: types.CoordinatorDescription{
				Identity: "id",
				Website:  "wb",
				Details:  "dt",
			}},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgUpdateCoordinatorDescription{Address: address},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdateCoordinatorDescription(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCoordinatorMsgServerDelete(t *testing.T) {
	k, ctx, addressCodec := keepertest.ProfileKeeper(t)
	srv := keeper.NewMsgServerImpl(k)

	address, err := addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateCoordinator(ctx, &types.MsgCreateCoordinator{Address: address})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDisableCoordinator
		err     error
	}{
		{
			desc:    "invalid address",
			request: &types.MsgDisableCoordinator{Address: "invalid"},
			err:     sdkerrors.ErrInvalidAddress,
		},
		{
			desc:    "unauthorized",
			request: &types.MsgDisableCoordinator{Address: unauthorizedAddr},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "key not found",
			request: &types.MsgDisableCoordinator{Address: address},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "completed",
			request: &types.MsgDisableCoordinator{Address: address},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DisableCoordinator(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
