package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"

	ignterrors "github.com/ignite/network/pkg/errors"
	"github.com/ignite/network/x/participation/types"
)

func (k msgServer) WithdrawAllocations(ctx context.Context, msg *types.MsgWithdrawAllocations) (*types.MsgWithdrawAllocationsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockTime := sdkCtx.BlockTime()

	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, ignterrors.Critical("failed to get launch params")
	}

	participantAddress, err := k.addressCodec.StringToBytes(msg.Participant)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "invalid participant address")
	}

	auction, err := k.fundraisingKeeper.GetAuction(ctx, msg.AuctionID)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "%d", msg.AuctionID)
	}

	// only prevent time-based restrictions on withdrawals if the auction's status is not `CANCELLED`
	if auction.GetStatus() != fundraisingtypes.AuctionStatusCancelled {
		if !blockTime.After(auction.GetStartTime().Add(params.WithdrawalDelay)) {
			return nil, sdkerrors.Wrapf(types.ErrAllocationWithdrawalTimeNotReached, "withdrawal for auction %d not yet allowed", msg.AuctionID)
		}
	}

	auctionUsedAllocations, err := k.AuctionUsedAllocations.Get(ctx, collections.Join(sdk.AccAddress(participantAddress), msg.AuctionID))
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrUsedAllocationsNotFound, "used allocations for auction %d not found", msg.AuctionID)
	}
	if auctionUsedAllocations.Withdrawn {
		return nil, sdkerrors.Wrapf(types.ErrAllocationsAlreadyWithdrawn, "allocations for auction %d already claimed", msg.AuctionID)
	}

	totalUsedAllocations, err := k.UsedAllocations.Get(ctx, msg.Participant)
	if err != nil {
		return nil, ignterrors.Criticalf("unable to find total used allocations entry for address %s", msg.Participant)
	}

	// decrease totalUsedAllocations making sure subtraction is feasible
	if totalUsedAllocations.NumAllocations.LT(auctionUsedAllocations.NumAllocations) {
		return nil, ignterrors.Critical("number of total used allocations cannot become negative")
	}
	totalUsedAllocations.NumAllocations = totalUsedAllocations.NumAllocations.Sub(auctionUsedAllocations.NumAllocations)

	auctionUsedAllocations.Withdrawn = true
	if err = k.AuctionUsedAllocations.Set(ctx, collections.Join(sdk.AccAddress(participantAddress), auctionUsedAllocations.AuctionID), auctionUsedAllocations); err != nil {
		return nil, ignterrors.Criticalf("unable to set auction used allocations entry for address %s", msg.Participant)
	}

	if err = k.UsedAllocations.Set(ctx, totalUsedAllocations.Address, totalUsedAllocations); err != nil {
		return nil, ignterrors.Criticalf("unable to set total used allocations entry for address %s", msg.Participant)
	}

	return &types.MsgWithdrawAllocationsResponse{}, sdkCtx.EventManager().EmitTypedEvent(&types.EventAllocationsWithdrawn{
		Participant: msg.Participant,
		AuctionID:   msg.AuctionID,
	})
}
