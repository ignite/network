package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/pkg/errors"

	networktypes "github.com/ignite/network/pkg/types"
	"github.com/ignite/network/x/monitoringc/types"
)

// OnRecvMonitoringPacket processes packet reception
func (k Keeper) OnRecvMonitoringPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data networktypes.MonitoringPacket,
) (packetAck networktypes.MonitoringPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// retrieve launch ID for channel ID
	lidFromCid, err := k.LaunchIDFromChannelID.Get(ctx, packet.DestinationChannel)
	if err != nil {
		return packetAck, errors.Wrapf(err, "no launch ID associated to channel ID %s", packet.DestinationChannel)
	}

	// save the latest received monitoring packet for documentation purpose
	err = k.MonitoringHistory.Set(ctx, lidFromCid.LaunchId, types.MonitoringHistory{
		LaunchId:               lidFromCid.LaunchId,
		LatestMonitoringPacket: data,
	})
	if err != nil {
	}

	// distribute reward from the signature count
	err = k.rewardKeeper.DistributeRewards(
		ctx,
		lidFromCid.LaunchId,
		data.SignatureCounts,
		data.BlockHeight,
		true,
	)

	return packetAck, err
}

// OnAcknowledgementMonitoringPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementMonitoringPacket(
	_ sdk.Context,
	_ channeltypes.Packet,
	_ networktypes.MonitoringPacket,
	_ channeltypes.Acknowledgement,
) error {
	return errors.New("not implemented")
}

// OnTimeoutMonitoringPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutMonitoringPacket(
	_ sdk.Context,
	_ channeltypes.Packet,
	_ networktypes.MonitoringPacket,
) error {
	return errors.New("not implemented")
}
