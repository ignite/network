package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func TestMsgRequestAddAccount(t *testing.T) {
	ctx, tk, ts := testkeeper.NewTestSetup(t, testkeeper.WithLaunchHooksMock())
	coordAddr, addr := sample.Address(r), sample.Address(r)

	fee := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1000)))
	hasFeeAddr := sample.Address(r)
	tk.Mint(ctx, hasFeeAddr, fee)

	type inputState struct {
		noCoordinator bool
		noChain       bool
		noAccount     bool
		coordinator   profiletypes.Coordinator
		chain         types.Chain
		account       types.GenesisAccount
		fee           sdk.Coins
	}

	tests := []struct {
		name                     string
		inputState               inputState
		msg                      types.MsgSendRequest
		wantID                   uint64
		wantApprove              bool
		expectedCommunityPoolAmt sdk.Coin
		err                      error
	}{
		{
			name: "should prevent sending a request for a non existing chain",
			inputState: inputState{
				noAccount:     true,
				noChain:       true,
				noCoordinator: true,
				fee:           sdk.Coins(nil),
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 10000),
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent sending a request for a launch triggered chain",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 0,
					Address:       sample.Address(r),
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        0,
					LaunchTriggered: true,
					IsMainnet:       false,
					CoordinatorID:   0,
				},
				fee: sdk.Coins(nil),
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 0),
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "should prevent sending a request not valid for mainnet for a mainnet chain",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 0,
					Address:       sample.Address(r),
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        1,
					LaunchTriggered: false,
					IsMainnet:       true,
					CoordinatorID:   1,
				},
				fee: sdk.Coins(nil),
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 1),
			err: types.ErrInvalidRequestForMainnet,
		},
		{
			name: "should prevent sending a request for a chain where coordinator is not found",
			inputState: inputState{
				noAccount:     true,
				noCoordinator: true,
				chain: types.Chain{
					LaunchID:        2,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   2,
				},
				fee: sdk.Coins(nil),
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 2),
			err: types.ErrChainInactive,
		},
		{
			name: "should prevent sending a request if it is sent by coordinator and can't be applied",
			inputState: inputState{
				account: types.GenesisAccount{
					Address:  addr,
					LaunchID: 3,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 3,
					Address:       coordAddr,
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        3,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   3,
				},
				fee: sdk.Coins(nil),
			},
			msg: sample.MsgSendRequestWithAddAccount(r, coordAddr, addr, 3),
			err: types.ErrRequestApplicationFailure,
		},
		{
			name: "should prevent sending a request for chain with inactive coordinator",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 4,
					Address:       sample.Address(r),
					Active:        false,
				},
				chain: types.Chain{
					LaunchID:        4,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   4,
				},
				fee: sdk.Coins(nil),
			},
			msg: sample.MsgSendRequestWithAddAccount(r, sample.Address(r), sample.Address(r), 4),
			err: profiletypes.ErrCoordinatorInactive,
		},
		{
			name: "should allow send a new request",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 5,
					Address:       coordAddr,
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        5,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   5,
				},
				fee: sdk.Coins(nil),
			},
			msg: *types.NewMsgSendRequest(
				sample.Address(r),
				5,
				types.NewAccountRemoval(sample.Address(r)),
			),
			wantID:      1,
			wantApprove: false,
		},
		{
			name: "should allow send a new request from the coordinator and apply it",
			inputState: inputState{
				noAccount:     true,
				noCoordinator: true,
				noChain:       true,
				fee:           sdk.Coins(nil),
			},
			msg:         sample.MsgSendRequestWithAddAccount(r, coordAddr, sample.Address(r), 5),
			wantID:      2,
			wantApprove: true,
		},
		{
			name: "should allow send a new valid request for a mainnet chain",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 6,
					Address:       sample.Address(r),
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        6,
					LaunchTriggered: false,
					IsMainnet:       true,
					CoordinatorID:   6,
				},
				fee: sdk.Coins(nil),
			},
			msg: *types.NewMsgSendRequest(
				sample.Address(r),
				6,
				types.NewValidatorRemoval(sample.Address(r)),
			),
			wantID:      1,
			wantApprove: false,
		},
		{
			name: "should prevent send a new request if sender has no balance",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 5,
					Address:       coordAddr,
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        5,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   5,
				},
				fee: sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1000))),
			},
			msg: *types.NewMsgSendRequest(
				sample.Address(r),
				5,
				types.NewAccountRemoval(sample.Address(r)),
			),
			wantApprove: false,
			err:         types.ErrFundCommunityPool,
		},
		{
			name: "should allow send a new request if sender sufficient balance",
			inputState: inputState{
				noAccount: true,
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 5,
					Address:       coordAddr,
					Active:        true,
				},
				chain: types.Chain{
					LaunchID:        5,
					LaunchTriggered: false,
					IsMainnet:       false,
					CoordinatorID:   5,
				},
				fee: fee,
			},
			msg: *types.NewMsgSendRequest(
				hasFeeAddr,
				5,
				types.NewAccountRemoval(sample.Address(r)),
			),
			wantID:                   3,
			expectedCommunityPoolAmt: sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1000)),
			wantApprove:              false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noCoordinator {
				err := tk.ProfileKeeper.Coordinator.Set(ctx, tt.inputState.coordinator.CoordinatorID, tt.inputState.coordinator)
				require.NoError(t, err)
				addr, err := tk.ProfileKeeper.AddressCodec().StringToBytes(tt.inputState.coordinator.Address)
				require.NoError(t, err)
				err = tk.ProfileKeeper.CoordinatorByAddress.Set(ctx, addr, profiletypes.CoordinatorByAddress{
					CoordinatorID: tt.inputState.coordinator.CoordinatorID,
					Address:       tt.inputState.coordinator.Address,
				})
				require.NoError(t, err)
			}
			if !tt.inputState.noChain {
				err := tk.LaunchKeeper.Chain.Set(ctx, tt.inputState.chain.LaunchID, tt.inputState.chain)
				require.NoError(t, err)
			}
			if !tt.inputState.noAccount {
				addr, err := tk.LaunchKeeper.AddressCodec().StringToBytes(tt.inputState.account.Address)
				require.NoError(t, err)
				err = tk.LaunchKeeper.GenesisAccount.Set(ctx, collections.Join(tt.inputState.account.LaunchID, sdk.AccAddress(addr)), tt.inputState.account)
				require.NoError(t, err)
			}
			if !tt.inputState.fee.Empty() {
				params, err := tk.LaunchKeeper.Params.Get(ctx)
				require.NoError(t, err)
				params.RequestFee = tt.inputState.fee
				err = tk.LaunchKeeper.Params.Set(ctx, params)
				require.NoError(t, err)
			}

			// hook call assertion
			if tt.err == nil {
				tk.HooksMocks.LaunchHooksMock.On(
					"RequestCreated",
					ctx,
					tt.msg.Creator,
					tt.msg.LaunchID,
					tt.wantID,
					tt.msg.Content,
				).Return(nil).Once()
			}

			got, err := ts.LaunchSrv.SendRequest(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			request, err := tk.LaunchKeeper.Request.Get(ctx, collections.Join(tt.msg.LaunchID, got.RequestID))
			require.NoError(t, err, "request not found")

			if !tt.wantApprove {
				require.Equal(t, types.Request_PENDING, request.Status)
			} else {
				require.Equal(t, types.Request_APPROVED, request.Status)
			}

			if !tt.inputState.fee.Empty() {
				feePool, err := tk.DistrKeeper.FeePool.Get(ctx)
				require.NoError(t, err)
				for _, decCoin := range feePool.CommunityPool {
					coin := sdk.NewCoin(decCoin.Denom, decCoin.Amount.TruncateInt())
					require.Equal(t, tt.expectedCommunityPoolAmt, coin)
				}
			}
		})
	}
}
