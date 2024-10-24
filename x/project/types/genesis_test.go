package types_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"

	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	tc2 "github.com/ignite/network/testutil/constructor"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/types"
)

func TestGenesisState_Validate(t *testing.T) {
	var (
		project1 = sample.Project(r, 0)
		project2 = sample.Project(r, 1)
		shares1  = sample.Shares(r)
		shares2  = sample.Shares(r)
		shares3  = sample.Shares(r)
		shares4  = sample.Shares(r)
	)
	sharesProject1 := types.IncreaseShares(shares1, shares2)
	project1.AllocatedShares = sharesProject1
	project1.CoordinatorId = 0

	sharesProject2 := types.IncreaseShares(shares3, shares4)
	project2.AllocatedShares = sharesProject2
	project2.CoordinatorId = 1

	for _, tc := range []struct {
		name         string
		genState     *types.GenesisState
		errorMessage string
	}{
		{
			name:     "should allow validation of valid default genesis",
			genState: types.DefaultGenesis(),
		},
		{
			name: "should allow validation of valid genesis",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
				ProjectChainsList: []types.ProjectChains{
					{
						ProjectId: project1.ProjectId,
					},
					{
						ProjectId: project2.ProjectId,
					},
				},
				ProjectList: []types.Project{
					project1,
					project2,
				},
				ProjectCount: 2,
				MainnetAccountList: []types.MainnetAccount{
					{
						ProjectId: project1.ProjectId,
						Address:   sample.Address(r),
						Shares:    shares1,
					},
					{
						ProjectId: project2.ProjectId,
						Address:   sample.Address(r),
						Shares:    shares3,
					},
				},
				TotalShares: networktypes.TotalShareNumber,
				Params:      types.DefaultParams(),
			},
		},
		{
			name: "should prevent validation of genesis with non existing project for mainnet account",
			genState: &types.GenesisState{
				ProjectChainsList: []types.ProjectChains{
					{
						ProjectId: 0,
					},
					{
						ProjectId: 1,
					},
				},
				ProjectList: []types.Project{
					sample.Project(r, 0),
					sample.Project(r, 1),
				},
				ProjectCount: 2,
				MainnetAccountList: []types.MainnetAccount{
					sample.MainnetAccount(r, 330, "330"),
				},
				TotalShares: networktypes.TotalShareNumber,
			},
			errorMessage: "project id 330 doesn't exist for mainnet account 330",
		},
		{
			name: "should prevent validation of genesis with non existing project for chains",
			genState: &types.GenesisState{
				ProjectChainsList: []types.ProjectChains{
					{
						ProjectId: 2,
					},
					{
						ProjectId: 4,
					},
				},
				ProjectList: []types.Project{
					sample.Project(r, 99),
					sample.Project(r, 88),
				},
				ProjectCount: 100,
				TotalShares:  networktypes.TotalShareNumber,
			},
			errorMessage: "project id 2 doesn't exist for chains",
		},
		{
			name: "should prevent validation of genesis with duplicated projectChains",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					sample.Project(r, 0),
				},
				ProjectCount: 1,
				ProjectChainsList: []types.ProjectChains{
					{
						ProjectId: 0,
					},
					{
						ProjectId: 0,
					},
				},
				TotalShares: networktypes.TotalShareNumber,
			},
			errorMessage: "duplicated index for projectChains",
		},
		{
			name: "should prevent validation of genesis with duplicated project",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					sample.Project(r, 0),
					sample.Project(r, 0),
				},
				ProjectCount: 2,
				TotalShares:  networktypes.TotalShareNumber,
			},
			errorMessage: "duplicated id for project",
		},
		{
			name: "should prevent validation of genesis with invalid project count",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					sample.Project(r, 1),
				},
				ProjectCount: 0,
				TotalShares:  networktypes.TotalShareNumber,
			},
			errorMessage: "project id should be lower or equal than the last id",
		},
		{
			name: "should prevent validation of genesis with invalid project",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					types.NewProject(
						0,
						invalidProjectName,
						sample.Uint64(r),
						sample.TotalSupply(r),
						sample.Metadata(r, 20),
						sample.Duration(r).Milliseconds(),
					),
				},
				ProjectCount: 1,
				TotalShares:  networktypes.TotalShareNumber,
			},
			errorMessage: "invalid project 0: project name can only contain alphanumerical characters or hyphen",
		},
		{
			name: "should prevent validation of genesis with duplicated mainnetAccount",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					sample.Project(r, 0),
				},
				ProjectCount: 1,
				MainnetAccountList: []types.MainnetAccount{
					{
						ProjectId: 0,
						Address:   "0",
					},
					{
						ProjectId: 0,
						Address:   "0",
					},
				},
				TotalShares: networktypes.TotalShareNumber,
			},
			errorMessage: "duplicated index for mainnetAccount",
		},
		{
			name: "should prevent validation of genesis with invalid allocations",
			genState: &types.GenesisState{
				ProjectList: []types.Project{
					{
						ProjectId:          0,
						ProjectName:        "test",
						CoordinatorId:      0,
						MainnetId:          0,
						MainnetInitialized: false,
						TotalSupply:        nil,
						AllocatedShares:    types.NewSharesFromCoins(tc2.Coins(t, fmt.Sprintf("%dstake", networktypes.TotalShareNumber+1))),
						Metadata:           nil,
					},
				},
				ProjectCount: 1,
				MainnetAccountList: []types.MainnetAccount{
					{
						ProjectId: 0,
						Address:   "0",
					},
				},
				TotalShares: networktypes.TotalShareNumber,
			},
			errorMessage: "invalid project 0: more allocated shares than total shares",
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.errorMessage != "" {
				require.Error(t, err)
				require.Equal(t, tc.errorMessage, err.Error())
				return
			}
			require.NoError(t, err)

			projectIDMap := make(map[uint64]types.Shares)
			for _, elem := range tc.genState.ProjectList {
				projectIDMap[elem.ProjectId] = elem.AllocatedShares
			}
			shares := make(map[uint64]types.Shares)

			for _, acc := range tc.genState.MainnetAccountList {
				// check if the project exists for mainnet accounts
				_, ok := projectIDMap[acc.ProjectId]
				require.True(t, ok)

				// sum mainnet account shares
				if _, ok := shares[acc.ProjectId]; !ok {
					shares[acc.ProjectId] = types.EmptyShares()
				}
				shares[acc.ProjectId] = types.IncreaseShares(
					shares[acc.ProjectId],
					acc.Shares,
				)
			}

			for projectID, share := range projectIDMap {
				// check if the project shares is equal all accounts shares
				accShares, ok := shares[projectID]
				require.True(t, ok)
				isLowerEqual := accShares.IsAllLTE(share)
				require.True(t, isLowerEqual)
			}
		})
	}
}

func TestGenesisState_ValidateParams(t *testing.T) {
	for _, tc := range []struct {
		name     string
		genState types.GenesisState
		valid    bool
	}{
		{
			name: "should prevent validation of genesis with max total supply below min total supply",
			genState: types.GenesisState{
				Params: types.NewParams(
					types.DefaultMinTotalSupply,
					types.DefaultMinTotalSupply.Sub(sdkmath.OneInt()),
					types.DefaultProjectCreationFee,
					types.DefaultMaxMetadataLength,
				),
			},
			valid: false,
		},
		{
			name: "should prevent validation of genesis with valid parameters",
			genState: types.GenesisState{
				Params: types.NewParams(
					types.DefaultMinTotalSupply,
					types.DefaultMinTotalSupply.Add(sdkmath.OneInt()),
					types.DefaultProjectCreationFee,
					types.DefaultMaxMetadataLength,
				),
			},
			valid: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
