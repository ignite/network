package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewMsgParticipate(participant string, auctionID uint64, tierID uint64) *MsgParticipate {
	return &MsgParticipate{
		Participant: participant,
		AuctionId:   auctionID,
		TierId:      tierID,
	}
}

func (msg MsgParticipate) Type() string {
	return sdk.MsgTypeURL(&MsgParticipate{})
}

func NewMsgWithdrawAllocations(participant string, auctionID uint64) *MsgWithdrawAllocations {
	return &MsgWithdrawAllocations{
		Participant: participant,
		AuctionId:   auctionID,
	}
}

func (msg MsgWithdrawAllocations) Type() string {
	return sdk.MsgTypeURL(&MsgWithdrawAllocations{})
}
