package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/launch module sentinel errors
var (
	ErrInvalidSigner               = sdkerrors.Register(ModuleName, 1102, "expected gov account as only signer for proposal message")
	ErrInvalidGenesisChainID       = sdkerrors.Register(ModuleName, 1103, "the genesis chain id is invalid")
	ErrInvalidInitialGenesis       = sdkerrors.Register(ModuleName, 1104, "the initial genesis is invalid")
	ErrTriggeredLaunch             = sdkerrors.Register(ModuleName, 1105, "launch is triggered for the chain")
	ErrNoAddressPermission         = sdkerrors.Register(ModuleName, 1106, "you must be the coordinator or address owner to perform this action")
	ErrInvalidCoins                = sdkerrors.Register(ModuleName, 1107, "the coin list is invalid")
	ErrInvalidVestingOption        = sdkerrors.Register(ModuleName, 1108, "invalid vesting option")
	ErrLaunchTimeTooLow            = sdkerrors.Register(ModuleName, 1109, "the remaining time is below authorized launch time")
	ErrNotTriggeredLaunch          = sdkerrors.Register(ModuleName, 1110, "the chain launch has not been triggered")
	ErrRevertDelayNotReached       = sdkerrors.Register(ModuleName, 1111, "the revert delay has not been reached")
	ErrRequestNotFound             = sdkerrors.Register(ModuleName, 1112, "request not found")
	ErrInvalidConsPubKey           = sdkerrors.Register(ModuleName, 1113, "the consensus public key is invalid")
	ErrInvalidGenTx                = sdkerrors.Register(ModuleName, 1114, "the gentx is invalid")
	ErrInvalidSelfDelegation       = sdkerrors.Register(ModuleName, 1115, "the self delegation is invalid")
	ErrInvalidPeer                 = sdkerrors.Register(ModuleName, 1116, "the peer is invalid")
	ErrAccountAlreadyExist         = sdkerrors.Register(ModuleName, 1117, "account already exists")
	ErrAccountNotFound             = sdkerrors.Register(ModuleName, 1118, "account not found")
	ErrValidatorAlreadyExist       = sdkerrors.Register(ModuleName, 1119, "validator already exists")
	ErrValidatorNotFound           = sdkerrors.Register(ModuleName, 1120, "validator not found")
	ErrChainInactive               = sdkerrors.Register(ModuleName, 1121, "the chain is inactive")
	ErrCreateChainFail             = sdkerrors.Register(ModuleName, 1122, "fail to create a new chain")
	ErrLaunchTimeTooHigh           = sdkerrors.Register(ModuleName, 1123, "the remaining time is above authorized launch time")
	ErrMinSelfDelegationNotReached = sdkerrors.Register(ModuleName, 1124, "the minimum self delgation is not reachead")
	ErrInvalidMetadataLength       = sdkerrors.Register(ModuleName, 1125, "metadata field too long")
	ErrChainHasProject             = sdkerrors.Register(ModuleName, 1126, "chain already is associated with a project")
	ErrAddChainToProject           = sdkerrors.Register(ModuleName, 1127, "unable to add chain to project")
	ErrChainMonitoringConnected    = sdkerrors.Register(ModuleName, 1128, "chain is already connected to monitoring")
	ErrRequestSettled              = sdkerrors.Register(ModuleName, 1129, "request is already settled")
	ErrInvalidRequestContent       = sdkerrors.Register(ModuleName, 1130, "request content is invalid")
	ErrInvalidRequestForMainnet    = sdkerrors.Register(ModuleName, 1131, "request is invalid for mainnet")
	ErrRequestApplicationFailure   = sdkerrors.Register(ModuleName, 1132, "request failed to be applied")
	ErrInvalidLaunchID             = sdkerrors.Register(ModuleName, 1133, "invalid launch ID")
	ErrCannotUpdateChain           = sdkerrors.Register(ModuleName, 1134, "cannot update chain")
	ErrFundCommunityPool           = sdkerrors.Register(ModuleName, 1135, "unable to fund community pool")
	ErrInvalidModuleName           = sdkerrors.Register(ModuleName, 1136, "invalid module name")
	ErrInvalidParamName            = sdkerrors.Register(ModuleName, 1137, "invalid param name")
	ErrChainNotFound               = sdkerrors.Register(ModuleName, 1138, "chain not found for launch id")
)
