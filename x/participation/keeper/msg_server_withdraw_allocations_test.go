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

func Test_msgServer_WithdrawAllocations(t *testing.T) {
	var (
		ctx, tk, ts         = testkeeper.NewTestSetup(t)
		auctioneer          = sample.AccAddress(r)
		validParticipant    = sample.AccAddress(r)
		invalidParticipant  = sample.AccAddress(r)
		auctionStartTime    = ctx.BlockTime().Add(time.Hour)
		auctionEndTime      = ctx.BlockTime().Add(time.Hour * 24 * 7)
		validWithdrawalTime = auctionStartTime.Add(time.Hour * 10)
		withdrawalDelay     = time.Hour * 5
	)

	params := types.DefaultParams()
	params.WithdrawalDelay = withdrawalDelay
	params.AllocationPrice = types.AllocationPrice{Bonded: sdkmath.NewInt(100)}
	err := tk.ParticipationKeeper.Params.Set(ctx, params)
	require.NoError(t, err)

	auctionSellingCoin := sample.CoinWithRange(r, params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64(),
		params.ParticipationTierList[1].Benefits.MaxBidAmount.Int64()+1000)

	// delegate some coins so participant has some allocations to use
	_, _, err = tk.DelegateN(ctx, r, validParticipant.String(), 100, 10)
	require.NoError(t, err)

	// initialize an auction
	tk.Mint(ctx, auctioneer.String(), sdk.NewCoins(auctionSellingCoin))
	auctionID := tk.CreateFixedPriceAuction(ctx, r, auctioneer.String(), auctionSellingCoin, auctionStartTime, auctionEndTime)

	// initialize another auction that will be set to `cancelled`
	tk.Mint(ctx, auctioneer.String(), sdk.NewCoins(auctionSellingCoin))
	cancelledAuctionID := tk.CreateFixedPriceAuction(ctx, r, auctioneer.String(), auctionSellingCoin, auctionStartTime, auctionEndTime)

	t.Run("should allow participation", func(t *testing.T) {
		_, err := ts.ParticipationSrv.Participate(ctx, &types.MsgParticipate{
			Participant: validParticipant.String(),
			AuctionId:   auctionID,
			TierId:      1,
		})
		require.NoError(t, err)
		_, err = ts.ParticipationSrv.Participate(ctx, &types.MsgParticipate{
			Participant: validParticipant.String(),
			AuctionId:   cancelledAuctionID,
			TierId:      1,
		})
		require.NoError(t, err)
	})

	t.Run("should allow auction cancel", func(t *testing.T) {
		err := tk.FundraisingKeeper.CancelAuction(ctx, fundraisingtypes.NewMsgCancelAuction(auctioneer.String(), cancelledAuctionID))
		require.NoError(t, err)
	})

	// manually insert entry for invalidParticipant for later test
	err = tk.ParticipationKeeper.AuctionUsedAllocations.Set(ctx, collections.Join(invalidParticipant, auctionID), types.AuctionUsedAllocations{
		Address:        invalidParticipant.String(),
		AuctionId:      auctionID,
		NumAllocations: sdkmath.OneInt(),
		Withdrawn:      true, // set withdrawn to true
	})
	require.NoError(t, err)

	tests := []struct {
		name      string
		msg       *types.MsgWithdrawAllocations
		blockTime time.Time
		err       error
	}{
		{
			name: "should allow to remove allocations",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant.String(),
				AuctionId:   auctionID,
			},
			blockTime: validWithdrawalTime,
		},
		{
			name: "should allow to remove allocations if auction status is cancelled",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant.String(),
				AuctionId:   cancelledAuctionID,
			},
			blockTime: auctionStartTime,
		},
		{
			name: "should return auction not found",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant.String(),
				AuctionId:   auctionID + 1000,
			},
			blockTime: validWithdrawalTime,
			err:       fundraisingtypes.ErrAuctionNotFound,
		},
		{
			name: "should prevent withdrawal before withdrawal delay has passed",
			msg: &types.MsgWithdrawAllocations{
				Participant: validParticipant.String(),
				AuctionId:   auctionID,
			},
			blockTime: auctionStartTime,
			err:       types.ErrAllocationWithdrawalTimeNotReached,
		},
		{
			name: "should return used allocations not found",
			msg: &types.MsgWithdrawAllocations{
				Participant: sample.Address(r),
				AuctionId:   auctionID,
			},
			blockTime: validWithdrawalTime,
			err:       types.ErrUsedAllocationsNotFound,
		},
		{
			name: "should prevent withdrawal if already claimed",
			msg: &types.MsgWithdrawAllocations{
				Participant: invalidParticipant.String(),
				AuctionId:   auctionID,
			},
			blockTime: validWithdrawalTime,
			err:       types.ErrAllocationsAlreadyWithdrawn,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preUsedAllocations, err := tk.ParticipationKeeper.UsedAllocations.Get(ctx, tt.msg.Participant)
			if tt.err == nil {
				// check if valid only when no error expected
				require.NoError(t, err)
			}

			preAuctionUsedAllocations, err := tk.ParticipationKeeper.AuctionUsedAllocations.Get(ctx, collections.Join(validParticipant, tt.msg.AuctionId))
			if tt.err == nil {
				// check if valid only when no error expected
				require.NoError(t, err)
				require.False(t, preAuctionUsedAllocations.Withdrawn)
			}

			// set wanted block time
			tmpCtx := ctx.WithBlockTime(tt.blockTime)
			_, err = ts.ParticipationSrv.WithdrawAllocations(tmpCtx, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			// check auctionUsedAllocations is set to `withdrawn`
			participantAddress, err := tk.ParticipationKeeper.AddressCodec().StringToBytes(tt.msg.Participant)
			require.NoError(t, err)
			postAuctionUsedAllocations, err := tk.ParticipationKeeper.AuctionUsedAllocations.Get(tmpCtx, collections.Join(sdk.AccAddress(participantAddress), tt.msg.AuctionId))
			require.NoError(t, err)
			require.True(t, postAuctionUsedAllocations.Withdrawn)
			require.Equal(t, preAuctionUsedAllocations.NumAllocations, postAuctionUsedAllocations.NumAllocations)

			// check usedAllocationEntry is correctly decreased
			postUsedAllocations, err := tk.ParticipationKeeper.UsedAllocations.Get(tmpCtx, tt.msg.Participant)
			require.NoError(t, err)
			calculated := preUsedAllocations.NumAllocations.Sub(preAuctionUsedAllocations.NumAllocations)
			require.True(t, postUsedAllocations.NumAllocations.Equal(calculated), "numAlloc %s not equal to calculated %s", postUsedAllocations.NumAllocations, calculated)
		})
	}
}
