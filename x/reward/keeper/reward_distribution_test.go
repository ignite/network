package keeper_test

import (
	"fmt"
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	ignterrors "github.com/ignite/network/pkg/errors"
	networktypes "github.com/ignite/network/pkg/types"
	tc "github.com/ignite/network/testutil/constructor"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	profiletypes "github.com/ignite/network/x/profile/types"
	"github.com/ignite/network/x/reward/keeper"
	"github.com/ignite/network/x/reward/types"
)

func TestCalculateRewards(t *testing.T) {
	type args struct {
		blockRatio sdkmath.LegacyDec
		sigRatio   sdkmath.LegacyDec
		coins      sdk.Coins
	}
	tests := []struct {
		name    string
		args    args
		want    sdk.Coins
		wantErr bool
	}{
		{
			name: "should give zero rewards with zero ratios and zero coins ",
			args: args{
				blockRatio: sdkmath.LegacyZeroDec(),
				sigRatio:   sdkmath.LegacyZeroDec(),
				coins:      sdk.NewCoins(),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "should give zero rewards with nil coins ",
			args: args{
				blockRatio: sdkmath.LegacyOneDec(),
				sigRatio:   sdkmath.LegacyOneDec(),
				coins:      nil,
			},
			want: sdk.NewCoins(),
		},
		{
			name: "should give 0 rewards with 0 block ratio ",
			args: args{
				blockRatio: sdkmath.LegacyZeroDec(),
				sigRatio:   sdkmath.LegacyOneDec(),
				coins:      tc.Coins(t, "10aaa,10bbb,10ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "should give 0 rewards with 0 signature ratio ",
			args: args{
				blockRatio: sdkmath.LegacyOneDec(),
				sigRatio:   sdkmath.LegacyZeroDec(),
				coins:      tc.Coins(t, "10aaa,10bbb,10ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "should give all rewards with full block and signature ratios ",
			args: args{
				blockRatio: sdkmath.LegacyOneDec(),
				sigRatio:   sdkmath.LegacyOneDec(),
				coins:      tc.Coins(t, "10aaa,10bbb,10ccc"),
			},
			want: tc.Coins(t, "10aaa,10bbb,10ccc"),
		},
		{
			name: "should give half rewards with 0.5 block ratio ",
			args: args{
				blockRatio: tc.Dec(t, "0.5"),
				sigRatio:   sdkmath.LegacyOneDec(),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "5aaa,50bbb,500ccc"),
		},
		{
			name: "should give half rewards with 0.5 signature ratio ",
			args: args{
				blockRatio: sdkmath.LegacyOneDec(),
				sigRatio:   tc.Dec(t, "0.5"),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "5aaa,50bbb,500ccc"),
		},
		{
			name: "should give 0.2 rewards with 0.5 block ratio and 0.4 signature ratio ",
			args: args{
				blockRatio: tc.Dec(t, "0.5"),
				sigRatio:   tc.Dec(t, "0.4"),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "2aaa,20bbb,200ccc"),
		},
		{
			name: "should be truncate with decimal rewards ",
			args: args{
				blockRatio: tc.Dec(t, "0.5"),
				sigRatio:   sdkmath.LegacyOneDec(),
				coins:      tc.Coins(t, "1aaa,11bbb,101ccc"),
			},
			want: tc.Coins(t, "5bbb,50ccc"),
		},
		{
			name: "should give 0.01 rewards with 0.1 block ratio and 0.1 signature ratio ",
			args: args{
				blockRatio: tc.Dec(t, "0.1"),
				sigRatio:   tc.Dec(t, "0.1"),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: tc.Coins(t, "1bbb,10ccc"),
		},
		{
			name: "should be empty coins rewards if all rewards are fully truncated",
			args: args{
				blockRatio: tc.Dec(t, "0.0001"),
				sigRatio:   sdkmath.LegacyOneDec(),
				coins:      tc.Coins(t, "10aaa,100bbb,1000ccc"),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "should return empty coins with empty coins ",
			args: args{
				blockRatio: sdkmath.LegacyOneDec(),
				sigRatio:   sdkmath.LegacyOneDec(),
				coins:      sdk.NewCoins(),
			},
			want: sdk.NewCoins(),
		},
		{
			name: "should prevent using block ratio greater than 1",
			args: args{
				blockRatio: tc.Dec(t, "1.000001"),
				sigRatio:   sdkmath.LegacyZeroDec(),
				coins:      sample.Coins(r),
			},
			wantErr: true,
		},
		{
			name: "should prevent using signature ratio greater than 1",
			args: args{
				blockRatio: sdkmath.LegacyZeroDec(),
				sigRatio:   tc.Dec(t, "1.000001"),
				coins:      sample.Coins(r),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := keeper.CalculateRewards(tt.args.blockRatio, tt.args.sigRatio, tt.args.coins)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.True(t, got.Equal(tt.want),
				fmt.Sprintf("want: %s, got: %s", tt.want.String(), got.String()),
			)
		})
	}
}

func TestKeeper_DistributeRewards(t *testing.T) {
	var (
		ctx, tk, _      = testkeeper.NewTestSetup(t)
		valFoo          = sample.Address(r)
		valBar          = sample.Address(r)
		valOpAddrFoo    = sample.Address(r)
		valOpAddrBar    = sample.Address(r)
		noProfileVal    = sample.Address(r)
		notFoundValAddr = sample.Address(r)
		provider        = sample.Address(r)
	)

	// set validator profiles
	err := tk.ProfileKeeper.Validator.Set(ctx, valFoo, profiletypes.Validator{
		Address:           valFoo,
		OperatorAddresses: []string{valOpAddrFoo},
	})
	require.NoError(t, err)

	err = tk.ProfileKeeper.ValidatorByOperatorAddress.Set(ctx, valOpAddrFoo, profiletypes.ValidatorByOperatorAddress{
		ValidatorAddress: valFoo,
		OperatorAddress:  valOpAddrFoo,
	})
	require.NoError(t, err)

	err = tk.ProfileKeeper.Validator.Set(ctx, valBar, profiletypes.Validator{
		Address:           valBar,
		OperatorAddresses: []string{valOpAddrBar},
	})
	require.NoError(t, err)

	err = tk.ProfileKeeper.ValidatorByOperatorAddress.Set(ctx, valOpAddrBar, profiletypes.ValidatorByOperatorAddress{
		ValidatorAddress: valBar,
		OperatorAddress:  valOpAddrBar,
	})
	require.NoError(t, err)

	err = tk.ProfileKeeper.ValidatorByOperatorAddress.Set(ctx, notFoundValAddr, profiletypes.ValidatorByOperatorAddress{
		ValidatorAddress: sample.Address(r),
		OperatorAddress:  notFoundValAddr,
	})
	require.NoError(t, err)

	type args struct {
		launchID        uint64
		signatureCounts networktypes.SignatureCounts
		lastBlockHeight int64
		closeRewardPool bool
	}
	tests := []struct {
		name         string
		rewardPool   types.RewardPool
		args         args
		wantBalances map[string]sdk.Coins
		err          error
	}{
		{
			name: "should allow distributing rewards",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 10,
				closeRewardPool: true,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "50aaa,50bbb"),
				valBar:   tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "should allow distributing reward with different signature ratios",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,1000bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,1000bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.2"),
					tc.SignatureCount(t, valOpAddrBar, "0.8"),
				),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "20aaa,200bbb"),
				valBar:   tc.Coins(t, "80aaa,800bbb"),
			},
		},
		{
			name: "should allow current reward height to influence block ratio for reward distribution",
			rewardPool: types.RewardPool{
				LaunchId:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 10,
				LastRewardHeight:    20,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 15,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "25aaa,25bbb"),
				valBar:   tc.Coins(t, "25aaa,25bbb"),
			},
		},
		{
			name: "should distribute all rewards if closing the reward pool ",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 5,
				closeRewardPool: true,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "50aaa,50bbb"),
				valFoo:   tc.Coins(t, "25aaa,25bbb"),
				valBar:   tc.Coins(t, "25aaa,25bbb"),
			},
		},
		{
			name: "should distribute part of the reward if last reward height not reached and reward pool not closed ",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 5,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "25aaa,25bbb"),
				valBar:   tc.Coins(t, "25aaa,25bbb"),
			},
		},
		{
			name: "should distribute all rewards if last block height greater than reward pool last reward height ",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "50aaa,50bbb"),
				valBar:   tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "should clamp block height to 1 if ratio GT 1",
			rewardPool: types.RewardPool{
				LaunchId:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 9,
				LastRewardHeight:    10,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
				lastBlockHeight: 11,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: sdk.NewCoins(),
				valFoo:   tc.Coins(t, "50aaa,50bbb"),
				valBar:   tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "should distribute rewards to the operator address for validator with no profile ",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.3"),
					tc.SignatureCount(t, valOpAddrBar, "0.3"),
					tc.SignatureCount(t, noProfileVal, "0.3"),
				),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider:     tc.Coins(t, "10aaa,10bbb"),
				valFoo:       tc.Coins(t, "30aaa,30bbb"),
				valBar:       tc.Coins(t, "30aaa,30bbb"),
				noProfileVal: tc.Coins(t, "30aaa,30bbb"),
			},
		},
		{
			name: "should refund all rewards if the reward pool is closed and no signature counts are reported",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 5,
				closeRewardPool: true,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "100aaa,100bbb"),
			},
		},
		{
			name: "should refund all rewards if the last reward height is reached and no signature counts are reported",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 10,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "100aaa,100bbb"),
			},
		},
		{
			name: "should refund all rewards relative to the block ratio if the reward " +
				"pool is not closed and no signature counts are reported",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 5,
				closeRewardPool: false,
			},
			wantBalances: map[string]sdk.Coins{
				provider: tc.Coins(t, "50aaa,50bbb"),
			},
		},
		{
			name: "should prevent invalid signature counts for negative reward pool with critical error",
			rewardPool: types.RewardPool{
				LaunchId:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 10,
				LastRewardHeight:    20,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.6"),
				),
				lastBlockHeight: 20,
				closeRewardPool: false,
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should return critical error for signatureRatio GT 1",
			rewardPool: types.RewardPool{
				LaunchId:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 10,
				LastRewardHeight:    20,
				Closed:              false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "1.0001"),
				),
				lastBlockHeight: 20,
				closeRewardPool: false,
			},
			err: ignterrors.ErrCritical,
		},
		{
			name: "should prevent distributing rewards with a non-existent reward pool",
			args: args{
				launchID: 99999,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
				),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrRewardPoolNotFound,
		},
		{
			name: "should prevent distributing rewards from a closed reward pool",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           true,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
				),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrRewardPoolClosed,
		},
		{
			name: "should prevent distributing rewards if signature counts are invalid",
			rewardPool: types.RewardPool{
				LaunchId:         1,
				Provider:         provider,
				InitialCoins:     tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:   tc.Coins(t, "100aaa,100bbb"),
				LastRewardHeight: 10,
				Closed:           false,
			},
			args: args{
				launchID: 1,
				signatureCounts: tc.SignatureCounts(1,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, "invalid-bech32-address", "0.5"),
				),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrInvalidSignatureCounts,
		},
		{
			name: "should prevent providing a last block height lower than the current reward height",
			rewardPool: types.RewardPool{
				LaunchId:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 5,
				LastRewardHeight:    10,
				Closed:              false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 1,
				closeRewardPool: false,
			},
			err: types.ErrInvalidLastBlockHeight,
		},
		{
			name: "should prevent providing a last block height equals to the current reward height",
			rewardPool: types.RewardPool{
				LaunchId:            1,
				Provider:            provider,
				InitialCoins:        tc.Coins(t, "100aaa,100bbb"),
				RemainingCoins:      tc.Coins(t, "100aaa,100bbb"),
				CurrentRewardHeight: 5,
				LastRewardHeight:    10,
				Closed:              false,
			},
			args: args{
				launchID:        1,
				signatureCounts: tc.SignatureCounts(1),
				lastBlockHeight: 5,
				closeRewardPool: false,
			},
			err: types.ErrInvalidLastBlockHeight,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set test reward pool if contains coins
			if tt.rewardPool.RemainingCoins != nil {
				err := tk.RewardKeeper.RewardPool.Set(ctx, tt.rewardPool.LaunchId, tt.rewardPool)
				require.NoError(t, err)
				err = tk.BankKeeper.MintCoins(ctx, types.ModuleName, tt.rewardPool.RemainingCoins)
				require.NoError(t, err)
			}

			err := tk.RewardKeeper.DistributeRewards(ctx,
				tt.args.launchID,
				tt.args.signatureCounts,
				tt.args.lastBlockHeight,
				tt.args.closeRewardPool,
			)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			rewardPool, err := tk.RewardKeeper.RewardPool.Get(ctx, tt.args.launchID)
			require.NoError(t, err)
			require.Equal(t, tt.rewardPool.InitialCoins, rewardPool.InitialCoins)
			require.Equal(t, tt.rewardPool.Provider, rewardPool.Provider)

			// check if reward pool should be closed
			if tt.args.closeRewardPool || tt.args.lastBlockHeight >= rewardPool.LastRewardHeight {
				require.True(t, rewardPool.Closed)
			} else {
				require.Equal(t, tt.args.lastBlockHeight, rewardPool.CurrentRewardHeight)
			}

			totalDistributedBalances := sdk.NewCoins()
			for wantAddr, wantBalance := range tt.wantBalances {
				t.Run(fmt.Sprintf("check balance %s", wantAddr), func(t *testing.T) {
					wantAcc, err := tk.RewardKeeper.AddressCodec().StringToBytes(wantAddr)
					require.NoError(t, err)

					balance := tk.BankKeeper.GetAllBalances(ctx, wantAcc)
					require.True(t, balance.Equal(wantBalance),
						fmt.Sprintf("address: %s,  want: %s, got: %s",
							wantAddr, wantBalance.String(), balance.String(),
						),
					)
					totalDistributedBalances = totalDistributedBalances.Add(balance...)

					// remove the test balance
					err = tk.BankKeeper.SendCoinsFromAccountToModule(ctx, wantAcc, types.ModuleName, balance)
					require.NoError(t, err)
					err = tk.BankKeeper.BurnCoins(ctx, types.ModuleName, balance)
					require.NoError(t, err)
				})
			}

			// assert currentRemainingCoins = previousRemainingCoins - distributedRewards
			expectedRemainingCoins, neg := tt.rewardPool.RemainingCoins.SafeSub(totalDistributedBalances...)
			require.False(t, neg, "more coins have been distributed than coins in remaining coins %s > %s",
				totalDistributedBalances.String(),
				tt.rewardPool.RemainingCoins.String(),
			)
			require.True(t, rewardPool.RemainingCoins.Equal(expectedRemainingCoins), "expected remaining coins %s, got %s",
				expectedRemainingCoins.String(),
				rewardPool.RemainingCoins.String(),
			)

			// remove the reward pool used for the test
			require.NoError(t, tk.RewardKeeper.RewardPool.Remove(ctx, tt.rewardPool.LaunchId))
		})
	}
}
