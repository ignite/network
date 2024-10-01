package monitoringp

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/monitoringp/keeper"
	"github.com/ignite/network/x/monitoringp/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) error {
	// Set if defined
	if genState.MonitoringInfo != nil {
		if err := k.MonitoringInfo.Set(ctx, *genState.MonitoringInfo); err != nil {
			return err
		}
	}
	// Set if defined
	if genState.ConnectionChannelID != nil {
		if err := k.ConnectionChannelID.Set(ctx, *genState.ConnectionChannelID); err != nil {
			return err
		}
	}
	// Set if defined
	if genState.ConsumerClientID != nil {
		if err := k.ConsumerClientID.Set(ctx, *genState.ConsumerClientID); err != nil {
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
	// Get all monitoringInfo
	monitoringInfo, err := k.MonitoringInfo.Get(ctx)
	if err == nil {
		genesis.MonitoringInfo = &monitoringInfo
	}
	// Get all connectionChannelID
	connectionChannelID, err := k.ConnectionChannelID.Get(ctx)
	if err == nil {
		genesis.ConnectionChannelID = &connectionChannelID
	}
	// Get all consumerClientID
	consumerClientID, err := k.ConsumerClientID.Get(ctx)
	if err == nil {
		genesis.ConsumerClientID = &consumerClientID
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
