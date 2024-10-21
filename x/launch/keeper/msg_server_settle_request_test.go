package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	ignterrors "github.com/ignite/network/pkg/errors"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func TestMsgSettleRequest(t *testing.T) {
	const numReq = 6

	var (
		coordinator1       = sample.Coordinator(r, sample.Address(r))
		coordinator2       = sample.Coordinator(r, sample.Address(r))
		disableCoordinator = sample.Coordinator(r, sample.Address(r))
		invalidChain       = uint64(1000)
		ctx, tk, ts        = testkeeper.NewTestSetup(t)
	)

	disableCoordinator.Active = false

	var err error
	coordinator1.CoordinatorId, err = tk.ProfileKeeper.AppendCoordinator(ctx, coordinator1)
	require.NoError(t, err)
	coordinator2.CoordinatorId, err = tk.ProfileKeeper.AppendCoordinator(ctx, coordinator2)
	require.NoError(t, err)
	disableCoordinator.CoordinatorId, err = tk.ProfileKeeper.AppendCoordinator(ctx, disableCoordinator)
	require.NoError(t, err)

	chains := createNChainForCoordinator(tk.LaunchKeeper, ctx, coordinator1.CoordinatorId, 4)
	chains[0].LaunchTriggered = true
	err = tk.LaunchKeeper.Chain.Set(ctx, chains[0].LaunchId, chains[0])
	require.NoError(t, err)
	chains[1].CoordinatorId = 99999
	err = tk.LaunchKeeper.Chain.Set(ctx, chains[1].LaunchId, chains[1])
	require.NoError(t, err)
	chains[3].CoordinatorId = disableCoordinator.CoordinatorId
	err = tk.LaunchKeeper.Chain.Set(ctx, chains[3].LaunchId, chains[3])
	require.NoError(t, err)

	requestSamples := make([]RequestSample, numReq)
	for i := 0; i < numReq; i++ {
		addr := sample.Address(r)
		requestSamples[i] = RequestSample{
			Content: sample.GenesisAccountContent(r, chains[2].LaunchId, addr),
			Creator: addr,
			Status:  types.Request_PENDING,
		}
	}

	// set one request to a non-pending status
	requestSamples[numReq-1].Status = types.Request_APPROVED
	requests := createRequestsFromSamples(tk.LaunchKeeper, ctx, chains[2].LaunchId, requestSamples)

	invalidContentRequest := types.Request{
		LaunchId: chains[2].LaunchId,
	}
	invalidContentRequestID, err := tk.LaunchKeeper.AppendRequest(ctx, invalidContentRequest)
	require.NoError(t, err)

	tests := []struct {
		name       string
		msg        types.MsgSettleRequest
		checkAddr  string
		wantStatus types.Request_Status
		err        error
	}{
		{
			name: "should prevent settling request for non existing chain",
			msg: types.MsgSettleRequest{
				LaunchId:  invalidChain,
				Signer:    coordinator1.Address,
				RequestId: requests[0].RequestId,
				Approve:   true,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent settling request with launch triggered chain",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[0].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: requests[0].RequestId,
				Approve:   true,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "should prevent settling request with coordinator not found",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[1].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: requests[0].RequestId,
				Approve:   true,
			},
			err: types.ErrChainInactive,
		},
		{
			name: "should prevent setting request if no address permission",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    coordinator2.Address,
				RequestId: requests[0].RequestId,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should prevent setting request if request already settled",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: requests[numReq-1].RequestId,
				Approve:   true,
			},
			err: types.ErrRequestSettled,
		},
		{
			name: "should prevent approving a request that does not exist",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: 99999999,
				Approve:   true,
			},
			err: types.ErrRequestNotFound,
		},
		{
			name: "should prevent applying a request with invalid contents",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: invalidContentRequestID,
				Approve:   true,
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should allow approving request from coordinator",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: requests[0].RequestId,
				Approve:   true,
			},
			wantStatus: types.Request_APPROVED,
			checkAddr:  requestSamples[0].Creator,
		},
		{
			name: "should allow approving a second request from coordinator",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: requests[1].RequestId,
				Approve:   true,
			},
			wantStatus: types.Request_APPROVED,
			checkAddr:  requestSamples[1].Creator,
		},
		{
			name: "should allow rejecting request from coordinator",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    coordinator1.Address,
				RequestId: requests[2].RequestId,
				Approve:   false,
			},
			wantStatus: types.Request_REJECTED,
			checkAddr:  requestSamples[2].Creator,
		},
		{
			name: "should allow rejecting request from request creator",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    requestSamples[3].Creator,
				RequestId: requests[3].RequestId,
				Approve:   false,
			},
			wantStatus: types.Request_REJECTED,
			checkAddr:  requestSamples[3].Creator,
		},
		{
			name: "should prevent rejecting a request from an account other than coordinator and request creator",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    requestSamples[3].Creator,
				RequestId: requests[4].RequestId,
				Approve:   false,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should prevent approving a request from an account other than coordinator",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[2].LaunchId,
				Signer:    requestSamples[5].Creator,
				RequestId: requests[5].RequestId,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should prevent settling request from a disabled coordinator",
			msg: types.MsgSettleRequest{
				LaunchId:  chains[3].LaunchId,
				Signer:    disableCoordinator.Address,
				RequestId: requests[5].RequestId,
				Approve:   true,
			},
			err: profiletypes.ErrCoordinatorInactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ts.LaunchSrv.SettleRequest(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			request, err := tk.LaunchKeeper.Request.Get(ctx, collections.Join(tt.msg.LaunchId, tt.msg.RequestId))
			require.NoError(t, err, "request not found")
			require.Equal(t, tt.wantStatus, request.Status)

			checkAddr, err := tk.LaunchKeeper.AddressCodec().StringToBytes(tt.checkAddr)
			_, err = tk.LaunchKeeper.GenesisAccount.Get(ctx, collections.Join(tt.msg.LaunchId, sdk.AccAddress(checkAddr)))
			if tt.msg.Approve {
				require.NoError(t, err, "request apply performed")
			} else {
				require.Error(t, err, "request apply not performed")
			}
		})
	}
}
