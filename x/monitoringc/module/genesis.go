package monitoringc

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/monitoringc/keeper"
	"github.com/ignite/network/x/monitoringc/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genState types.GenesisState) error {
	// Set all the verifiedClientID
	for _, elem := range genState.VerifiedClientIdList {
		if err := k.VerifiedClientID.Set(ctx, elem.LaunchId, elem); err != nil {
			return err
		}
	}

	// Set all the providerClientID
	for _, elem := range genState.ProviderClientIdList {
		if err := k.ProviderClientID.Set(ctx, elem.LaunchId, elem); err != nil {
			return err
		}
	}

	// Set all the launchIDFromVerifiedClientID
	for _, elem := range genState.LaunchIdFromVerifiedClientIdList {
		if err := k.LaunchIDFromVerifiedClientID.Set(ctx, elem.ClientId, elem); err != nil {
			return err
		}
	}

	// Set all the launchIDFromChannelID
	for _, elem := range genState.LaunchIdFromChannelIdList {
		if err := k.LaunchIDFromChannelID.Set(ctx, elem.ChannelId, elem); err != nil {
			return err
		}
	}

	// Set all the monitoringHistory
	for _, elem := range genState.MonitoringHistoryList {
		if err := k.MonitoringHistory.Set(ctx, elem.LaunchId, elem); err != nil {
			return err
		}
	}

	// this line is used by starport scaffolding # genesis/module/init

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			return errors.Wrap(err, "could not claim port capability")
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	genesis.PortId = k.GetPort(ctx)

	if err := k.VerifiedClientID.Walk(ctx, nil, func(_ uint64, val types.VerifiedClientID) (stop bool, err error) {
		genesis.VerifiedClientIdList = append(genesis.VerifiedClientIdList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.ProviderClientID.Walk(ctx, nil, func(_ uint64, val types.ProviderClientID) (stop bool, err error) {
		genesis.ProviderClientIdList = append(genesis.ProviderClientIdList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.LaunchIDFromVerifiedClientID.Walk(ctx, nil, func(_ string, val types.LaunchIDFromVerifiedClientID) (stop bool, err error) {
		genesis.LaunchIdFromVerifiedClientIdList = append(genesis.LaunchIdFromVerifiedClientIdList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.LaunchIDFromChannelID.Walk(ctx, nil, func(_ string, val types.LaunchIDFromChannelID) (stop bool, err error) {
		genesis.LaunchIdFromChannelIdList = append(genesis.LaunchIdFromChannelIdList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.MonitoringHistory.Walk(ctx, nil, func(_ uint64, val types.MonitoringHistory) (stop bool, err error) {
		genesis.MonitoringHistoryList = append(genesis.MonitoringHistoryList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
