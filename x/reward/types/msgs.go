package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgSetRewards(provider string, launchID uint64, coins sdk.Coins, lastRewardHeight int64) *MsgSetRewards {
	return &MsgSetRewards{
		Provider:         provider,
		LaunchID:         launchID,
		Coins:            coins,
		LastRewardHeight: lastRewardHeight,
	}
}

func (msg MsgSetRewards) Type() string {
	return sdk.MsgTypeURL(&MsgSetRewards{})
}

func (msg *MsgSetRewards) ValidateBasic() error {
	if err := msg.Coins.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidRewardPoolCoins, err.Error())
	}

	if msg.LastRewardHeight < 0 {
		return sdkerrors.Wrap(ErrInvalidRewardHeight, "last reward height must be non-negative")
	}
	return nil
}
