// Package networksuite provides base test suite for tests that need a local network instance
package networksuite

import (
	"math/rand"
	"strconv"

	sdkmath "cosmossdk.io/math"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	claim "github.com/ignite/modules/x/claim/types"

	"github.com/ignite/network/testutil/network"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	launch "github.com/ignite/network/x/launch/types"
	monitoringc "github.com/ignite/network/x/monitoringc/types"
	participation "github.com/ignite/network/x/participation/types"
	profile "github.com/ignite/network/x/profile/types"
	project "github.com/ignite/network/x/project/types"
	reward "github.com/ignite/network/x/reward/types"
)

// NetworkTestSuite is a test suite for query tests that initializes a network instance
type NetworkTestSuite struct {
	suite.Suite
	Network            *network.Network
	LaunchState        launch.GenesisState
	ProjectState       project.GenesisState
	ClaimState         claim.GenesisState
	MonitoringcState   monitoringc.GenesisState
	ParticipationState participation.GenesisState
	ProfileState       profile.GenesisState
	RewardState        reward.GenesisState
}

// SetupSuite setups the local network with a genesis state
func (nts *NetworkTestSuite) SetupSuite() {
	var (
		r   = sample.Rand()
		cfg = network.DefaultConfig()
	)

	updateGenesisConfigState := func(moduleName string, moduleState proto.Message) {
		buf, err := cfg.Codec.MarshalJSON(moduleState)
		require.NoError(nts.T(), err)
		cfg.GenesisState[moduleName] = buf
	}

	// initialize launch
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[launch.ModuleName], &nts.LaunchState))
	nts.LaunchState = populateLaunch(r, nts.LaunchState)
	updateGenesisConfigState(launch.ModuleName, &nts.LaunchState)

	// initialize project
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[project.ModuleName], &nts.ProjectState))
	nts.ProjectState = populateProject(r, nts.ProjectState)
	updateGenesisConfigState(project.ModuleName, &nts.ProjectState)

	// initialize claim
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[claim.ModuleName], &nts.ClaimState))
	nts.ClaimState = populateClaim(r, nts.ClaimState)
	updateGenesisConfigState(claim.ModuleName, &nts.ClaimState)

	// initialize monitoring consumer
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[monitoringc.ModuleName], &nts.MonitoringcState))
	nts.MonitoringcState = populateMonitoringc(nts.MonitoringcState)
	updateGenesisConfigState(monitoringc.ModuleName, &nts.MonitoringcState)

	// initialize participation
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[participation.ModuleName], &nts.ParticipationState))
	nts.ParticipationState = populateParticipation(r, nts.ParticipationState)
	updateGenesisConfigState(participation.ModuleName, &nts.ParticipationState)

	// initialize profile
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[profile.ModuleName], &nts.ProfileState))
	nts.ProfileState = populateProfile(r, nts.ProfileState)
	updateGenesisConfigState(profile.ModuleName, &nts.ProfileState)

	// initialize reward
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[reward.ModuleName], &nts.RewardState))
	nts.RewardState = populateReward(nts.RewardState)
	updateGenesisConfigState(reward.ModuleName, &nts.RewardState)

	nts.Network = network.New(nts.T(), cfg)
}

func populateLaunch(r *rand.Rand, launchState launch.GenesisState) launch.GenesisState {
	// add chains
	for i := 0; i < 5; i++ {
		chain := sample.Chain(r, uint64(i), uint64(i))
		launchState.ChainList = append(
			launchState.ChainList,
			chain,
		)
	}

	// add genesis accounts
	for i := 0; i < 5; i++ {
		launchState.GenesisAccountList = append(
			launchState.GenesisAccountList,
			sample.GenesisAccount(r, 0, sample.Address(r)),
		)
	}

	// add vesting accounts
	for i := 0; i < 5; i++ {
		launchState.VestingAccountList = append(
			launchState.VestingAccountList,
			sample.VestingAccount(r, 0, sample.Address(r)),
		)
	}

	// add genesis validators
	for i := 0; i < 5; i++ {
		launchState.GenesisValidatorList = append(
			launchState.GenesisValidatorList,
			sample.GenesisValidator(r, uint64(0), sample.Address(r)),
		)
	}

	// add param chagne
	//for i := 0; i < 5; i++ {
	//	launchState.ParamChange = append(
	//		launchState.ParamChanges,
	//		sample.ParamChange(r, uint64(0)),
	//	)
	//}

	// add request
	for i := 0; i < 5; i++ {
		request := sample.Request(r, 0, sample.Address(r))
		request.RequestID = uint64(i)
		launchState.RequestList = append(
			launchState.RequestList,
			request,
		)
	}

	return launchState
}

func populateProject(r *rand.Rand, projectState project.GenesisState) project.GenesisState {
	// add projects
	for i := 0; i < 5; i++ {
		prjt := project.Project{
			ProjectID: uint64(i),
		}
		nullify.Fill(&prjt)
		projectState.ProjectList = append(projectState.ProjectList, prjt)
	}

	// add project chains
	for i := 0; i < 5; i++ {
		projectState.ProjectChainsList = append(projectState.ProjectChainsList, project.ProjectChains{
			ProjectID: uint64(i),
			Chains:    []uint64{uint64(i)},
		})
	}

	// add mainnet accounts
	projectID := uint64(5)
	for i := 0; i < 5; i++ {
		projectState.MainnetAccountList = append(
			projectState.MainnetAccountList,
			sample.MainnetAccount(r, projectID, sample.Address(r)),
		)
	}

	return projectState
}

func populateClaim(r *rand.Rand, claimState claim.GenesisState) claim.GenesisState {
	claimState.AirdropSupply.Supply = sample.Coin(r)
	totalSupply := sdkmath.ZeroInt()
	for i := 0; i < 5; i++ {
		// fill claim records
		accSupply := sdkmath.NewIntFromUint64(r.Uint64() % 1000)
		claimRecord := claim.ClaimRecord{
			Claimable: accSupply,
			Address:   sample.Address(r),
		}
		totalSupply = totalSupply.Add(accSupply)
		nullify.Fill(&claimRecord)
		claimState.ClaimRecordList = append(claimState.ClaimRecordList, claimRecord)
	}
	claimState.AirdropSupply.Supply.Amount = totalSupply

	// add missions
	for i := 0; i < 5; i++ {
		mission := claim.Mission{
			MissionID: uint64(i),
			Weight:    sdkmath.LegacyNewDec(r.Int63()),
		}
		nullify.Fill(&mission)
		claimState.MissionList = append(claimState.MissionList, mission)
	}

	return claimState
}

func populateMonitoringc(monitoringcState monitoringc.GenesisState) monitoringc.GenesisState {
	// add launch ID from channel ID
	for i := 0; i < 5; i++ {
		launchIDFromChannelID := monitoringc.LaunchIDFromChannelID{
			ChannelID: strconv.Itoa(i),
		}
		nullify.Fill(&launchIDFromChannelID)
		monitoringcState.LaunchIDFromChannelIDList = append(
			monitoringcState.LaunchIDFromChannelIDList,
			launchIDFromChannelID,
		)
	}

	// add monitoring history
	for i := 0; i < 5; i++ {
		monitoringHistory := monitoringc.MonitoringHistory{
			LaunchID: uint64(i),
		}
		nullify.Fill(&monitoringHistory)
		monitoringcState.MonitoringHistoryList = append(monitoringcState.MonitoringHistoryList, monitoringHistory)
	}

	// add provider client ID
	for i := 0; i < 5; i++ {
		providerClientID := monitoringc.ProviderClientID{
			LaunchID: uint64(i),
		}
		nullify.Fill(&providerClientID)
		monitoringcState.ProviderClientIDList = append(monitoringcState.ProviderClientIDList, providerClientID)
	}

	// add verified client IDs
	for i := 0; i < 5; i++ {
		verifiedClientID := monitoringc.VerifiedClientID{
			LaunchID: uint64(i),
		}
		nullify.Fill(&verifiedClientID)
		monitoringcState.VerifiedClientIDList = append(monitoringcState.VerifiedClientIDList, verifiedClientID)
	}

	return monitoringcState
}

func populateParticipation(r *rand.Rand, participationState participation.GenesisState) participation.GenesisState {
	// add used allocations
	for i := 0; i < 5; i++ {
		usedAllocations := participation.UsedAllocations{
			Address:        sample.Address(r),
			NumAllocations: sample.Int(r),
		}
		nullify.Fill(&usedAllocations)
		participationState.UsedAllocationsList = append(participationState.UsedAllocationsList, usedAllocations)
	}

	// add auction used allocations
	address := sample.Address(r)
	for i := 0; i < 5; i++ {
		auctionUsedAllocations := participation.AuctionUsedAllocations{
			Address:        address,
			AuctionID:      uint64(i),
			NumAllocations: sample.Int(r),
		}
		nullify.Fill(&auctionUsedAllocations)
		participationState.AuctionUsedAllocationsList = append(participationState.AuctionUsedAllocationsList, auctionUsedAllocations)
	}

	return participationState
}

func populateProfile(r *rand.Rand, profileState profile.GenesisState) profile.GenesisState {
	// add coordinators
	for i := 0; i < 5; i++ {
		profileState.CoordinatorList = append(
			profileState.CoordinatorList,
			profile.Coordinator{CoordinatorID: uint64(i)},
		)
	}

	// add coordinator by address
	for i := 0; i < 5; i++ {
		profileState.CoordinatorsByAddress = append(
			profileState.CoordinatorsByAddress,
			profile.CoordinatorByAddress{Address: sample.Address(r)},
		)
	}

	// add validator
	for i := 0; i < 5; i++ {
		profileState.ValidatorList = append(profileState.ValidatorList, profile.Validator{
			Address: sample.Address(r),
		})
	}

	return profileState
}

func populateReward(rewardState reward.GenesisState) reward.GenesisState {
	// add reward pool
	for i := 0; i < 5; i++ {
		rewardPool := reward.RewardPool{
			LaunchID: uint64(i),
		}
		nullify.Fill(&rewardPool)
		rewardState.RewardPoolList = append(rewardState.RewardPoolList, rewardPool)
	}

	return rewardState
}
