package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
	projecttypes "github.com/ignite/network/x/project/types"
)

func TestMsgEditChain(t *testing.T) {
	var (
		coordNoExist    = sample.Address(r)
		launchIDNoExist = uint64(1000)
		ctx, tk, ts     = testkeeper.NewTestSetup(t)
	)

	// Create coordinators
	_, coordAddr := ts.CreateCoordinator(ctx, r)
	_, coordAddr2 := ts.CreateCoordinator(ctx, r)

	coordAddress := coordAddr.String()
	coordAddress2 := coordAddr2.String()

	// Create a chain
	launchID := ts.CreateChain(ctx, r, coordAddress, "", false, 0)

	// create a project
	projectID := ts.CreateProject(ctx, r, coordAddress)

	// create a chain with an existing project
	launchIDHasProject := ts.CreateChain(ctx, r, coordAddress, "", true, projectID)

	// create a project
	validProjectID := ts.CreateProject(ctx, r, coordAddress)

	// create a project from a different address
	projectDifferentCoordinator := ts.CreateProject(ctx, r, coordAddress2)

	// Create a new chain for more tests
	launchID2 := ts.CreateChain(ctx, r, coordAddress, "", false, 0)

	// create a new project and add a chainProjects entry to it
	projectDuplicateChain := ts.CreateProject(ctx, r, coordAddress)

	err := tk.ProjectKeeper.AddChainToProject(ctx, projectDuplicateChain, launchID2)
	require.NoError(t, err)

	// create message with an invalid metadata length
	msgEditChainInvalidMetadata := sample.MsgEditChain(r,
		coordAddress,
		launchID,
		true,
		validProjectID,
		false,
	)
	launchParams, err := tk.LaunchKeeper.Params.Get(ctx)
	require.NoError(t, err)
	msgEditChainInvalidMetadata.Metadata = sample.Metadata(r, launchParams.MaxMetadataLength+1)

	for _, tc := range []struct {
		name string
		msg  types.MsgEditChain
		err  error
	}{
		{
			name: "should allow setting a project ID",
			msg: sample.MsgEditChain(r,
				coordAddress,
				launchID,
				true,
				validProjectID,
				false,
			),
		},
		{
			name: "should allow editing metadata",
			msg: sample.MsgEditChain(r,
				coordAddress,
				launchID,
				false,
				0,
				true,
			),
		},
		{
			name: "should prevent editing chain from non existent launch id",
			msg: sample.MsgEditChain(r,
				coordAddress,
				launchIDNoExist,
				true,
				0,
				false,
			),
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent editing chain with non existent coordinator",
			msg: sample.MsgEditChain(r,
				coordNoExist,
				launchID,
				true,
				0,
				false,
			),
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should prevent editing chain with invalid coordinator",
			msg: sample.MsgEditChain(r,
				coordAddress2,
				launchID,
				true,
				0,
				false,
			),
			err: profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "should prevent setting project id for chain with a project",
			msg: sample.MsgEditChain(r,
				coordAddress,
				launchIDHasProject,
				true,
				0,
				false,
			),
			err: types.ErrChainHasProject,
		},
		{
			name: "should prevent setting project id where project does not exist",
			msg: sample.MsgEditChain(r,
				coordAddress,
				launchID2,
				true,
				999,
				false,
			),
			err: projecttypes.ErrProjectNotFound,
		},
		{
			name: "should prevent setting project id where project has a different coordinator",
			msg: sample.MsgEditChain(r,
				coordAddress,
				launchID2,
				true,
				projectDifferentCoordinator,
				false,
			),
			err: profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "should prevent setting project id where project chain entry is duplicated",
			msg: sample.MsgEditChain(r,
				coordAddress,
				launchID2,
				true,
				projectDuplicateChain,
				false,
			),
			err: types.ErrAddChainToProject,
		},
		{
			name: "should prevent edit a chain with invalid metadata length",
			msg:  msgEditChainInvalidMetadata,
			err:  types.ErrInvalidMetadataLength,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Fetch the previous state of the chain to perform checks
			var previousChain types.Chain
			var found bool
			if tc.err == nil {
				previousChain, err = tk.LaunchKeeper.GetChain(ctx, tc.msg.LaunchID)
				require.NoError(t, err)
			}

			// Send the message
			_, err := ts.LaunchSrv.EditChain(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// The chain must continue to exist in the store
			chain, err := tk.LaunchKeeper.GetChain(ctx, tc.msg.LaunchID)
			require.NoError(t, err)

			// Unchanged values
			require.EqualValues(t, previousChain.CoordinatorID, chain.CoordinatorID)
			require.EqualValues(t, previousChain.CreatedAt, chain.CreatedAt)
			require.EqualValues(t, previousChain.LaunchTime, chain.LaunchTime)
			require.EqualValues(t, previousChain.LaunchTriggered, chain.LaunchTriggered)

			if len(tc.msg.Metadata) > 0 {
				require.EqualValues(t, tc.msg.Metadata, chain.Metadata)
			} else {
				require.EqualValues(t, previousChain.Metadata, chain.Metadata)
			}

			if tc.msg.SetProjectID {
				require.True(t, chain.HasProject)
				require.EqualValues(t, tc.msg.ProjectID, chain.ProjectID)
				// ensure project exist
				_, err := tk.ProjectKeeper.GetProject(ctx, chain.ProjectID)
				require.NoError(t, err)
				// ensure project chains exist
				projectChains, err := tk.ProjectKeeper.GetProjectChains(ctx, chain.ProjectID)
				require.NoError(t, err)

				// check that the chain launch ID is in the project chains
				found = false
				for _, chainID := range projectChains.Chains {
					if chainID == chain.LaunchID {
						found = true
						break
					}
				}

				require.True(t, found)
			}
		})
	}
}
