package keeper_test

import (
	"testing"

	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"

	networktypes "github.com/ignite/network/pkg/types"
	tc "github.com/ignite/network/testutil/constructor"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/monitoringc/types"
	profiletypes "github.com/ignite/network/x/profile/types"
	rewardtypes "github.com/ignite/network/x/reward/types"
)

func Test_OnRecvMonitoringPacket(t *testing.T) {
	var (
		ctx, tk, _     = testkeeper.NewTestSetup(t)
		invalidChannel = "invalidchannel"
		validChannel   = "monitoringtest"
		chain          = sample.Chain(r, 0, 0)
		valFoo         = sample.Address(r)
		valBar         = sample.Address(r)
		valOpAddrFoo   = sample.Address(r)
		valOpAddrBar   = sample.Address(r)
		coins          = sample.Coins(r)
	)

	err := tk.MonitoringConsumerKeeper.LaunchIDFromChannelID.Set(ctx, invalidChannel, types.LaunchIDFromChannelID{
		ChannelID: invalidChannel,
		LaunchID:  10000,
	})
	require.NoError(t, err)
	chain.LaunchID, err = tk.LaunchKeeper.AppendChain(ctx, chain)
	require.NoError(t, err)
	err = tk.MonitoringConsumerKeeper.LaunchIDFromChannelID.Set(ctx, validChannel, types.LaunchIDFromChannelID{
		ChannelID: validChannel,
		LaunchID:  chain.LaunchID,
	})
	require.NoError(t, err)

	t.Run("should allow set reward pool", func(t *testing.T) {
		err := tk.RewardKeeper.RewardPool.Set(ctx, chain.LaunchID, rewardtypes.RewardPool{
			LaunchID:         chain.LaunchID,
			Provider:         sample.Address(r),
			InitialCoins:     coins,
			RemainingCoins:   coins,
			LastRewardHeight: 1,
			Closed:           false,
		})
		require.NoError(t, err)
		err = tk.BankKeeper.MintCoins(ctx, rewardtypes.ModuleName, coins)
		require.NoError(t, err)
	})

	// set validator profiles
	err = tk.ProfileKeeper.Validator.Set(ctx, valFoo, profiletypes.Validator{
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

	tests := []struct {
		name   string
		packet channeltypes.Packet
		data   networktypes.MonitoringPacket
		valid  bool
	}{
		{
			name: "should successfully distribute rewards",
			packet: channeltypes.Packet{
				DestinationChannel: validChannel,
			},
			data: networktypes.MonitoringPacket{
				BlockHeight: 10,
				SignatureCounts: tc.SignatureCounts(10,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
			},
			valid: true,
		},
		{
			name:   "should prevent invalid data",
			packet: channeltypes.Packet{},
			data: networktypes.MonitoringPacket{
				BlockHeight: 0,
				SignatureCounts: networktypes.SignatureCounts{
					BlockCount: 1,
				},
			},
			valid: false,
		},
		{
			name: "should prevent no launch ID associated to channel ID",
			packet: channeltypes.Packet{
				DestinationChannel: "invalid",
			},
			data: networktypes.MonitoringPacket{
				BlockHeight: 1,
				SignatureCounts: networktypes.SignatureCounts{
					BlockCount: 1,
				},
			},
			valid: false,
		},
		{
			name: "should fail distribute rewards",
			packet: channeltypes.Packet{
				DestinationChannel: invalidChannel,
			},
			data: networktypes.MonitoringPacket{
				BlockHeight: 1,
				SignatureCounts: networktypes.SignatureCounts{
					BlockCount: 1,
				},
			},
			valid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tk.MonitoringConsumerKeeper.OnRecvMonitoringPacket(ctx, tt.packet, tt.data)
			if !tt.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func Test_OnAcknowledgementMonitoringPacket(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should return not implemented", func(t *testing.T) {
		err := tk.MonitoringConsumerKeeper.OnAcknowledgementMonitoringPacket(
			ctx,
			channeltypes.Packet{},
			networktypes.MonitoringPacket{},
			channeltypes.Acknowledgement{},
		)
		require.EqualError(t, err, "not implemented")
	})
}

func Test_OnTimeoutMonitoringPacket(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	t.Run("should return not implemented", func(t *testing.T) {
		err := tk.MonitoringConsumerKeeper.OnTimeoutMonitoringPacket(
			ctx,
			channeltypes.Packet{},
			networktypes.MonitoringPacket{},
		)
		require.EqualError(t, err, "not implemented")
	})
}
