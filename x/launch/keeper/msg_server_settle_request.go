package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

func (k msgServer) SettleRequest(ctx context.Context, msg *types.MsgSettleRequest) (*types.MsgSettleRequestResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Signer); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid signer address %s", err.Error())
	}

	chain, err := k.GetChain(ctx, msg.LaunchId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.LaunchId)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchId)
	}

	coordinator, err := k.profileKeeper.GetCoordinator(ctx, chain.CoordinatorId)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrChainInactive,
			"the chain %d coordinator %d not found", chain.LaunchId, chain.CoordinatorId)
	}

	if !coordinator.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordinatorInactive,
			"the chain %d coordinator inactive", chain.LaunchId)
	}

	if msg.Approve && msg.Signer != coordinator.Address {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Signer)
	}

	// first check if the request exists
	request, err := k.Request.Get(ctx, collections.Join(msg.LaunchId, msg.RequestId))
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrRequestNotFound, "failed to get request %d for chain %d", msg.RequestId, msg.LaunchId)
	}

	if request.Status != types.Request_PENDING {
		return nil, sdkerrors.Wrapf(types.ErrRequestSettled,
			"request %d is not pending",
			msg.RequestId,
		)
	}

	if msg.Signer != request.Creator && msg.Signer != coordinator.Address {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Signer)
	}

	// apply request if approving and update status
	if msg.Approve {
		err := ApplyRequest(ctx, k.Keeper, chain, request, coordinator)
		if err != nil {
			return nil, err
		}
		request.Status = types.Request_APPROVED
	} else {
		request.Status = types.Request_REJECTED
	}

	if err := k.Request.Set(ctx, collections.Join(request.LaunchId, request.RequestId), request); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidRequestContent, "failed to set request %d for launch %d", request.RequestId, request.LaunchId)
	}

	err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventRequestSettled{
		LaunchId:  msg.LaunchId,
		RequestId: request.RequestId,
		Approved:  msg.Approve,
	})

	return &types.MsgSettleRequestResponse{}, err
}
