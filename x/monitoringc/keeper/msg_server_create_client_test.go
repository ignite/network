package keeper_test

import (
	"encoding/base64"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	launchtypes "github.com/ignite/network/x/launch/types"
	"github.com/ignite/network/x/monitoringc/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func Test_msgServer_CreateClient(t *testing.T) {
	var (
		coordAddr    = sample.Address(r)
		invalidChain = uint64(1000)

		ctx, tk, ts = testkeeper.NewTestSetup(t)

		consPubKeyStr = "jP0v8F0e2kSAS367V/QAikddQPze+V36v7lhkv1Iqgg="
		cs            = networktypes.NewConsensusState(
			"2022-02-08T15:12:36.161481Z",
			"A13E761948413E405EA4F09BEC9F37632F739404108FE1635CB3529B61DA9FD7",
			"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		)
		vs = networktypes.NewValidatorSet(
			networktypes.NewValidator(consPubKeyStr, 0, 100),
		)
	)

	selfDelegation, err := sdk.ParseCoinNormalized("1000stake")
	require.NoError(t, err)
	consPubKey, err := base64.StdEncoding.DecodeString(consPubKeyStr)
	require.NoError(t, err)

	// create a coordinator and a chain with a genesis validator
	resCoord, err := ts.ProfileSrv.CreateCoordinator(ctx, profiletypes.NewMsgCreateCoordinator(
		coordAddr,
		"",
		"",
		"",
	))
	require.NoError(t, err)
	initialGenesis := launchtypes.NewDefaultInitialGenesis()
	resCreateChain, err := ts.LaunchSrv.CreateChain(ctx, launchtypes.NewMsgCreateChain(
		coordAddr,
		"orbit-1",
		sample.String(r, 10),
		sample.String(r, 10),
		initialGenesis,
		false,
		0,
		sample.Coins(r),
		sample.Metadata(r, 20),
	))
	require.NoError(t, err)
	chainWithInvalidChainID := sample.Chain(r, resCreateChain.LaunchId+1, resCoord.CoordinatorId)
	chainWithInvalidChainID.GenesisChainId = "invalid_chain_id"
	err = tk.LaunchKeeper.Chain.Set(ctx, chainWithInvalidChainID.LaunchId, chainWithInvalidChainID)
	require.NoError(t, err)
	_, err = ts.LaunchSrv.SendRequest(ctx, launchtypes.NewMsgSendRequest(
		coordAddr,
		resCreateChain.LaunchId,
		launchtypes.NewGenesisValidator(
			resCreateChain.LaunchId,
			sample.Address(r),
			sample.Bytes(r, 100),
			consPubKey,
			selfDelegation,
			sample.GenesisValidatorPeer(r),
		),
	))
	require.NoError(t, err)
	_, err = ts.LaunchSrv.TriggerLaunch(ctx, launchtypes.NewMsgTriggerLaunch(
		coordAddr,
		resCreateChain.LaunchId,
		ctx.BlockTime().Add(launchtypes.DefaultMinLaunchTime),
	))
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgCreateClient
		err  error
	}{
		{
			name: "invalid chain ID",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				chainWithInvalidChainID.LaunchId,
				cs,
				vs,
				networktypes.DefaultUnbondingPeriod,
				networktypes.DefaultRevisionHeight,
			),
			err: types.ErrInvalidClientState,
		},
		{
			name: "invalid consensus state",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				resCreateChain.LaunchId,
				networktypes.NewConsensusState(
					"",
					"",
					"",
				),
				vs,
				networktypes.DefaultUnbondingPeriod,
				networktypes.DefaultRevisionHeight,
			),
			err: types.ErrInvalidConsensusState,
		},
		{
			name: "chain doesn't exist",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				invalidChain,
				cs,
				vs,
				networktypes.DefaultUnbondingPeriod,
				networktypes.DefaultRevisionHeight,
			),
			err: launchtypes.ErrChainNotFound,
		},
		{
			name: "empty validator set",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				resCreateChain.LaunchId,
				sample.ConsensusState(0),
				networktypes.ValidatorSet{},
				networktypes.DefaultUnbondingPeriod,
				networktypes.DefaultRevisionHeight,
			),
			err: types.ErrInvalidValidatorSet,
		},
		{
			name: "invalid validator set",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				resCreateChain.LaunchId,
				sample.ConsensusState(0),
				sample.ValidatorSet(1),
				networktypes.DefaultUnbondingPeriod,
				networktypes.DefaultRevisionHeight,
			),
			err: types.ErrInvalidValidatorSetHash,
		},
		{
			name: "verified client should be created",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				resCreateChain.LaunchId,
				cs,
				vs,
				networktypes.DefaultUnbondingPeriod,
				networktypes.DefaultRevisionHeight,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := ts.MonitoringcSrv.CreateClient(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// verify the client is created
			verifiedClients, err := tk.MonitoringConsumerKeeper.VerifiedClientID.Get(ctx, tt.msg.LaunchId)
			require.NoError(t, err, "verified client ID should be added in the list")
			require.EqualValues(t, tt.msg.LaunchId, verifiedClients.LaunchId)
			require.Contains(t, verifiedClients.ClientIdList, res.ClientId)

			launchIDFromClient, err := tk.MonitoringConsumerKeeper.LaunchIDFromVerifiedClientID.Get(ctx, res.ClientId)
			require.NoError(t, err, "launch ID should be registered for the verified client ID")
			require.EqualValues(t, res.ClientId, launchIDFromClient.ClientId)
			require.EqualValues(t, tt.msg.LaunchId, launchIDFromClient.LaunchId)

			// IBC client should be created
			clientState, found := tk.IBCKeeper.ClientKeeper.GetClientState(ctx, res.ClientId)
			require.True(t, found, "IBC consumer client state should be created")
			cs, ok := clientState.(*ibctmtypes.ClientState)
			require.True(t, ok)
			require.EqualValues(t, tt.msg.RevisionHeight, cs.LatestHeight.RevisionHeight)
			require.EqualValues(t, time.Second*time.Duration(tt.msg.UnbondingPeriod), cs.UnbondingPeriod)
		})
	}
}
