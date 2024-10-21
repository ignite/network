package keeper

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	"github.com/cometbft/cometbft/light"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	committypes "github.com/cosmos/ibc-go/v8/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"

	"github.com/ignite/network/pkg/chainid"
	"github.com/ignite/network/x/monitoringp/types"
)

// InitializeConsumerClient initializes the consumer IBC client and set it in the store
func (k Keeper) InitializeConsumerClient(ctx sdk.Context) (string, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return "", err
	}

	// initialize the client state
	clientState, err := k.initializeClientState(ctx, params.ConsumerChainId)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidClientState, err.Error())
	}
	if err := clientState.Validate(); err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidClientState, err.Error())
	}

	// get consensus state from param
	tmConsensusState, err := params.ConsumerConsensusState.ToTendermintConsensusState()
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidConsensusState, err.Error())
	}

	// create IBC client for consumer
	clientID, err := k.clientKeeper.CreateClient(ctx, clientState, &tmConsensusState)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrClientCreationFailure, err.Error())
	}

	// register the IBC client
	err = k.ConsumerClientID.Set(ctx, types.ConsumerClientID{
		ClientId: clientID,
	})
	return clientID, err
}

// initializeClientState initializes the client state provided for the IBC client
func (k Keeper) initializeClientState(ctx sdk.Context, chainID string) (*ibctmtypes.ClientState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	_, revisionNumber, err := chainid.ParseGenesisChainID(chainID)
	if err != nil {
		return nil, err
	}

	unbondingPeriod := params.ConsumerUnbondingPeriod
	revisionHeight := params.ConsumerRevisionHeight

	return ibctmtypes.NewClientState(
		chainID,
		ibctmtypes.NewFractionFromTm(light.DefaultTrustLevel),
		time.Second*time.Duration(unbondingPeriod)-1,
		time.Second*time.Duration(unbondingPeriod),
		time.Minute*10,
		clienttypes.NewHeight(revisionNumber, revisionHeight),
		committypes.GetSDKSpecs(),
		[]string{"upgrade", "upgradedIBCState"},
	), nil
}
