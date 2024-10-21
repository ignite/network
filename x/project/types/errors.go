package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/project module sentinel errors
var (
	ErrInvalidSigner             = sdkerrors.Register(ModuleName, 1101, "expected gov account as only signer for proposal message")
	ErrInvalidTotalSupply        = sdkerrors.Register(ModuleName, 1102, "invalid total supply")
	ErrProjectNotFound           = sdkerrors.Register(ModuleName, 1103, "project not found")
	ErrMainnetInitialized        = sdkerrors.Register(ModuleName, 1104, "mainnet initialized")
	ErrInvalidShares             = sdkerrors.Register(ModuleName, 1105, "invalid shares")
	ErrTotalSharesLimit          = sdkerrors.Register(ModuleName, 1106, "allocated shares greater than total shares")
	ErrAccountNotFound           = sdkerrors.Register(ModuleName, 1107, "account not found")
	ErrSharesDecrease            = sdkerrors.Register(ModuleName, 1108, "shares can't be decreased")
	ErrVouchersMinting           = sdkerrors.Register(ModuleName, 1109, "vouchers can't be minted")
	ErrInvalidVouchers           = sdkerrors.Register(ModuleName, 1110, "invalid vouchers")
	ErrNoMatchVouchers           = sdkerrors.Register(ModuleName, 1111, "vouchers don't match to project")
	ErrInsufficientVouchers      = sdkerrors.Register(ModuleName, 1112, "account with insufficient vouchers")
	ErrInvalidProjectName        = sdkerrors.Register(ModuleName, 1113, "invalid project name")
	ErrInvalidSupplyRange        = sdkerrors.Register(ModuleName, 1114, "invalid total supply range")
	ErrInvalidMetadataLength     = sdkerrors.Register(ModuleName, 1115, "metadata field too long")
	ErrMainnetLaunchTriggered    = sdkerrors.Register(ModuleName, 1116, "mainnet launch already triggered")
	ErrInvalidSpecialAllocations = sdkerrors.Register(ModuleName, 1117, "invalid special allocations")
	ErrInvalidMainnetInfo        = sdkerrors.Register(ModuleName, 1118, "invalid mainnet info")
	ErrCannotUpdateProject       = sdkerrors.Register(ModuleName, 1119, "cannot update project")
	ErrFundCommunityPool         = sdkerrors.Register(ModuleName, 1120, "unable to fund community pool")
	ErrProjectChainsNotFound     = sdkerrors.Register(ModuleName, 1121, "project chains not found")
	ErrMainnetAccountNotFound    = sdkerrors.Register(ModuleName, 1122, "mainnet account not found")
)
