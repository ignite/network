package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
	monitoringctypes "github.com/ignite/network/x/monitoringc/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func TestMsgRevertLaunch(t *testing.T) {
	ctx, tk, ts := testkeeper.NewTestSetup(t)
	launchParams, err := tk.LaunchKeeper.Params.Get(ctx)
	require.NoError(t, err)

	type inputState struct {
		noChain            bool
		noCoordinator      bool
		noVerifiedClientID bool
		chain              types.Chain
		coordinator        profiletypes.Coordinator
		verifiedClientID   string
		blockTime          time.Time
		blockHeight        int64
	}
	sampleTime := sample.Time(r)
	sampleAddr := sample.Address(r)

	for _, tt := range []struct {
		name       string
		inputState inputState
		msg        types.MsgRevertLaunch
		err        error
	}{
		{
			name: "should allow reverting launch if revert delay reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:        0,
					CoordinatorId:   0,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 0,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(launchParams.RevertDelay),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    0,
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should allow reverting launch if revert delay reached and chain has no monitoring connection but verified client ID",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:        0,
					CoordinatorId:   0,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 0,
					Address:       sampleAddr,
					Active:        true,
				},
				verifiedClientID: "test-client-id-1",
				blockTime:        sampleTime.Add(launchParams.RevertDelay),
				blockHeight:      100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    0,
				Coordinator: sampleAddr,
			},
		},
		{
			name: "should prevent reverting launch if revert delay not reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:        1,
					CoordinatorId:   1,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 1,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(launchParams.RevertDelay - time.Second),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    1,
				Coordinator: sampleAddr,
			},
			err: types.ErrRevertDelayNotReached,
		},
		{
			name: "should prevent reverting launch if revert delay not reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:        2,
					CoordinatorId:   2,
					LaunchTriggered: false,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 2,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(launchParams.RevertDelay),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    2,
				Coordinator: sampleAddr,
			},
			err: types.ErrNotTriggeredLaunch,
		},
		{
			name: "should allow reverting launch if revert delay reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:        3,
					CoordinatorId:   3,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				noCoordinator:      true,
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(launchParams.RevertDelay),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    3,
				Coordinator: sample.Address(r),
			},
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should allow reverting launch if revert delay reached",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:        4,
					CoordinatorId:   1000,
					LaunchTriggered: true,
					LaunchTime:      sampleTime,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 4,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(launchParams.RevertDelay),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    4,
				Coordinator: sampleAddr,
			},
			err: profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "should prevent reverting launch with non existent chain id",
			inputState: inputState{
				noChain: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 5,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(launchParams.RevertDelay),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    1000,
				Coordinator: sampleAddr,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent reverting launch if monitoring module is connected",
			inputState: inputState{
				chain: types.Chain{
					LaunchId:            6,
					CoordinatorId:       6,
					LaunchTriggered:     true,
					LaunchTime:          sampleTime,
					MonitoringConnected: true,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorId: 6,
					Address:       sampleAddr,
					Active:        true,
				},
				noVerifiedClientID: true,
				blockTime:          sampleTime.Add(launchParams.RevertDelay),
				blockHeight:        100,
			},
			msg: types.MsgRevertLaunch{
				LaunchId:    6,
				Coordinator: sampleAddr,
			},
			err: types.ErrChainMonitoringConnected,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noChain {
				err = tk.LaunchKeeper.Chain.Set(ctx, tt.inputState.chain.LaunchId, tt.inputState.chain)
				require.NoError(t, err)
			}
			if !tt.inputState.noCoordinator {
				err = tk.ProfileKeeper.Coordinator.Set(ctx, tt.inputState.coordinator.CoordinatorId, tt.inputState.coordinator)
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
			if !tt.inputState.noVerifiedClientID {
				err = tk.MonitoringConsumerKeeper.VerifiedClientID.Set(ctx, tt.inputState.chain.LaunchId, monitoringctypes.VerifiedClientID{
					LaunchId:     tt.inputState.chain.LaunchId,
					ClientIdList: []string{tt.inputState.verifiedClientID},
				})
				require.NoError(t, err)
				err = tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Set(ctx, tt.inputState.verifiedClientID, monitoringctypes.LaunchIDFromVerifiedClientID{
					LaunchId: tt.inputState.chain.LaunchId,
					ClientId: tt.inputState.verifiedClientID,
				})
				require.NoError(t, err)
			}

			// Send the message
			_, err := ts.LaunchSrv.RevertLaunch(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// Check value of chain
			chain, err := tk.LaunchKeeper.GetChain(ctx, tt.msg.LaunchId)
			require.NoError(t, err)
			require.False(t, chain.LaunchTriggered)

			// check that monitoringc client ids are removed
			_, err = tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, tt.msg.LaunchId)
			require.Error(t, err)
			_, err = tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Get(ctx, tt.inputState.verifiedClientID)
			require.Error(t, err)
		})
	}
}
