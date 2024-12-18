package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/participation/types"
)

func Test_msgServer_Participate(t *testing.T) {
	var (
		ctx, tk, ts                      = testkeeper.NewTestSetup(t)
		auctioneer                       = sample.Address(r)
		registrationPeriod               = time.Hour * 5 // 5 hours before start
		startTime                        = ctx.BlockTime().Add(time.Hour * 10)
		startTimeLowerRegistrationPeriod = time.Unix(int64((registrationPeriod - time.Hour).Seconds()), 0)
		endTime                          = ctx.BlockTime().Add(time.Hour * 24 * 7)
		validRegistrationTime            = ctx.BlockTime().Add(time.Hour * 6)
		allocationPrice                  = types.AllocationPrice{Bonded: sdkmath.NewInt(100)}
		addrsWithDelsTier                = []string{sample.Address(r), sample.Address(r), sample.Address(r), sample.Address(r)}
		availableAllocsTier              = make([]sdkmath.Int, len(addrsWithDelsTier))
		auctionRegistrationPeriodID      uint64
		auctionStartedID                 uint64
		auctionLowerRegistrationPeriodID uint64
		auctionCancelledID               uint64
	)

	params := types.DefaultParams()
	params.AllocationPrice = allocationPrice
	params.RegistrationPeriod = registrationPeriod
	err := tk.ParticipationKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	sellingCoin1 := sample.CoinWithRange(r, params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64(),
		params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64()+1000)

	sellingCoin2 := sample.CoinWithRange(r, params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64(),
		params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64()+1000)

	t.Run("should allow initialize auctions", func(t *testing.T) {
		tk.Mint(ctx, auctioneer, sdk.NewCoins(sellingCoin1))
		auctionRegistrationPeriodID = tk.CreateFixedPriceAuction(ctx, r, auctioneer, sellingCoin1, startTime, endTime)
		// initialize auction with edge case start time (forcefully set to status standby)
		tk.Mint(ctx, auctioneer, sdk.NewCoins(sellingCoin2))
		auctionLowerRegistrationPeriodID = tk.CreateFixedPriceAuction(ctx, r, auctioneer, sellingCoin2, startTimeLowerRegistrationPeriod, endTime)
		auctionLowerRegistrationPeriod, err := tk.FundraisingKeeper.GetAuction(ctx, auctionLowerRegistrationPeriodID)
		require.NoError(t, err)
		err = auctionLowerRegistrationPeriod.SetStatus(fundraisingtypes.AuctionStatusStandBy)
		require.NoError(t, err)
		err = tk.FundraisingKeeper.Auction.Set(ctx, auctionLowerRegistrationPeriod.GetId(), auctionLowerRegistrationPeriod)
		require.NoError(t, err)
		// initialize auction that is already started
		tk.Mint(ctx, auctioneer, sdk.NewCoins(sellingCoin1))
		auctionStartedID = tk.CreateFixedPriceAuction(ctx, r, auctioneer, sellingCoin1, ctx.BlockTime(), endTime)
		// initialize auction that will be set to `cancelled`
		tk.Mint(ctx, auctioneer, sdk.NewCoins(sellingCoin1))
		auctionCancelledID = tk.CreateFixedPriceAuction(ctx, r, auctioneer, sellingCoin1, startTime, endTime)
		// cancel auction
		err = tk.FundraisingKeeper.CancelAuction(ctx, fundraisingtypes.NewMsgCancelAuction(auctioneer, auctionCancelledID))
		require.NoError(t, err)
	})

	// add delegations
	for i := 0; i < len(addrsWithDelsTier); i++ {
		_, _, err := tk.DelegateN(ctx, r, addrsWithDelsTier[i], 100, 10)
		require.NoError(t, err)
		availableAllocsTier[i], err = tk.ParticipationKeeper.GetAvailableAllocations(ctx, addrsWithDelsTier[i])
		require.NoError(t, err)
		require.EqualValues(t, sdkmath.NewInt(10), availableAllocsTier[i])
	}

	tests := []struct {
		name                  string
		msg                   *types.MsgParticipate
		desiredUsedAlloc      sdkmath.Int
		currentAvailableAlloc sdkmath.Int
		blockTime             time.Time
		err                   error
	}{
		{
			name: "should allow valid message tier 1",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[0],
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      1,
			},
			desiredUsedAlloc:      sdkmath.OneInt(),
			currentAvailableAlloc: availableAllocsTier[0],
			blockTime:             validRegistrationTime,
		},
		{
			name: "should allow valid message tier 2",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      2,
			},
			desiredUsedAlloc:      sdkmath.NewInt(2),
			currentAvailableAlloc: availableAllocsTier[1],
			blockTime:             validRegistrationTime,
		},
		{
			name: "should allow participation when registration period is longer than range between Unix time 0 and auction's start time",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[2],
				AuctionId:   auctionLowerRegistrationPeriodID,
				TierId:      1,
			},
			desiredUsedAlloc:      sdkmath.OneInt(),
			currentAvailableAlloc: availableAllocsTier[2],
			blockTime:             time.Unix(1, 0),
		},
		{
			name: "should prevent invalid address",
			msg: &types.MsgParticipate{
				Participant: "",
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      1,
			},
			err:       types.ErrInvalidSigner,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating twice in the same auction",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[0],
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      1,
			},
			err:       types.ErrAlreadyParticipating,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent if user has insufficient available allocations",
			msg: &types.MsgParticipate{
				Participant: sample.Address(r),
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      1,
			},
			err:       types.ErrInsufficientAllocations,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating using a non existent tier",
			msg: &types.MsgParticipate{
				Participant: sample.Address(r),
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      111,
			},
			err:       types.ErrTierNotFound,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating in a non existent auction",
			msg: &types.MsgParticipate{
				Participant: sample.Address(r),
				AuctionId:   auctionLowerRegistrationPeriodID + 1000,
				TierId:      1,
			},
			err:       fundraisingtypes.ErrAuctionNotFound,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating if auction cancelled",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionId:   auctionCancelledID,
				TierId:      1,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: validRegistrationTime,
		},
		{
			name: "should prevent participating if auction started",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionId:   auctionStartedID,
				TierId:      1,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: startTime.Add(time.Hour),
		},
		{
			name: "should prevent participating before registration period",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[1],
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      2,
			},
			err:       types.ErrParticipationNotAllowed,
			blockTime: ctx.BlockTime(),
		},
		{
			name: "should prevent participating if tier amount greater than auction max bid amount",
			msg: &types.MsgParticipate{
				Participant: addrsWithDelsTier[3],
				AuctionId:   auctionRegistrationPeriodID,
				TierId:      4,
			},
			err:       types.ErrInvalidBidder,
			blockTime: validRegistrationTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set wanted block time
			tmpCtx := ctx.WithBlockTime(tt.blockTime)

			_, err := ts.ParticipationSrv.Participate(tmpCtx, tt.msg)

			// check error
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			tier, found := types.GetTierFromID(types.DefaultParticipationTierList, tt.msg.TierId)
			require.True(t, found)

			// check auction contains allowed bidder
			allowedBidders, err := tk.FundraisingKeeper.GetAllowedBiddersByAuction(tmpCtx, tt.msg.AuctionId)
			require.NoError(t, err)
			require.Contains(t, allowedBidders, fundraisingtypes.AllowedBidder{
				Bidder:       tt.msg.Participant,
				MaxBidAmount: tier.Benefits.MaxBidAmount,
			})

			// check used allocations entry for bidder
			usedAllocations, err := tk.ParticipationKeeper.UsedAllocations.Get(tmpCtx, tt.msg.Participant)
			require.NoError(t, err)
			require.EqualValues(t, tt.desiredUsedAlloc, usedAllocations.NumAllocations)

			// check valid auction used allocations entry for bidder exists
			participantAddress, err := tk.ParticipationKeeper.AddressCodec().StringToBytes(tt.msg.Participant)
			require.NoError(t, err)
			auctionUsedAllocations, err := tk.ParticipationKeeper.AuctionUsedAllocations.Get(tmpCtx, collections.Join(sdk.AccAddress(participantAddress), tt.msg.AuctionId))
			require.NoError(t, err)
			require.Equal(t, tier.RequiredAllocations, auctionUsedAllocations.NumAllocations)
			require.False(t, auctionUsedAllocations.Withdrawn)

			// check that available allocations has decreased accordingly according to tier used
			availableAlloc, err := tk.ParticipationKeeper.GetAvailableAllocations(tmpCtx, tt.msg.Participant)
			require.NoError(t, err)
			require.True(t, found)
			require.EqualValues(t, tt.currentAvailableAlloc.Sub(tier.RequiredAllocations), availableAlloc)
		})
	}
}
