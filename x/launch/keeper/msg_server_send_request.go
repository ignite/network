package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func (k msgServer) SendRequest(ctx context.Context, msg *types.MsgSendRequest) (*types.MsgSendRequestResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	creator, err := k.addressCodec.StringToBytes(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid creator address %s", err.Error())
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get launch params")
	}

	chain, err := k.GetChain(ctx, msg.LaunchID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchID)
	}

	// check if request is valid for mainnet
	err = msg.Content.IsValidForMainnet()
	if err != nil && chain.IsMainnet {
		return nil, sdkerrors.Wrap(types.ErrInvalidRequestForMainnet, err.Error())
	}

	// no request can be sent if the launch of the chain is triggered
	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	coordinator, err := k.profileKeeper.GetCoordinator(ctx, chain.CoordinatorID)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %d coordinator %d not found", chain.LaunchID, chain.CoordinatorID)
	}

	// only chain with active coordinator can receive a request
	if !coordinator.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordinatorInactive,
			"the chain %d coordinator is inactive", chain.LaunchID)
	}

	// create the request from the content
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	request := types.Request{
		LaunchID:  msg.LaunchID,
		Creator:   msg.Creator,
		CreatedAt: sdkCtx.BlockTime().Unix(),
		Content:   msg.Content,
		Status:    types.Request_PENDING,
	}

	approved := false
	if msg.Creator == coordinator.Address {
		err := ApplyRequest(ctx, k.Keeper, chain, request, coordinator)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrRequestApplicationFailure, err.Error())
		}
		approved = true
		request.Status = types.Request_APPROVED
	}

	// deduct request fee if set

	if !params.RequestFee.Empty() {
		if err = k.distributionKeeper.FundCommunityPool(ctx, params.RequestFee, creator); err != nil {
			return nil, sdkerrors.Wrap(types.ErrFundCommunityPool, err.Error())
		}
	}

	request.RequestID, err = k.AppendRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	if err = sdkCtx.EventManager().EmitTypedEvent(&types.EventRequestCreated{
		Creator: msg.Creator,
		Request: request,
	}); err != nil {
		return nil, err
	}

	// call request created hook
	err = k.RequestCreated(ctx, msg.Creator, msg.LaunchID, request.RequestID, msg.Content)
	return &types.MsgSendRequestResponse{
		RequestID:    request.RequestID,
		AutoApproved: approved,
	}, err
}
