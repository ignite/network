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
		return nil, sdkerrors.Wrap(err, "invalid participant address")
	}

	availableAlloc, err := k.GetAvailableAllocations(ctx, msg.Participant)
	if err != nil {
		return nil, err
	}

	// check if auction exists
	auction, err := k.fundraisingKeeper.GetAuction(ctx, msg.AuctionID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.AuctionID)
	}

	// check if auction is not in standby
	if auction.GetStatus() != fundraisingtypes.AuctionStatusStandBy {
		return nil, sdkerrors.Wrapf(types.ErrParticipationNotAllowed, "auction %d is not in standby", msg.AuctionID)
	}

	// check if auction allows participation at this time
	isRegistrationEnabled, err := k.IsRegistrationEnabled(ctx, auction.GetStartTime())
	if err != nil {
		return nil, ignterrors.Criticalf("failed to get project params %s", err.Error())
	}
	if !isRegistrationEnabled {
		return nil, sdkerrors.Wrapf(types.ErrParticipationNotAllowed, "participation period for auction %d not yet started", msg.AuctionID)
	}

	// check if the user is already added as an allowed bidder for the auction
	_, err = k.AuctionUsedAllocations.Get(ctx, collections.Join(sdk.AccAddress(participantAddress), msg.AuctionID))
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyParticipating,
			"address %s is already a participant for auction %d",
			msg.Participant, msg.AuctionID)
	}

	tier, found := types.GetTierFromID(params.ParticipationTierList, msg.TierID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTierNotFound, "tier %d not found", msg.TierID)
	}

	// check if user has enough available allocations to cover tier
	if tier.RequiredAllocations.GT(availableAlloc) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientAllocations,
			"available allocations %s is less than required allocations %s for tier %d",
			availableAlloc.String(), tier.RequiredAllocations.String(), tier.TierID)
	}

	allowedBidder := fundraisingtypes.AllowedBidder{
		Bidder:       msg.Participant,
		MaxBidAmount: tier.Benefits.MaxBidAmount,
	}
	if err := k.fundraisingKeeper.AddAllowedBidders(
		ctx, msg.AuctionID,
		[]fundraisingtypes.AllowedBidder{allowedBidder},
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBidder, err.Error())
	}

	// set used allocations
	numUsedAllocations := sdkmath.ZeroInt()
	used, err := k.UsedAllocations.Get(ctx, msg.Participant)
	if errors.Is(err, collections.ErrNotFound) {
		numUsedAllocations = used.NumAllocations
	} else if err != nil {
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
	err = k.AuctionUsedAllocations.Set(ctx, collections.Join(sdk.AccAddress(participantAddress), msg.AuctionID), types.AuctionUsedAllocations{
		Address:        msg.Participant,
		AuctionID:      msg.AuctionID,
		NumAllocations: tier.RequiredAllocations,
		Withdrawn:      false,
	})
	if err != nil {
		return nil, ignterrors.Criticalf("failed to set auction used allocations %s", err.Error())
	}

	return &types.MsgParticipateResponse{}, sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvent(&types.EventAllocationsUsed{
		Participant:    msg.Participant,
		AuctionID:      msg.AuctionID,
		NumAllocations: tier.RequiredAllocations,
	})
}
