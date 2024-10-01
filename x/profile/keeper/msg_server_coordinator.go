package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/profile/types"
)

func (k msgServer) CreateCoordinator(ctx context.Context, msg *types.MsgCreateCoordinator) (*types.MsgCreateCoordinatorResponse, error) {
	address, err := k.addressCodec.StringToBytes(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid address")
	}

	if coordinator, err := k.CoordinatorByAddress.Get(ctx, address); err != nil {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAlreadyExist, "coordinatorID: %d", coordinator.CoordinatorID)
	}

	nextID, err := k.CoordinatorSeq.Next(ctx)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "failed to get next id")
	}

	coordinator := types.Coordinator{
		CoordinatorID: nextID,
		Address:       msg.Address,
		Description:   msg.Description,
		Active:        true,
	}

	if err = k.Coordinator.Set(ctx, nextID, coordinator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator")
	}

	if err = k.CoordinatorByAddress.Set(ctx, address, types.CoordinatorByAddress{
		Address:       msg.Address,
		CoordinatorID: nextID,
	}); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator by address")
	}

	return &types.MsgCreateCoordinatorResponse{CoordinatorID: nextID}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
		&types.EventCoordinatorCreated{
			CoordinatorID: nextID,
			Address:       msg.Address,
		})
}

func (k msgServer) UpdateCoordinatorDescription(ctx context.Context, msg *types.MsgUpdateCoordinatorDescription) (*types.MsgUpdateCoordinatorDescriptionResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	address, err := k.addressCodec.StringToBytes(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid address")
	}

	coordByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if !errors.IsOf(err, collections.ErrNotFound) {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAddressNotFound, "coordinator address %s not found", msg.Address)
	}

	coordinator, err := k.Coordinator.Get(ctx, coordByAddress.CoordinatorID)
	if err != nil {
		if errors.IsOf(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", 10)) // TODO
		}

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

	if err := k.Coordinator.Set(ctx, coordByAddress.CoordinatorID, coordinator); err != nil {
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
		return nil, errorsmod.Wrap(err, "invalid address")
	}
	coordByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if !errors.IsOf(err, collections.ErrNotFound) {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAddressNotFound, "coordinator address %s not found", msg.Address)
	}

	newAddress, err := k.addressCodec.StringToBytes(msg.NewAddress)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid new address")
	}
	if newCoord, err := k.CoordinatorByAddress.Get(ctx, newAddress); err != nil {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAlreadyExist, "new address already have a coordinator: %d", newCoord.CoordinatorID)
	}

	coordinator, err := k.Coordinator.Get(ctx, coordByAddress.CoordinatorID)
	if err != nil {
		if errors.IsOf(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", 10)) // TODO
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to get coordinator")
	}

	// Check if the coordinator is inactive
	if !coordinator.Active {
		return nil, errors.Criticalf("inactive coordinator address should not exist in store, ID: %d", coordByAddress.CoordinatorID)
	}

	coordinator.Address = msg.NewAddress

	// Remove the old coordinator by addrless and create a new one
	if err = k.CoordinatorByAddress.Remove(ctx, address); err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrLogic, "failed to remove coordinator by address %s", msg.Address)
	}
	if err = k.CoordinatorByAddress.Set(ctx, newAddress, types.CoordinatorByAddress{
		Address:       msg.NewAddress,
		CoordinatorID: coordByAddress.CoordinatorID,
	}); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator by address")
	}
	if err := k.Coordinator.Set(ctx, coordByAddress.CoordinatorID, coordinator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update coordinator address")
	}

	return &types.MsgUpdateCoordinatorAddressResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
		&types.EventCoordinatorAddressUpdated{
			CoordinatorID: coordByAddress.CoordinatorID,
			NewAddress:    msg.NewAddress,
		})
}

func (k msgServer) DisableCoordinator(ctx context.Context, msg *types.MsgDisableCoordinator) (*types.MsgDisableCoordinatorResponse, error) {
	address, err := k.addressCodec.StringToBytes(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid address")
	}

	coordByAddress, err := k.CoordinatorByAddress.Get(ctx, address)
	if !errors.IsOf(err, collections.ErrNotFound) {
		return nil, errorsmod.Wrapf(types.ErrCoordinatorAddressNotFound, "coordinator address %s not found", msg.Address)
	}

	// Checks that the element exists
	coordinator, err := k.Coordinator.Get(ctx, coordByAddress.CoordinatorID)
	if err != nil {
		if errors.IsOf(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", coordByAddress.CoordinatorID))
		}

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
				coordByAddress.CoordinatorID)
	}

	// disable by setting to inactive and remove CoordByAddress
	coordinator.Active = false
	if err := k.Coordinator.Set(ctx, coordinator.CoordinatorID, coordinator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to set coordinator")
	}

	if err := k.CoordinatorByAddress.Remove(ctx, address); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to delete coordinator")
	}

	return &types.MsgDisableCoordinatorResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(
		&types.EventCoordinatorDisabled{
			CoordinatorID: coordByAddress.CoordinatorID,
			Address:       msg.Address,
		})
}
