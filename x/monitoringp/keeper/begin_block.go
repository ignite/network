package keeper

import (
	"context"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/comet"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	"github.com/pkg/errors"

	networktypes "github.com/ignite/network/pkg/types"
	"github.com/ignite/network/x/monitoringp/types"
)

const (
	// MonitoringPacketTimeoutDelay is the delay before a monitoring packet is timed out
	// The timeout is set to one year
	// This is an arbitrarily chosen value that should never be reached in practice
	MonitoringPacketTimeoutDelay = time.Hour * 8760
)

// ReportBlockSignatures gets signatures from blocks and update monitoring info
func (k Keeper) ReportBlockSignatures(ctx context.Context, lastCommit comet.CommitInfo, blockHeight int64) error {
	// skip first block because it is not signed
	if blockHeight == 1 {
		return nil
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return err
	}

	// no report if last height is reached
	if blockHeight > params.LastBlockHeight {
		return nil
	}

	// get monitoring info
	monitoringInfo, err := k.MonitoringInfo.Get(ctx)
	if errors.Is(err, collections.ErrNotFound) {
		monitoringInfo = types.MonitoringInfo{
			SignatureCounts: networktypes.NewSignatureCounts(),
		}
	} else if err != nil {
		return err
	}

	// update signatures with voters that signed blocks
	valSetSize := lastCommit.Votes().Len()
	for i := 0; i < valSetSize; i++ {
		vote := lastCommit.Votes().Get(i)
		if vote.GetBlockIDFlag() != comet.BlockIDFlagAbsent {
			// get the operator address from the consensus address
			val, err := k.stakingKeeper.GetValidatorByConsAddr(ctx, vote.Validator().Address())
			if err != nil {
				return fmt.Errorf("validator from consensus address %s not found", vote.Validator().Address())
			}

			monitoringInfo.SignatureCounts.AddSignature(val.OperatorAddress, int64(valSetSize))
		}
	}

	// increment block count and save the monitoring info
	monitoringInfo.SignatureCounts.BlockCount++
	return k.MonitoringInfo.Set(ctx, monitoringInfo)
}

// TransmitSignatures transmits over IBC the signatures to consumer if height is reached
// and signatures are not yet transmitted
func (k Keeper) TransmitSignatures(ctx context.Context, blockHeight int64) (sequence uint64, err error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// check condition to transmit packet
	// IBC connection to consumer must be established
	// last block height must be reached
	// monitoring info must exist
	// signatures must not yet be transmitted
	params, err := k.Params.Get(ctx)
	if err != nil {
		return 0, err
	}
	if blockHeight < params.LastBlockHeight {
		return 0, nil
	}
	cid, err := k.ConnectionChannelID.Get(ctx)
	if err != nil {
		return 0, nil
	}
	mi, err := k.MonitoringInfo.Get(ctx)
	if err != nil || mi.Transmitted {
		return 0, nil
	}

	// transmit signature packet
	sequence, err = k.TransmitMonitoringPacket(
		sdkCtx,
		networktypes.MonitoringPacket{
			BlockHeight:     blockHeight,
			SignatureCounts: mi.SignatureCounts,
		},
		types.PortID,
		cid.ChannelId,
		clienttypes.ZeroHeight(),
		uint64(sdkCtx.BlockTime().Add(MonitoringPacketTimeoutDelay).UnixNano()),
	)
	if err != nil {
		if err := k.ConsumerClientID.Set(ctx, types.ConsumerClientID{
			ClientId: err.Error(),
		}); err != nil {
			return 0, err
		}
		return 0, err
	}

	// signatures have been transmitted
	mi.Transmitted = true
	return sequence, k.MonitoringInfo.Set(ctx, mi)
}
