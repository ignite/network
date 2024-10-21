package monitoringp

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/monitoringp/keeper"
	"github.com/ignite/network/x/monitoringp/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genState types.GenesisState) error {
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		return err
	}

	// Set if defined
	if genState.ConsumerClientId != nil {
		if err := k.ConsumerClientID.Set(ctx, *genState.ConsumerClientId); err != nil {
			return err
		}
	}
	// Set if defined
	if genState.ConnectionChannelId != nil {
		if err := k.ConnectionChannelID.Set(ctx, *genState.ConnectionChannelId); err != nil {
			return err
		}
	}
	// Set if defined
	if genState.MonitoringInfo != nil {
		if err := k.MonitoringInfo.Set(ctx, *genState.MonitoringInfo); err != nil {
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

	// initialize and setup the consumer IBC client
	if genState.Params.ConsumerConsensusState.Timestamp != "" {
		_, err := k.InitializeConsumerClient(ctx)
		if err != nil {
			return errors.Wrap(err, "couldn't initialize the consumer client ID")
		}
	}

	return nil
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
	// Get all consumerClientID
	consumerClientID, err := k.ConsumerClientID.Get(ctx)
	if err == nil {
		genesis.ConsumerClientId = &consumerClientID
	}
	// Get all connectionChannelID
	connectionChannelID, err := k.ConnectionChannelID.Get(ctx)
	if err == nil {
		genesis.ConnectionChannelId = &connectionChannelID
	}
	// Get all monitoringInfo
	monitoringInfo, err := k.MonitoringInfo.Get(ctx)
	if err == nil {
		genesis.MonitoringInfo = &monitoringInfo
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis, nil
}
