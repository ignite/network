package app

import (
	"time"

	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	authzmodulev1 "cosmossdk.io/api/cosmos/authz/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	circuitmodulev1 "cosmossdk.io/api/cosmos/circuit/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	crisismodulev1 "cosmossdk.io/api/cosmos/crisis/module/v1"
	distrmodulev1 "cosmossdk.io/api/cosmos/distribution/module/v1"
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	feegrantmodulev1 "cosmossdk.io/api/cosmos/feegrant/module/v1"
	genutilmodulev1 "cosmossdk.io/api/cosmos/genutil/module/v1"
	govmodulev1 "cosmossdk.io/api/cosmos/gov/module/v1"
	groupmodulev1 "cosmossdk.io/api/cosmos/group/module/v1"
	mintmodulev1 "cosmossdk.io/api/cosmos/mint/module/v1"
	nftmodulev1 "cosmossdk.io/api/cosmos/nft/module/v1"
	paramsmodulev1 "cosmossdk.io/api/cosmos/params/module/v1"
	slashingmodulev1 "cosmossdk.io/api/cosmos/slashing/module/v1"
	stakingmodulev1 "cosmossdk.io/api/cosmos/staking/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	upgrademodulev1 "cosmossdk.io/api/cosmos/upgrade/module/v1"
	vestingmodulev1 "cosmossdk.io/api/cosmos/vesting/module/v1"
	"cosmossdk.io/core/appconfig"
	circuittypes "cosmossdk.io/x/circuit/types"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	"cosmossdk.io/x/nft"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	fundraisingmodulev1 "github.com/ignite/modules/api/modules/fundraising/module/v1"
	claimtypes "github.com/ignite/modules/x/claim/types"
	_ "github.com/ignite/modules/x/fundraising/module"
	fundraisingmoduletypes "github.com/ignite/modules/x/fundraising/types"
	"google.golang.org/protobuf/types/known/durationpb"

	launchmodulev1 "github.com/ignite/network/api/network/launch/module/v1"
	monitoringcmodulev1 "github.com/ignite/network/api/network/monitoringc/module/v1"
	monitoringpmodulev1 "github.com/ignite/network/api/network/monitoringp/module/v1"
	participationmodulev1 "github.com/ignite/network/api/network/participation/module/v1"
	profilemodulev1 "github.com/ignite/network/api/network/profile/module/v1"
	projectmodulev1 "github.com/ignite/network/api/network/project/module/v1"
	rewardmodulev1 "github.com/ignite/network/api/network/reward/module/v1"
	networktypes "github.com/ignite/network/pkg/types"
	_ "github.com/ignite/network/x/launch/module"
	launchmoduletypes "github.com/ignite/network/x/launch/types"
	_ "github.com/ignite/network/x/monitoringc/module"
	monitoringcmoduletypes "github.com/ignite/network/x/monitoringc/types"
	_ "github.com/ignite/network/x/monitoringp/module"
	monitoringpmoduletypes "github.com/ignite/network/x/monitoringp/types"
	_ "github.com/ignite/network/x/participation/module"
	participationmoduletypes "github.com/ignite/network/x/participation/types"
	_ "github.com/ignite/network/x/profile/module"
	profilemoduletypes "github.com/ignite/network/x/profile/types"
	_ "github.com/ignite/network/x/project/module"
	projectmoduletypes "github.com/ignite/network/x/project/types"
	_ "github.com/ignite/network/x/reward/module"
	rewardmoduletypes "github.com/ignite/network/x/reward/types"
)

var (
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	genesisModuleOrder = []string{
		// cosmos-sdk/ibc modules
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibcexported.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		nft.ModuleName,
		group.ModuleName,
		consensustypes.ModuleName,
		circuittypes.ModuleName,
		// chain modules
		fundraisingmoduletypes.ModuleName,
		profilemoduletypes.ModuleName,
		launchmoduletypes.ModuleName,
		rewardmoduletypes.ModuleName,
		projectmoduletypes.ModuleName,
		participationmoduletypes.ModuleName,
		monitoringcmoduletypes.ModuleName,
		monitoringpmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/initGenesis
	}

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	beginBlockers = []string{
		// cosmos sdk modules
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authz.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		capabilitytypes.ModuleName,
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		fundraisingmoduletypes.ModuleName,
		profilemoduletypes.ModuleName,
		launchmoduletypes.ModuleName,
		rewardmoduletypes.ModuleName,
		projectmoduletypes.ModuleName,
		participationmoduletypes.ModuleName,
		monitoringcmoduletypes.ModuleName,
		monitoringpmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/beginBlockers
	}

	endBlockers = []string{
		// cosmos sdk modules
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		feegrant.ModuleName,
		group.ModuleName,
		genutiltypes.ModuleName,
		// ibc modules
		ibcexported.ModuleName,
		ibctransfertypes.ModuleName,
		capabilitytypes.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// chain modules
		fundraisingmoduletypes.ModuleName,
		profilemoduletypes.ModuleName,
		launchmoduletypes.ModuleName,
		rewardmoduletypes.ModuleName,
		projectmoduletypes.ModuleName,
		participationmoduletypes.ModuleName,
		monitoringcmoduletypes.ModuleName,
		monitoringpmoduletypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/endBlockers
	}

	preBlockers = []string{
		upgradetypes.ModuleName,
		// this line is used by starport scaffolding # stargate/app/preBlockers
	}

	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: distrtypes.ModuleName},
		{Account: minttypes.ModuleName, Permissions: []string{authtypes.Minter}},
		{Account: stakingtypes.BondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: stakingtypes.NotBondedPoolName, Permissions: []string{authtypes.Burner, stakingtypes.ModuleName}},
		{Account: govtypes.ModuleName, Permissions: []string{authtypes.Burner}},
		{Account: nft.ModuleName},
		{Account: ibctransfertypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: ibcfeetypes.ModuleName},
		{Account: icatypes.ModuleName},
		{Account: projectmoduletypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: rewardmoduletypes.ModuleName},
		{Account: fundraisingmoduletypes.ModuleName},
		{Account: monitoringcmoduletypes.ModuleName},
		{Account: monitoringpmoduletypes.ModuleName},
		{Account: claimtypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		// this line is used by starport scaffolding # stargate/app/maccPerms
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		distrtypes.ModuleName,
		minttypes.ModuleName,
		stakingtypes.BondedPoolName,
		stakingtypes.NotBondedPoolName,
		nft.ModuleName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// appConfig application configuration (used by depinject)
	appConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: runtime.ModuleName,
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName:       networktypes.Name,
					PreBlockers:   preBlockers,
					BeginBlockers: beginBlockers,
					EndBlockers:   endBlockers,
					InitGenesis:   genesisModuleOrder,
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: genesisModuleOrder,
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: nil,
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             networktypes.AccountAddressPrefix,
					ModuleAccountPermissions: moduleAccPerms,
					// By default modules authority is the governance module. This is configurable with the following:
					// Authority: "group", // A custom module authority can be set using a module name
					// Authority: "cosmos1cwwv22j5ca08ggdv9c2uky355k908694z577tv", // or a specific address
				}),
			},
			{
				Name:   nft.ModuleName,
				Config: appconfig.WrapAny(&nftmodulev1.Module{}),
			},
			{
				Name:   vestingtypes.ModuleName,
				Config: appconfig.WrapAny(&vestingmodulev1.Module{}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name: stakingtypes.ModuleName,
				Config: appconfig.WrapAny(&stakingmodulev1.Module{
					// NOTE: specifying a prefix is only necessary when using bech32 addresses
					// If not specfied, the auth Bech32Prefix appended with "valoper" and "valcons" is used by default
					Bech32PrefixValidator: networktypes.AccountAddressPrefix + "valoper",
					Bech32PrefixConsensus: networktypes.AccountAddressPrefix + "valcons",
				}),
			},
			{
				Name:   slashingtypes.ModuleName,
				Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
			},
			{
				Name:   paramstypes.ModuleName,
				Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   authz.ModuleName,
				Config: appconfig.WrapAny(&authzmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   distrtypes.ModuleName,
				Config: appconfig.WrapAny(&distrmodulev1.Module{}),
			},
			{
				Name:   evidencetypes.ModuleName,
				Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
			},
			{
				Name:   minttypes.ModuleName,
				Config: appconfig.WrapAny(&mintmodulev1.Module{}),
			},
			{
				Name: group.ModuleName,
				Config: appconfig.WrapAny(&groupmodulev1.Module{
					MaxExecutionPeriod: durationpb.New(time.Second * 1209600),
					MaxMetadataLen:     255,
				}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   govtypes.ModuleName,
				Config: appconfig.WrapAny(&govmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   circuittypes.ModuleName,
				Config: appconfig.WrapAny(&circuitmodulev1.Module{}),
			},
			{
				Name:   fundraisingmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&fundraisingmodulev1.Module{}),
			},
			{
				Name:   profilemoduletypes.ModuleName,
				Config: appconfig.WrapAny(&profilemodulev1.Module{}),
			},
			{
				Name:   launchmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&launchmodulev1.Module{}),
			},
			{
				Name:   rewardmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&rewardmodulev1.Module{}),
			},
			{
				Name:   projectmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&projectmodulev1.Module{}),
			},
			{
				Name:   participationmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&participationmodulev1.Module{}),
			},
			{
				Name:   monitoringcmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&monitoringcmodulev1.Module{}),
			},
			{
				Name:   monitoringpmoduletypes.ModuleName,
				Config: appconfig.WrapAny(&monitoringpmodulev1.Module{}),
			},
			// this line is used by starport scaffolding # stargate/app/moduleConfig
		},
	})
)
