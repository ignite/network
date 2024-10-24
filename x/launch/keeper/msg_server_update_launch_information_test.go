package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func TestMsgUpdateLaunchInformation(t *testing.T) {
	ctx, tk, ts := testkeeper.NewTestSetup(t)
	coordNoExist := sample.Address(r)
	launchIDNoExist := uint64(1000)

	// Create coordinators
	coordID, coordAddr := ts.CreateCoordinator(ctx, r)
	coordAddress := coordAddr.String()

	_, coordAddr2 := ts.CreateCoordinator(ctx, r)
	coordAddress2 := coordAddr2.String()

	// Create a chain
	launchID := uint64(1)
	chain := sample.Chain(r, launchID, coordID)
	err := tk.LaunchKeeper.Chain.Set(ctx, chain.LaunchId, chain)
	require.NoError(t, err)

	launchIDLaunchTriggered := uint64(2)
	chain = sample.Chain(r, launchIDLaunchTriggered, coordID)
	chain.LaunchTriggered = true
	err = tk.LaunchKeeper.Chain.Set(ctx, chain.LaunchId, chain)
	require.NoError(t, err)

	for _, tc := range []struct {
		name string
		msg  types.MsgUpdateLaunchInformation
		err  error
	}{
		{
			name: "should allow updating genesis chain ID",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress,
				launchID,
				true,
				false,
				false,
				false,
			),
		},
		{
			name: "should allow updating source",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress,
				launchID,
				false,
				true,
				false,
				false,
			),
		},
		{
			name: "should allow updating initial genesis with default genesis",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress,
				launchID,
				false,
				false,
				true,
				false,
			),
		},
		{
			name: "should allow updating initial genesis with genesis url",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress,
				launchID,
				false,
				false,
				true,
				true,
			),
		},
		{
			name: "should allow updating source and initial genesis",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress,
				launchID,
				false,
				true,
				true,
				true,
			),
		},
		{
			name: "should prevent updating for non existent launch id",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress,
				launchIDNoExist,
				false,
				true,
				false,
				false,
			),
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent updating from non existent coordinator",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordNoExist,
				launchID,
				false,
				true,
				false,
				false,
			),
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should prevent updating from invalid coordinator",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress2,
				launchID,
				false,
				true,
				false,
				false,
			),
			err: profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "should prevent updating if chain launch already triggered",
			msg: sample.MsgUpdateLaunchInformation(r,
				coordAddress,
				launchIDLaunchTriggered,
				false,
				true,
				false,
				false,
			),
			err: types.ErrTriggeredLaunch,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Fetch the previous state of the chain to perform checks
			var previousChain types.Chain
			var err error
			if tc.err == nil {
				previousChain, err = tk.LaunchKeeper.GetChain(ctx, tc.msg.LaunchId)
				require.NoError(t, err)
			}

			// Send the message
			_, err = ts.LaunchSrv.UpdateLaunchInformation(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// The chain must continue to exist in the store
			chain, err := tk.LaunchKeeper.GetChain(ctx, tc.msg.LaunchId)
			require.NoError(t, err)

			// Unchanged values
			require.EqualValues(t, previousChain.CoordinatorId, chain.CoordinatorId)
			require.EqualValues(t, previousChain.CreatedAt, chain.CreatedAt)
			require.EqualValues(t, previousChain.LaunchTime, chain.LaunchTime)
			require.EqualValues(t, previousChain.LaunchTriggered, chain.LaunchTriggered)

			// Compare changed values
			if tc.msg.GenesisChainId != "" {
				require.EqualValues(t, tc.msg.GenesisChainId, chain.GenesisChainId)
			} else {
				require.EqualValues(t, previousChain.GenesisChainId, chain.GenesisChainId)
			}
			if tc.msg.SourceUrl != "" {
				require.EqualValues(t, tc.msg.SourceUrl, chain.SourceUrl)
				require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)
			} else {
				require.EqualValues(t, previousChain.SourceUrl, chain.SourceUrl)
				require.EqualValues(t, previousChain.SourceHash, chain.SourceHash)
			}

			if tc.msg.InitialGenesis != nil {
				require.EqualValues(t, *tc.msg.InitialGenesis, chain.InitialGenesis)
			} else {
				require.EqualValues(t, previousChain.InitialGenesis, chain.InitialGenesis)
			}
		})
	}
}
