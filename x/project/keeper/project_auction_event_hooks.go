package keeper

import (
	"context"
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"

	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/project/types"
)

// EmitProjectAuctionCreated emits EventProjectAuctionCreated event if an auction is created for a project from a coordinator
func (k Keeper) EmitProjectAuctionCreated(
	ctx context.Context,
	auctionID uint64,
	auctioneer string,
	sellingCoin sdk.Coin,
) (bool, error) {
	projectID, err := types.VoucherProject(sellingCoin.Denom)
	if err != nil {
		// not a project auction
		return false, nil
	}

	// verify the auctioneer is the coordinator of the project
	project, err := k.GetProject(ctx, projectID)
	if err != nil {
		return false, sdkerrors.Wrapf(types.ErrProjectNotFound,
			"voucher %s is associated to an non-existing project %d",
			sellingCoin.Denom,
			projectID,
		)
	}
	coordinator, err := k.profileKeeper.GetCoordinator(ctx, project.CoordinatorId)
	if err != nil {
		return false, sdkerrors.Wrapf(profiletypes.ErrCoordinatorInvalid,
			"project %d coordinator doesn't exist %d",
			projectID,
			project.CoordinatorId,
		)
	}

	// if the coordinator if the auctioneer, we emit a ProjectAuctionCreated event
	if coordinator.Address != auctioneer {
		return false, nil
	}

	err = sdk.UnwrapSDKContext(ctx).EventManager().EmitTypedEvents(
		&types.EventProjectAuctionCreated{
			ProjectId: projectID,
			AuctionId: auctionID,
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

// ProjectAuctionEventHooks returns a ProjectAuctionEventHooks associated with the project keeper
func (k Keeper) ProjectAuctionEventHooks() ProjectAuctionEventHooks {
	return ProjectAuctionEventHooks{
		projectKeeper: k,
	}
}

// ProjectAuctionEventHooks implements fundraising hooks and emit events on auction creation
type ProjectAuctionEventHooks struct {
	projectKeeper Keeper
}

// Implements FundraisingHooks interface
var _ fundraisingtypes.FundraisingHooks = ProjectAuctionEventHooks{}

// AfterFixedPriceAuctionCreated emits a ProjectAuctionCreated event if created for a project
func (h ProjectAuctionEventHooks) AfterFixedPriceAuctionCreated(
	ctx context.Context,
	auctionID uint64,
	auctioneer string,
	_ sdkmath.LegacyDec,
	sellingCoin sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ time.Time,
	_ time.Time,
) error {
	_, err := h.projectKeeper.EmitProjectAuctionCreated(ctx, auctionID, auctioneer, sellingCoin)
	return err
}

// AfterBatchAuctionCreated emits a ProjectAuctionCreated event if created for a project
func (h ProjectAuctionEventHooks) AfterBatchAuctionCreated(
	ctx context.Context,
	auctionID uint64,
	auctioneer string,
	_ sdkmath.LegacyDec,
	_ sdkmath.LegacyDec,
	sellingCoin sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ uint32,
	_ sdkmath.LegacyDec,
	_ time.Time,
	_ time.Time,
) error {
	_, err := h.projectKeeper.EmitProjectAuctionCreated(ctx, auctionID, auctioneer, sellingCoin)
	return err
}

// BeforeFixedPriceAuctionCreated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeFixedPriceAuctionCreated(
	_ context.Context,
	_ string,
	_ sdkmath.LegacyDec,
	_ sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ time.Time,
	_ time.Time,
) error {
	return nil
}

// BeforeBatchAuctionCreated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeBatchAuctionCreated(
	_ context.Context,
	_ string,
	_ sdkmath.LegacyDec,
	_ sdkmath.LegacyDec,
	_ sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ uint32,
	_ sdkmath.LegacyDec,
	_ time.Time,
	_ time.Time,
) error {
	return nil
}

// BeforeAuctionCanceled implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeAuctionCanceled(
	_ context.Context,
	_ uint64,
	_ string,
) error {
	return nil
}

// BeforeBidPlaced implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeBidPlaced(
	_ context.Context,
	_ uint64,
	_ uint64,
	_ string,
	_ fundraisingtypes.BidType,
	_ sdkmath.LegacyDec,
	_ sdk.Coin,
) error {
	return nil
}

// BeforeBidModified implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeBidModified(
	_ context.Context,
	_ uint64,
	_ uint64,
	_ string,
	_ fundraisingtypes.BidType,
	_ sdkmath.LegacyDec,
	_ sdk.Coin,
) error {
	return nil
}

// BeforeAllowedBiddersAdded implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeAllowedBiddersAdded(
	_ context.Context,
	_ []fundraisingtypes.AllowedBidder,
) error {
	return nil
}

// BeforeAllowedBidderUpdated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeAllowedBidderUpdated(
	_ context.Context,
	_ uint64,
	_ sdk.AccAddress,
	_ sdkmath.Int,
) error {
	return nil
}

// BeforeSellingCoinsAllocated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeSellingCoinsAllocated(
	_ context.Context,
	_ uint64,
	_ map[string]sdkmath.Int,
	_ map[string]sdkmath.Int,
) error {
	return nil
}
