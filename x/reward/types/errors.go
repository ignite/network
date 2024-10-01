package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/reward module sentinel errors
var (
	ErrInvalidSigner          = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrInvalidRewardPoolCoins = sdkerrors.Register(ModuleName, 1102, "invalid coins for reward pool")
	ErrInvalidCoordinatorID   = sdkerrors.Register(ModuleName, 1103, "invalid coordinator id for reward pool")
	ErrRewardPoolNotFound     = sdkerrors.Register(ModuleName, 1104, "reward pool not found")
	ErrRewardPoolClosed       = sdkerrors.Register(ModuleName, 1105, "reward pool is closed")
	ErrInvalidSignatureCounts = sdkerrors.Register(ModuleName, 1106, "invalid signature counts")
	ErrInvalidLastBlockHeight = sdkerrors.Register(ModuleName, 1107, "invalid last block height")
	ErrInvalidRewardHeight    = sdkerrors.Register(ModuleName, 1108, "invalid reward height")
	ErrInsufficientFunds      = sdkerrors.Register(ModuleName, 1109, "insufficient funds")
)
