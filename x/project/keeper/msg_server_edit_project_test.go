package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func TestMsgUpdateProjectName(t *testing.T) {
	var (
		coordAddr          = sample.Address(r)
		coordAddrNoProject = sample.Address(r)
		project            = sample.Project(r, 0)

		ctx, tk, ts = testkeeper.NewTestSetup(t)
	)

	params, err := tk.ProjectKeeper.Params.Get(ctx)
	require.NoError(t, err)
	maxMetadataLength := params.MaxMetadataLength

	t.Run("should allow creation of coordinators", func(t *testing.T) {
		desc1 := sample.CoordinatorDescription(r)
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:  coordAddr,
			Identity: desc1.Identity,
			Details:  desc1.Details,
			Website:  desc1.Website,
		})
		require.NoError(t, err)
		project.CoordinatorId = res.CoordinatorId
		project.ProjectId, err = tk.ProjectKeeper.AppendProject(ctx, project)
		require.NoError(t, err)

		desc2 := sample.CoordinatorDescription(r)
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:  coordAddrNoProject,
			Identity: desc2.Identity,
			Details:  desc2.Details,
			Website:  desc2.Website,
		})
		require.NoError(t, err)
	})

	for _, tc := range []struct {
		name string
		msg  types.MsgEditProject
		err  error
	}{
		{
			name: "should allow edit name and metadata",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectId:   project.ProjectId,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should allow edit name",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectId:   project.ProjectId,
				Name:        sample.ProjectName(r),
				Metadata:    []byte{},
			},
		},
		{
			name: "should allow edit metadata",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectId:   project.ProjectId,
				Name:        "",
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should fail if invalid project id",
			msg: types.MsgEditProject{
				Coordinator: coordAddr,
				ProjectId:   100,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with invalid coordinator address",
			msg: types.MsgEditProject{
				Coordinator: sample.Address(r),
				ProjectId:   project.ProjectId,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should fail with wrong coordinator id",
			msg: types.MsgEditProject{
				Coordinator: coordAddrNoProject,
				ProjectId:   project.ProjectId,
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "should fail when the change had too long metadata",
			msg: types.MsgEditProject{
				ProjectId:   0,
				Coordinator: sample.Address(r),
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, maxMetadataLength+1),
			},
			err: types.ErrInvalidMetadataLength,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			previousProject, err := tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectId)
			if err != nil && tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			_, err = ts.ProjectSrv.EditProject(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			project, err := tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectId)
			require.NoError(t, err)

			if len(tc.msg.Name) > 0 {
				require.EqualValues(t, tc.msg.Name, project.ProjectName)
			} else {
				require.EqualValues(t, previousProject.ProjectName, project.ProjectName)
			}

			if len(tc.msg.Metadata) > 0 {
				require.EqualValues(t, tc.msg.Metadata, project.Metadata)
			} else {
				require.EqualValues(t, previousProject.Metadata, project.Metadata)
			}
		})
	}
}
