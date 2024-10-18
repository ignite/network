package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/participation module sentinel errors
var (
	ErrInvalidSigner                      = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrAuctionNotFound                    = sdkerrors.Register(ModuleName, 1102, "auction not found")
	ErrInvalidBidder                      = sdkerrors.Register(ModuleName, 1103, "invalid bidder")
	ErrInvalidAllocationAmount            = sdkerrors.Register(ModuleName, 1104, "invalid allocation amount")
	ErrTierNotFound                       = sdkerrors.Register(ModuleName, 1105, "tier not found")
	ErrInsufficientAllocations            = sdkerrors.Register(ModuleName, 1106, "insufficient allocations")
	ErrAlreadyParticipating               = sdkerrors.Register(ModuleName, 1107, "address is already participating")
	ErrParticipationNotAllowed            = sdkerrors.Register(ModuleName, 1108, "unable to participate to auction")
	ErrAllocationWithdrawalTimeNotReached = sdkerrors.Register(ModuleName, 1109, "unable to withdraw allocations")
	ErrUsedAllocationsNotFound            = sdkerrors.Register(ModuleName, 1110, "used allocations not found")
	ErrAllocationsAlreadyWithdrawn        = sdkerrors.Register(ModuleName, 1111, "used allocations already withdrawn")
	ErrInvalidAddress                     = sdkerrors.Register(ModuleName, 1112, "invalid participant address")
	ErrInvalidDelegations                 = sdkerrors.Register(ModuleName, 1113, "invalid degalations")
)
