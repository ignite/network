package types

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/pkg/chainid"
)

func NewMsgCreateChain(
	coordinator,
	genesisChainID,
	sourceUrl,
	sourceHash string,
	initialGenesis InitialGenesis,
	hasProject bool,
	projectID uint64,
	accountBalance sdk.Coins,
	metadata []byte,
) *MsgCreateChain {
	return &MsgCreateChain{
		Coordinator:    coordinator,
		GenesisChainID: genesisChainID,
		SourceURL:      sourceUrl,
		SourceHash:     sourceHash,
		InitialGenesis: initialGenesis,
		HasProject:     hasProject,
		ProjectID:      projectID,
		AccountBalance: accountBalance,
		Metadata:       metadata,
	}
}

func (msg MsgCreateChain) Type() string {
	return sdk.MsgTypeURL(&MsgCreateChain{})
}

func (msg *MsgCreateChain) ValidateBasic() error {
	if _, _, err := chainid.ParseGenesisChainID(msg.GenesisChainID); err != nil {
		return sdkerrors.Wrapf(ErrInvalidGenesisChainID, err.Error())
	}

	if err := msg.InitialGenesis.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidInitialGenesis, err.Error())
	}

	// Coins must be valid
	if !msg.AccountBalance.IsValid() {
		return sdkerrors.Wrap(ErrInvalidCoins, "default account balance sdk.Coins is not valid")
	}

	return nil
}

func NewMsgEditChain(
	coordinator string,
	launchID uint64,
	setProjectID bool,
	projectID uint64,
	metadata []byte,
) *MsgEditChain {
	return &MsgEditChain{
		Coordinator:  coordinator,
		LaunchID:     launchID,
		SetProjectID: setProjectID,
		ProjectID:    projectID,
		Metadata:     metadata,
	}
}

func (msg MsgEditChain) Type() string {
	return sdk.MsgTypeURL(&MsgEditChain{})
}

func (msg *MsgEditChain) ValidateBasic() error {
	if len(msg.Metadata) == 0 && !msg.SetProjectID {
		return sdkerrors.Wrap(ErrCannotUpdateChain, "no value to edit")
	}

	return nil
}

func NewMsgUpdateLaunchInformation(
	coordinator string,
	launchID uint64,
	genesisChainID,
	sourceURL,
	sourceHash string,
	initialGenesis *InitialGenesis,
) *MsgUpdateLaunchInformation {
	return &MsgUpdateLaunchInformation{
		Coordinator:    coordinator,
		LaunchID:       launchID,
		GenesisChainID: genesisChainID,
		SourceURL:      sourceURL,
		SourceHash:     sourceHash,
		InitialGenesis: initialGenesis,
	}
}

func (msg MsgUpdateLaunchInformation) Type() string {
	return sdk.MsgTypeURL(&MsgUpdateLaunchInformation{})
}

func (msg *MsgUpdateLaunchInformation) ValidateBasic() error {
	if msg.GenesisChainID != "" {
		if _, _, err := chainid.ParseGenesisChainID(msg.GenesisChainID); err != nil {
			return sdkerrors.Wrapf(ErrInvalidGenesisChainID, err.Error())
		}
	}

	if msg.GenesisChainID == "" && msg.SourceURL == "" && msg.InitialGenesis == nil {
		return sdkerrors.Wrap(ErrCannotUpdateChain, "no value to edit")
	}

	if msg.InitialGenesis != nil {
		if err := msg.InitialGenesis.Validate(); err != nil {
			return sdkerrors.Wrap(ErrInvalidInitialGenesis, err.Error())
		}
	}

	return nil
}

func NewMsgSendRequest(creator string, launchID uint64, content RequestContent) *MsgSendRequest {
	return &MsgSendRequest{
		Creator:  creator,
		LaunchID: launchID,
		Content:  content,
	}
}

func (msg MsgSendRequest) Type() string {
	return sdk.MsgTypeURL(&MsgSendRequest{})
}

func (msg *MsgSendRequest) ValidateBasic() error {
	if err := msg.Content.Validate(msg.LaunchID); err != nil {
		return sdkerrors.Wrapf(ErrInvalidRequestContent, err.Error())
	}
	return nil
}

func NewMsgSettleRequest(signer string, launchID uint64, requestID uint64, approve bool) *MsgSettleRequest {
	return &MsgSettleRequest{
		Signer:    signer,
		LaunchID:  launchID,
		RequestID: requestID,
		Approve:   approve,
	}
}

func (msg MsgSettleRequest) Type() string {
	return sdk.MsgTypeURL(&MsgSettleRequest{})
}

func NewMsgTriggerLaunch(coordinator string, launchID uint64, launchTime time.Time) *MsgTriggerLaunch {
	return &MsgTriggerLaunch{
		Coordinator: coordinator,
		LaunchID:    launchID,
		LaunchTime:  launchTime,
	}
}

func (msg MsgTriggerLaunch) Type() string {
	return sdk.MsgTypeURL(&MsgTriggerLaunch{})
}

func NewMsgRevertLaunch(coordinator string, launchID uint64) *MsgRevertLaunch {
	return &MsgRevertLaunch{
		Coordinator: coordinator,
		LaunchID:    launchID,
	}
}

func (msg MsgRevertLaunch) Type() string {
	return sdk.MsgTypeURL(&MsgRevertLaunch{})
}