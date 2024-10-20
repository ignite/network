package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func TestMsgUpdateTotalSupply(t *testing.T) {
	var (
		coordID     uint64
		coordAddr1  = sample.Address(r)
		coordAddr2  = sample.Address(r)
		ctx, tk, ts = testkeeper.NewTestSetup(t)
	)

	t.Run("should allow creating coordinators", func(t *testing.T) {
		res, err := ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr1,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
		coordID = res.CoordinatorId
		res, err = ts.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
			Address:     coordAddr2,
			Description: sample.CoordinatorDescription(r),
		})
		require.NoError(t, err)
	})

	// Set a regular project and a project with an already initialized mainnet
	project := sample.Project(r, 0)
	project.CoordinatorId = coordID
	require.NoError(t, tk.ProjectKeeper.Project.Set(ctx, 0, project))

	project = sample.Project(r, 1)
	project.CoordinatorId = coordID
	project.MainnetInitialized = true
	require.NoError(t, tk.ProjectKeeper.Project.Set(ctx, 1, project))

	for _, tc := range []struct {
		name string
		msg  types.MsgUpdateTotalSupply
		err  error
	}{
		{
			name: "should update total supply",
			msg: types.MsgUpdateTotalSupply{
				ProjectId:         0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should allow update total supply again",
			msg: types.MsgUpdateTotalSupply{
				ProjectId:         0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should fail if project not found",
			msg: types.MsgUpdateTotalSupply{
				ProjectId:         100,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: types.ErrProjectNotFound,
		},
		{
			name: "should fail with non existing coordinator",
			msg: types.MsgUpdateTotalSupply{
				ProjectId:         0,
				Coordinator:       sample.Address(r),
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: profiletypes.ErrCoordinatorAddressNotFound,
		},
		{
			name: "should fail if coordinator is not associated with project",
			msg: types.MsgUpdateTotalSupply{
				ProjectId:         0,
				Coordinator:       coordAddr2,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "cannot update total supply when mainnet is initialized",
			msg: types.MsgUpdateTotalSupply{
				ProjectId:         1,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
			err: types.ErrMainnetInitialized,
		},
		{
			name: "should fail if total supply outside of valid range",
			msg: types.MsgUpdateTotalSupply{
				ProjectId:         0,
				Coordinator:       coordAddr1,
				TotalSupplyUpdate: sample.CoinsWithRange(r, 10, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var previousTotalSupply sdk.Coins
			if tc.err == nil {
				project, err := tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectId)
				require.NoError(t, err)
				previousTotalSupply = project.TotalSupply
			}

			_, err := ts.ProjectSrv.UpdateTotalSupply(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			project, err := tk.ProjectKeeper.GetProject(ctx, tc.msg.ProjectId)
			require.NoError(t, err)
			require.True(t, project.TotalSupply.Equal(
				types.UpdateTotalSupply(previousTotalSupply, tc.msg.TotalSupplyUpdate),
			))
		})
	}
}
