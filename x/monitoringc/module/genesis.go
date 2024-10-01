package monitoringc

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/monitoringc/keeper"
	"github.com/ignite/network/x/monitoringc/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) error {
	// Set all the launchIDFromChannelID
	for _, elem := range genState.LaunchIDFromChannelIDList {
		if err := k.LaunchIDFromChannelID.Set(ctx, elem.ChannelID, elem); err != nil {
			return err
		}
	}
	// Set all the launchIDFromVerifiedClientID
	for _, elem := range genState.LaunchIDFromVerifiedClientIDList {
		if err := k.LaunchIDFromVerifiedClientID.Set(ctx, elem.ClientID, elem); err != nil {
			return err
		}
	}
	// Set all the monitoringHistory
	for _, elem := range genState.MonitoringHistoryList {
		if err := k.MonitoringHistory.Set(ctx, elem.LaunchID, elem); err != nil {
			return err
		}
	}

	// Set all the verifiedClientID
	for _, elem := range genState.VerifiedClientIDList {
		if err := k.VerifiedClientID.Set(ctx, elem.LaunchID, elem); err != nil {
			return err
		}
	}
	// Set all the providerClientID
	for _, elem := range genState.ProviderClientIDList {
		if err := k.ProviderClientID.Set(ctx, elem.LaunchID, elem); err != nil {
			return err
		}
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortID)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortID) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortID)
		if err != nil {
			return errors.Wrap(err, "could not claim port capability")
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	genesis.PortID = k.GetPort(ctx)
	if err := k.LaunchIDFromChannelID.Walk(ctx, nil, func(_ string, val types.LaunchIDFromChannelID) (stop bool, err error) {
		genesis.LaunchIDFromChannelIDList = append(genesis.LaunchIDFromChannelIDList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.LaunchIDFromVerifiedClientID.Walk(ctx, nil, func(_ string, val types.LaunchIDFromVerifiedClientID) (stop bool, err error) {
		genesis.LaunchIDFromVerifiedClientIDList = append(genesis.LaunchIDFromVerifiedClientIDList, val)
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

	if err := k.VerifiedClientID.Walk(ctx, nil, func(_ uint64, val types.VerifiedClientID) (stop bool, err error) {
		genesis.VerifiedClientIDList = append(genesis.VerifiedClientIDList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	if err := k.ProviderClientID.Walk(ctx, nil, func(_ uint64, val types.ProviderClientID) (stop bool, err error) {
		genesis.ProviderClientIDList = append(genesis.ProviderClientIDList, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
