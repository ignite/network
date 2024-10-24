package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"
	"github.com/pkg/errors"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/participation/types"
)

func (k msgServer) Participate(ctx context.Context, msg *types.MsgParticipate) (*types.MsgParticipateResponse, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get launch params")
	}

	participantAddress, err := k.addressCodec.StringToBytes(msg.Participant)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSigner, "invalid participant address %s", err.Error())
	}

	availableAlloc, err := k.GetAvailableAllocations(ctx, msg.Participant)
	if err != nil {
		return nil, err
	}

	// check if auction exists
	auction, err := k.fundraisingKeeper.GetAuction(ctx, msg.AuctionId)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.AuctionId)
	}

	// check if auction is not in standby
	if auction.GetStatus() != fundraisingtypes.AuctionStatusStandBy {
		return nil, sdkerrors.Wrapf(types.ErrParticipationNotAllowed, "auction %d is not in standby", msg.AuctionId)
	}

	// check if auction allows participation at this time
	isRegistrationEnabled, err := k.IsRegistrationEnabled(ctx, auction.GetStartTime())
	if err != nil {
		return nil, ignterrors.Criticalf("failed to get project params %s", err.Error())
	}
	if !isRegistrationEnabled {
		return nil, sdkerrors.Wrapf(types.ErrParticipationNotAllowed, "participation period for auction %d not yet started", msg.AuctionId)
	}

	// check if the user is already added as an allowed bidder for the auction
	_, err = k.AuctionUsedAllocations.Get(ctx, collections.Join(sdk.AccAddress(participantAddress), msg.AuctionId))
	if err == nil {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyParticipating,
			"address %s is already a participant for auction %d",
			msg.Participant, msg.AuctionId)
	} else if !errors.Is(err, collections.ErrNotFound) {
		return nil, ignterrors.Critical(err.Error())
	}

	tier, found := types.GetTierFromID(params.ParticipationTierList, msg.TierId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTierNotFound, "tier %d not found", msg.TierId)
	}

	// check if user has enough available allocations to cover tier
	if tier.RequiredAllocations.GT(availableAlloc) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientAllocations,
			"available allocations %s is less than required allocations %s for tier %d",
			availableAlloc.String(), tier.RequiredAllocations.String(), tier.TierId)
	}

	allowedBidder := fundraisingtypes.AllowedBidder{
		Bidder:       msg.Participant,
		MaxBidAmount: tier.Benefits.MaxBidAmount,
	}
	if err := k.fundraisingKeeper.AddAllowedBidders(
		ctx, msg.AuctionId,
		[]fundraisingtypes.AllowedBidder{allowedBidder},
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBidder, err.Error())
	}

	// set used allocations
	numUsedAllocations := sdkmath.ZeroInt()
	used, err := k.UsedAllocations.Get(ctx, msg.Participant)
	if err == nil {
		numUsedAllocations = used.NumAllocations
	} else if !errors.Is(err, collections.ErrNotFound) {
		return nil, ignterrors.Criticalf("failed to get used allocations %s", err.Error())
	}

	numUsedAllocations = numUsedAllocations.Add(tier.RequiredAllocations)
	err = k.UsedAllocations.Set(ctx, msg.Participant, types.UsedAllocations{
		Address:        msg.Participant,
		NumAllocations: numUsedAllocations,
	})
	if err != nil {
		return nil, ignterrors.Criticalf("failed to set used allocations %s", err.Error())
	}

	// set auction used allocations
	err = k.AuctionUsedAllocations.Set(ctx, collections.Join(sdk.AccAddress(participantAddress), msg.AuctionId), types.AuctionUsedAllocations{
		Address:        msg.Participant,
		AuctionId:      msg.AuctionId,
		NumAllocations: tier.RequiredAllocations,
		Withdrawn:      false,
	})
	if err != nil {
		return nil, ignterrors.Criticalf("failed to set auction used allocations %s", err.Error())
	}

	return &types.MsgParticipateResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventAllocationsUsed{
		Participant:    msg.Participant,
		AuctionId:      msg.AuctionId,
		NumAllocations: tier.RequiredAllocations,
	})
}
