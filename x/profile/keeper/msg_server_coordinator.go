package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/profile/types"
)

func (k msgServer) CreateCoordinator(ctx context.Context, msg *types.MsgCreateCoordinator) (*types.MsgCreateCoordinatorResponse, error) {
	address, err := k.addressCodec.StringToBytes(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid address %s", err.Error())
	}

	if coordinator, err := k.CoordinatorByAddress.Get(ctx, address); err == nil {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAlreadyExist, "coordinatorID: %d", coordinator.CoordinatorId)
	}

	nextID, err := k.CoordinatorSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	coordinator := types.Coordinator{
		CoordinatorId: nextID,
		Address:       msg.Address,
		Description:   msg.Description,
		Active:        true,
	}

	if err = k.Coordinator.Set(ctx, nextID, coordinator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator")
	}

	if err = k.CoordinatorByAddress.Set(ctx, address, types.CoordinatorByAddress{
		Address:       msg.Address,
		CoordinatorId: nextID,
	}); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator by address")
	}

	return &types.MsgCreateCoordinatorResponse{CoordinatorId: nextID}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
		&types.EventCoordinatorCreated{
			CoordinatorId: nextID,
			Address:       msg.Address,
		})
}

func (k msgServer) UpdateCoordinatorDescription(ctx context.Context, msg *types.MsgUpdateCoordinatorDescription) (*types.MsgUpdateCoordinatorDescriptionResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	address, err := k.addressCodec.StringToBytes(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid address %s", err.Error())
	}

	coordByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAddressNotFound, "coordinator address %s not found", msg.Address)
	}

	coordinator, err := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get coordinator")
	}

	// Checks if the msg address is the same as the current owner
	if msg.Address != coordinator.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if len(msg.Description.Identity) > 0 {
		coordinator.Description.Identity = msg.Description.Identity
	}
	if len(msg.Description.Website) > 0 {
		coordinator.Description.Website = msg.Description.Website
	}
	if len(msg.Description.Details) > 0 {
		coordinator.Description.Details = msg.Description.Details
	}

	if err := k.Coordinator.Set(ctx, coordByAddress.CoordinatorId, coordinator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update coordinator")
	}

	return &types.MsgUpdateCoordinatorDescriptionResponse{}, nil
}

func (k msgServer) UpdateCoordinatorAddress(ctx context.Context, msg *types.MsgUpdateCoordinatorAddress) (*types.MsgUpdateCoordinatorAddressResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	address, err := k.addressCodec.StringToBytes(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid address %s", err.Error())
	}
	coordByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAddressNotFound, "coordinator address %s not found", msg.Address)
	}

	newAddress, err := k.addressCodec.StringToBytes(msg.NewAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid new address %s", err.Error())
	}
	if newCoord, err := k.CoordinatorByAddress.Get(ctx, newAddress); err == nil {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAlreadyExist, "new address already have a coordinator: %d", newCoord.CoordinatorId)
	}

	coordinator, err := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get coordinator")
	}

	// Check if the coordinator is inactive
	if !coordinator.Active {
		return nil, errors.Criticalf("inactive coordinator address should not exist in store, ID: %d", coordByAddress.CoordinatorId)
	}

	coordinator.Address = msg.NewAddress

	// Remove the old coordinator by addrless and create a new one
	if err = k.CoordinatorByAddress.Remove(ctx, address); err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrLogic, "failed to remove coordinator by address %s", msg.Address)
	}
	if err = k.CoordinatorByAddress.Set(ctx, newAddress, types.CoordinatorByAddress{
		Address:       msg.NewAddress,
		CoordinatorId: coordByAddress.CoordinatorId,
	}); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator by address")
	}
	if err := k.Coordinator.Set(ctx, coordByAddress.CoordinatorId, coordinator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update coordinator address")
	}

	return &types.MsgUpdateCoordinatorAddressResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
		&types.EventCoordinatorAddressUpdated{
			CoordinatorId: coordByAddress.CoordinatorId,
			NewAddress:    msg.NewAddress,
		})
}

func (k msgServer) DisableCoordinator(ctx context.Context, msg *types.MsgDisableCoordinator) (*types.MsgDisableCoordinatorResponse, error) {
	address, err := k.addressCodec.StringToBytes(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid address %s", err.Error())
	}

	coordByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAddressNotFound, "coordinator address %s not found", msg.Address)
	}

	// Checks that the element exists
	coordinator, err := k.GetCoordinator(ctx, coordByAddress.CoordinatorId)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get coordinator")
	}

	// Checks if the msg address is the same as the current owner
	if msg.Address != coordinator.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check if the coordinator is inactive
	if !coordinator.Active {
		return nil,
			errors.Criticalf("inactive coordinator address should not exist in store, ID: %d",
				coordByAddress.CoordinatorId)
	}

	// disable by setting to inactive and remove CoordByAddress
	coordinator.Active = false
	if err := k.Coordinator.Set(ctx, coordinator.CoordinatorId, coordinator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator")
	}

	if err := k.CoordinatorByAddress.Remove(ctx, address); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete coordinator")
	}

	return &types.MsgDisableCoordinatorResponse{
			CoordinatorId: coordinator.CoordinatorId,
		}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
			&types.EventCoordinatorDisabled{
				CoordinatorId: coordByAddress.CoordinatorId,
				Address:       msg.Address,
			})
}
