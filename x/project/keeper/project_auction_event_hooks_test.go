package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	tc "github.com/ignite/network/testutil/constructor"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

func TestKeeper_EmitProjectAuctionCreated(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	type inputState struct {
		noProject     bool
		noCoordinator bool
		project       types.Project
		coordinator   profiletypes.Coordinator
	}

	coordinator := sample.Address(r)

	tests := []struct {
		name        string
		inputState  inputState
		auctionId   uint64
		auctioneer  string
		sellingCoin sdk.Coin
		emitted     bool
		err         error
	}{
		{
			name: "should prevent emitting event if selling coin is not a voucher",
			inputState: inputState{
				noProject:     true,
				noCoordinator: true,
			},
			sellingCoin: tc.Coin(t, "1000foo"),
			emitted:     false,
		},
		{
			name: "should return error if selling coin is a voucher of a non existing project",
			inputState: inputState{
				noProject:     true,
				noCoordinator: true,
			},
			sellingCoin: tc.Coin(t, "1000"+types.VoucherDenom(5, "foo")),
			err:         types.ErrProjectNotFound,
		},
		{
			name: "should return error if selling coin is a voucher of a project with non existing coordinator",
			inputState: inputState{
				project: types.Project{
					ProjectID:     10,
					CoordinatorID: 20,
				},
				noCoordinator: true,
			},
			sellingCoin: tc.Coin(t, "1000"+types.VoucherDenom(10, "foo")),
			err:         profiletypes.ErrCoordinatorInvalid,
		},
		{
			name: "should prevent emitting event if the auctioneer is not the coordinator of the project",
			inputState: inputState{
				project: types.Project{
					ProjectID:     100,
					CoordinatorID: 200,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 200,
					Address:       sample.Address(r),
				},
			},
			auctioneer:  sample.Address(r),
			sellingCoin: tc.Coin(t, "1000"+types.VoucherDenom(100, "foo")),
			emitted:     false,
		},
		{
			name: "should allow emitting event if the auctioneer is the coordinator of the project",
			inputState: inputState{
				project: types.Project{
					ProjectID:     1000,
					CoordinatorID: 2000,
				},
				coordinator: profiletypes.Coordinator{
					CoordinatorID: 2000,
					Address:       coordinator,
				},
			},
			auctioneer:  coordinator,
			sellingCoin: tc.Coin(t, "1000"+types.VoucherDenom(1000, "foo")),
			emitted:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// initialize input state
			if !tt.inputState.noProject {
				err := tk.ProjectKeeper.Project.Set(ctx, tt.inputState.project.ProjectID, tt.inputState.project)
				require.NoError(t, err)
			}
			if !tt.inputState.noCoordinator {
				err := tk.ProfileKeeper.Coordinator.Set(ctx, tt.inputState.coordinator.CoordinatorID, tt.inputState.coordinator)
				require.NoError(t, err)
			}

			emitted, err := tk.ProjectKeeper.EmitProjectAuctionCreated(ctx, tt.auctionId, tt.auctioneer, tt.sellingCoin)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
				require.EqualValues(t, tt.emitted, emitted)
			}

			// clean state
			if !tt.inputState.noProject {
				err := tk.ProjectKeeper.Project.Remove(ctx, tt.inputState.project.ProjectID)
				require.NoError(t, err)
			}
			if !tt.inputState.noCoordinator {
				err := tk.ProfileKeeper.Coordinator.Remove(ctx, tt.inputState.coordinator.CoordinatorID)
				require.NoError(t, err)
			}
		})
	}
}
