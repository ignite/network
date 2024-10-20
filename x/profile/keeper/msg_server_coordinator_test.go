package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/types"
)

func TestMsgCreateCoordinator(t *testing.T) {
	var (
		msg1        = sample.MsgCreateCoordinator(sample.Address(r))
		msg2        = sample.MsgCreateCoordinator(sample.Address(r))
		ctx, tk, ts = keepertest.NewTestSetup(t)
	)
	tests := []struct {
		name   string
		msg    types.MsgCreateCoordinator
		wantId uint64
		err    error
	}{
		{
			name:   "should allow creating a coordinator",
			msg:    msg1,
			wantId: 0,
		},
		{
			name:   "should allow creating a second coordinator",
			msg:    msg2,
			wantId: 1,
		},
		{
			name: "should prevent creating with an existing coordinator address",
			msg:  msg2,
			err:  types.ErrCoordinatorAlreadyExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.ProfileSrv.CreateCoordinator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			address, err := tk.ProfileKeeper.AddressCodec().StringToBytes(tt.msg.Address)
			require.NoError(t, err)
			coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
			require.NoError(t, err)
			require.EqualValues(t, tt.wantId, coordByAddr.CoordinatorId)
			require.EqualValues(t, tt.wantId, got.CoordinatorId)

			coord, err := tk.ProfileKeeper.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			require.NoError(t, err, "coordinator id not found")
			require.EqualValues(t, tt.msg.Address, coord.Address)
			require.EqualValues(t, tt.msg.Identity, coord.Description.Identity)
			require.EqualValues(t, tt.msg.Website, coord.Description.Website)
			require.EqualValues(t, tt.msg.Details, coord.Description.Details)
			require.EqualValues(t, coordByAddr.CoordinatorId, coord.CoordinatorId)
			require.EqualValues(t, true, coord.Active)
		})
	}
}

func TestMsgDisableCoordinator(t *testing.T) {
	var (
		addr        = sample.Address(r)
		msgCoord    = sample.MsgCreateCoordinator(sample.Address(r))
		ctx, tk, ts = keepertest.NewTestSetup(t)
	)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCoord)
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgDisableCoordinator
		err  error
	}{
		{
			name: "should prevent disabling a non existing coordinator",
			msg:  types.MsgDisableCoordinator{Address: addr},
			err:  types.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should allow disabling an active coordinator",
			msg:  types.MsgDisableCoordinator{Address: msgCoord.Address},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.ProfileSrv.DisableCoordinator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			address, err := tk.ProfileKeeper.AddressCodec().StringToBytes(tt.msg.Address)
			require.NoError(t, err)
			_, err = tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
			require.ErrorIs(t, err, types.ErrCoordinatorAddressNotFound)

			coord, err := tk.ProfileKeeper.GetCoordinator(ctx, got.CoordinatorId)
			require.NoError(t, err)
			require.EqualValues(t, false, coord.Active)
		})
	}
}

func TestMsgUpdateCoordinatorAddress(t *testing.T) {
	var (
		addr         = sample.Address(r)
		nonExistAddr = sample.Address(r)
		coord1       = sample.MsgCreateCoordinator(sample.Address(r))
		coord2       = sample.MsgCreateCoordinator(sample.Address(r))
		ctx, tk, ts  = keepertest.NewTestSetup(t)
	)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &coord1)
	require.NoError(t, err)
	_, err = ts.ProfileSrv.CreateCoordinator(ctx, &coord2)
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorAddress
		err  error
	}{
		{
			name: "should prevent updating a non existing coordinator",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: nonExistAddr,
			},
			err: types.ErrCoordinatorAddressNotFound,
		}, {
			name: "should prevent updating with an address already associated to a coordinator",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord1.Address,
				NewAddress: coord2.Address,
			},
			err: types.ErrCoordinatorAlreadyExist,
		}, {
			name: "should allow updating coordinator address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord1.Address,
				NewAddress: addr,
			},
		}, {
			name: "should allow updating coordinator address a second time",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    coord2.Address,
				NewAddress: coord1.Address,
			},
		}, {
			name: "should prevent updating from previous coordinator address",
			msg: types.MsgUpdateCoordinatorAddress{
				Address:    addr,
				NewAddress: coord1.Address,
			},
			err: types.ErrCoordinatorAlreadyExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ts.ProfileSrv.UpdateCoordinatorAddress(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			address, err := tk.ProfileKeeper.AddressCodec().StringToBytes(tt.msg.Address)
			require.NoError(t, err)
			_, err = tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
			require.ErrorIs(t, err, types.ErrCoordinatorAddressNotFound, "old coordinator address was not removed")

			newAddress, err := tk.ProfileKeeper.AddressCodec().StringToBytes(tt.msg.NewAddress)
			require.NoError(t, err)
			coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, newAddress)
			require.NoError(t, err, "coordinator by address not found")
			require.EqualValues(t, tt.msg.NewAddress, coordByAddr.Address)

			coord, err := tk.ProfileKeeper.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			require.NoError(t, err, "coordinator id not found")
			require.EqualValues(t, tt.msg.NewAddress, coord.Address)
			require.EqualValues(t, coordByAddr.CoordinatorId, coord.CoordinatorId)
		})
	}
}

func TestMsgUpdateCoordinatorDescription(t *testing.T) {
	var (
		addr        = sample.Address(r)
		msgCoord    = sample.MsgCreateCoordinator(sample.Address(r))
		ctx, tk, ts = keepertest.NewTestSetup(t)
	)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCoord)
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorDescription
		err  error
	}{
		{
			name: "should prevent updating description of non existing coordinator",
			msg:  sample.MsgUpdateCoordinatorDescription(addr),
			err:  types.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should allow updating one value of coordinator description",
			msg: types.MsgUpdateCoordinatorDescription{
				Address:  msgCoord.Address,
				Identity: "update",
			},
		},
		{
			name: "should allow updating all values of coordinator description",
			msg:  sample.MsgUpdateCoordinatorDescription(msgCoord.Address),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, err := tk.ProfileKeeper.AddressCodec().StringToBytes(tt.msg.Address)
			require.NoError(t, err)

			var oldCoord types.Coordinator
			if tt.err == nil {
				coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
				require.NoError(t, err, "coordinator by address not found")
				oldCoord, err = tk.ProfileKeeper.GetCoordinator(ctx, coordByAddr.CoordinatorId)
				require.NoError(t, err, "coordinator not found")
			}

			_, err = ts.ProfileSrv.UpdateCoordinatorDescription(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			coordByAddr, err := tk.ProfileKeeper.GetCoordinatorByAddress(ctx, address)
			require.NoError(t, err, "coordinator by address not found")
			coord, err := tk.ProfileKeeper.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			require.NoError(t, err, "coordinator not found")
			require.EqualValues(t, tt.msg.Address, coord.Address)
			require.EqualValues(t, coordByAddr.CoordinatorId, coord.CoordinatorId)

			if len(tt.msg.Identity) > 0 {
				require.EqualValues(t, tt.msg.Identity, coord.Description.Identity)
			} else {
				require.EqualValues(t, oldCoord.Description.Identity, coord.Description.Identity)
			}

			if len(tt.msg.Website) > 0 {
				require.EqualValues(t, tt.msg.Website, coord.Description.Website)
			} else {
				require.EqualValues(t, oldCoord.Description.Website, coord.Description.Website)
			}

			if len(tt.msg.Details) > 0 {
				require.EqualValues(t, tt.msg.Details, coord.Description.Details)
			} else {
				require.EqualValues(t, oldCoord.Description.Details, coord.Description.Details)
			}
		})
	}
}
