package keeper_test

import (
	"testing"
	"time"

	ibctmtypes "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/monitoringp/types"
)

func TestKeeper_InitializeConsumerClient(t *testing.T) {
	t.Run("initialize consumer client", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

		// set params with valid values
		err := tk.MonitoringProviderKeeper.Params.Set(ctx, types.NewParams(
			1000,
			types.DefaultConsumerChainID,
			sample.ConsensusState(0),
			networktypes.DefaultUnbondingPeriod,
			networktypes.DefaultRevisionHeight,
		))
		require.NoError(t, err)
		clientID, err := tk.MonitoringProviderKeeper.InitializeConsumerClient(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, clientID)

		consumerClientID, err := tk.MonitoringProviderKeeper.ConsumerClientID.Get(ctx)
		require.NoError(t, err, "consumer client ID should be registered in the store")
		require.EqualValues(t, clientID, consumerClientID.ClientId)

		// IBC client should be created
		clientState, found := tk.IBCKeeper.ClientKeeper.GetClientState(ctx, clientID)
		require.True(t, found, "IBC consumer client state should be created")

		cs, ok := clientState.(*ibctmtypes.ClientState)
		require.True(t, ok)

		params, err := tk.MonitoringProviderKeeper.Params.Get(ctx)
		require.NoError(t, err)
		require.EqualValues(t, params.ConsumerRevisionHeight, cs.LatestHeight.RevisionHeight)
		require.EqualValues(t, time.Second*time.Duration(params.ConsumerUnbondingPeriod), cs.UnbondingPeriod)
	})

	t.Run("invalid consumer consensus state", func(t *testing.T) {
		ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

		// default params contain an empty consensus state, therefore invalid
		_, err := tk.MonitoringProviderKeeper.InitializeConsumerClient(ctx)
		require.ErrorIs(t, err, types.ErrInvalidConsensusState)
	})
}
