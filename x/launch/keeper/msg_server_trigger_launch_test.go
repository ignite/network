package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func TestMsgTriggerLaunch(t *testing.T) {
	ctx, tk, ts := testkeeper.NewTestSetup(t)

	type inputState struct {
		noChain       bool
		noCoordinator bool
		chain         types.Chain
		coordinator   profiletypes.Coordinator
		blockTime     time.Time
		blockHeight   int64
	}
	sampleTime := sample.Time(r)
	sampleAddr := sample.Address(r)

	for _, tt := range []struct {
		name       string
		inputState inputState
		msg        types.MsgTriggerLaunch
		err        error
	}{
		{
			name: "should allow triggering a chain launch",
			inputState: inputState{
				chain: sample.Chain(r, 0, 0),
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 0,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    0,
				LaunchTime:  types.DefaultMinLaunchTime,
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should allow triggering a chain launch  with maximum launch time",
			inputState: inputState{
				chain: sample.Chain(r, 10, 10),
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 10,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 5000,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    10,
				LaunchTime:  types.DefaultMaxLaunchTime,
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should prevent triggering a chain launch from a non existing chain",
			inputState: inputState{
				noChain: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 1,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    1000,
				LaunchTime:  types.DefaultMinLaunchTime,
				Coordinator: sampleAddr,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent triggering a chain launch from a non existent coordinator",
			inputState: inputState{
				chain:         sample.Chain(r, 2, 2),
				noCoordinator: true,
				blockTime:     sampleTime,
				blockHeight:   100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    2,
				LaunchTime:  types.DefaultMinLaunchTime,
				Coordinator: sample.Address(r),
			},
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should prevent triggering a chain launch from an invalid coordinator",
			inputState: inputState{
				chain: sample.Chain(r, 3, 1000),
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 3,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    3,
				LaunchTime:  types.DefaultMinLaunchTime,
				Coordinator: sampleAddr,
			},
			err: profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "should prevent triggering a chain launch with chain launch already triggered",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:        5,
					CoordinatorId:   5,
					LaunchTriggered: true,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 5,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    5,
				LaunchTime:  types.DefaultMinLaunchTime,
				Coordinator: sampleAddr,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "should prevent triggering a chain launch with launch time too low",
			inputState: inputState{
				chain: sample.Chain(r, 6, 6),
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 6,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    6,
				LaunchTime:  types.DefaultMinLaunchTime - time.Second,
				Coordinator: sampleAddr,
			},
			err: types.ErrLaunchTimeTooLow,
		},
		{
			name: "should prevent triggering a chain launch with launch time too high",
			inputState: inputState{
				chain: sample.Chain(r, 7, 7),
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 7,
					Address:       sampleAddr,
					Active:        true,
				},
				blockTime:   sampleTime,
				blockHeight: 100,
			},
			msg: types.MsgTriggerLaunch{
				LaunchId:    7,
				LaunchTime:  types.DefaultMaxLaunchTime + time.Second,
				Coordinator: sampleAddr,
			},
			err: types.ErrLaunchTimeTooHigh,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noChain {
				err := tk.LaunchKeeper.Chain.Set(ctx, tt.inputState.chain.LaunchId, tt.inputState.chain)
				require.NoError(t, err)
			}
			if !tt.inputState.noCoordinator {
				err := tk.ProfileKeeper.Coordinator.Set(ctx, tt.inputState.coordinator.CoordinatorId, tt.inputState.coordinator)
				require.NoError(t, err)
				addr, err := tk.ProfileKeeper.AddressCodec().StringToBytes(tt.inputState.coordinator.Address)
				require.NoError(t, err)
				err = tk.ProfileKeeper.CoordinatorByAddress.Set(ctx, addr, profiletypes.CoordinatorByAddress{
					Address:       tt.inputState.coordinator.Address,
					CoordinatorId: tt.inputState.coordinator.CoordinatorId,
				})
				require.NoError(t, err)
			}
			if !tt.inputState.blockTime.IsZero() {
				ctx = ctx.WithBlockTime(tt.inputState.blockTime)
			}
			if tt.inputState.blockHeight > 0 {
				ctx = ctx.WithBlockHeight(tt.inputState.blockHeight)
			}

			// Send the message
			_, err := ts.LaunchSrv.TriggerLaunch(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// Check values
			chain, err := tk.LaunchKeeper.GetChain(ctx, tt.msg.LaunchId)
			require.NoError(t, err)
			require.True(t, chain.LaunchTriggered)
			require.EqualValues(t, sampleTime.Add(tt.msg.LaunchTime), chain.LaunchTime)
			require.EqualValues(t, tt.inputState.blockHeight, chain.ConsumerRevisionHeight)
		})
	}
}
