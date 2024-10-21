package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/pkg/chainid"
)

func NewMsgCreateProject(coordinator string, projectName string, totalSupply sdk.Coins, metadata []byte) *MsgCreateProject {
	return &MsgCreateProject{
		Coordinator: coordinator,
		ProjectName: projectName,
		TotalSupply: totalSupply,
		Metadata:    metadata,
	}
}

func (msg MsgCreateProject) Type() string {
	return sdk.MsgTypeURL(&MsgCreateProject{})
}

func (msg *MsgCreateProject) ValidateBasic() error {
	if err := CheckProjectName(msg.ProjectName); err != nil {
		return sdkerrors.Wrap(ErrInvalidProjectName, err.Error())
	}

	if !msg.TotalSupply.IsValid() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is not a valid Coins object")
	}

	return nil
}

func NewMsgEditProject(coordinator string, projectID uint64, name string, metadata []byte) *MsgEditProject {
	return &MsgEditProject{
		Coordinator: coordinator,
		ProjectId:   projectID,
		Name:        name,
		Metadata:    metadata,
	}
}

func (msg MsgEditProject) Type() string {
	return sdk.MsgTypeURL(&MsgEditProject{})
}

func (msg *MsgEditProject) ValidateBasic() error {
	if len(msg.Name) == 0 && len(msg.Metadata) == 0 {
		return sdkerrors.Wrap(ErrCannotUpdateProject, "must modify at least one field (name or metadata)")
	}

	if len(msg.Name) != 0 {
		if err := CheckProjectName(msg.Name); err != nil {
			return sdkerrors.Wrap(ErrInvalidProjectName, err.Error())
		}
	}

	return nil
}

func NewMsgUpdateTotalSupply(coordinator string, projectID uint64, totalSupplyUpdate sdk.Coins) *MsgUpdateTotalSupply {
	return &MsgUpdateTotalSupply{
		Coordinator:       coordinator,
		ProjectId:         projectID,
		TotalSupplyUpdate: totalSupplyUpdate,
	}
}

func (msg MsgUpdateTotalSupply) Type() string {
	return sdk.MsgTypeURL(&MsgUpdateTotalSupply{})
}

func (msg *MsgUpdateTotalSupply) ValidateBasic() error {
	if !msg.TotalSupplyUpdate.IsValid() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is not a valid Coins object")
	}

	if msg.TotalSupplyUpdate.Empty() {
		return sdkerrors.Wrap(ErrInvalidTotalSupply, "total supply is empty")
	}

	return nil
}

func NewMsgUpdateSpecialAllocations(coordinator string, projectID uint64, specialAllocations SpecialAllocations) *MsgUpdateSpecialAllocations {
	return &MsgUpdateSpecialAllocations{
		Coordinator:        coordinator,
		ProjectId:          projectID,
		SpecialAllocations: specialAllocations,
	}
}

func (msg MsgUpdateSpecialAllocations) Type() string {
	return sdk.MsgTypeURL(&MsgUpdateSpecialAllocations{})
}

func (msg *MsgUpdateSpecialAllocations) ValidateBasic() error {
	if err := msg.SpecialAllocations.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidSpecialAllocations, err.Error())
	}
	return nil
}

func NewMsgInitializeMainnet(coordinator string, projectID uint64, sourceURL string, sourceHash string, mainnetChainID string) *MsgInitializeMainnet {
	return &MsgInitializeMainnet{
		Coordinator:    coordinator,
		ProjectId:      projectID,
		SourceUrl:      sourceURL,
		SourceHash:     sourceHash,
		MainnetChainId: mainnetChainID,
	}
}

func (msg MsgInitializeMainnet) Type() string {
	return sdk.MsgTypeURL(&MsgInitializeMainnet{})
}

func (msg *MsgInitializeMainnet) ValidateBasic() error {
	if msg.SourceUrl == "" {
		return sdkerrors.Wrap(ErrInvalidMainnetInfo, "empty source URL")
	}
	if msg.SourceHash == "" {
		return sdkerrors.Wrap(ErrInvalidMainnetInfo, "empty source hash")
	}
	if _, _, err := chainid.ParseGenesisChainID(msg.MainnetChainId); err != nil {
		return sdkerrors.Wrapf(ErrInvalidMainnetInfo, err.Error())
	}

	return nil
}

func NewMsgMintVouchers(coordinator string, projectID uint64, shares Shares) *MsgMintVouchers {
	return &MsgMintVouchers{
		Coordinator: coordinator,
		ProjectId:   projectID,
		Shares:      shares,
	}
}

func (msg MsgMintVouchers) Type() string {
	return sdk.MsgTypeURL(&MsgMintVouchers{})
}

func (msg *MsgMintVouchers) ValidateBasic() error {
	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidShares, sdk.Coins(msg.Shares).String())
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidShares, "shares is empty")
	}

	return nil
}

func NewMsgBurnVouchers(sender string, projectID uint64, vouchers sdk.Coins) *MsgBurnVouchers {
	return &MsgBurnVouchers{
		Sender:    sender,
		ProjectId: projectID,
		Vouchers:  vouchers,
	}
}

func (msg MsgBurnVouchers) Type() string {
	return sdk.MsgTypeURL(&MsgBurnVouchers{})
}

func (msg *MsgBurnVouchers) ValidateBasic() error {
	if !msg.Vouchers.IsValid() {
		return sdkerrors.Wrap(ErrInvalidVouchers, msg.Vouchers.String())
	}

	if msg.Vouchers.Empty() {
		return sdkerrors.Wrap(ErrInvalidVouchers, "vouchers is empty")
	}

	if err := CheckVouchers(msg.Vouchers, msg.ProjectId); err != nil {
		return sdkerrors.Wrap(ErrNoMatchVouchers, err.Error())
	}
	return nil
}

func NewMsgRedeemVouchers(sender string, projectID uint64, account string, vouchers sdk.Coins) *MsgRedeemVouchers {
	return &MsgRedeemVouchers{
		Sender:    sender,
		ProjectId: projectID,
		Account:   account,
		Vouchers:  vouchers,
	}
}

func (msg MsgRedeemVouchers) Type() string {
	return sdk.MsgTypeURL(&MsgRedeemVouchers{})
}

func (msg *MsgRedeemVouchers) ValidateBasic() error {
	if !msg.Vouchers.IsValid() {
		return sdkerrors.Wrap(ErrInvalidVouchers, msg.Vouchers.String())
	}

	if msg.Vouchers.Empty() {
		return sdkerrors.Wrap(ErrInvalidVouchers, "vouchers is empty")
	}

	if err := CheckVouchers(msg.Vouchers, msg.ProjectId); err != nil {
		return sdkerrors.Wrap(ErrNoMatchVouchers, err.Error())
	}
	return nil
}

func NewMsgUnredeemVouchers(sender string, projectID uint64, shares Shares) *MsgUnredeemVouchers {
	return &MsgUnredeemVouchers{
		Sender:    sender,
		ProjectId: projectID,
		Shares:    shares,
	}
}

func (msg MsgUnredeemVouchers) Type() string {
	return sdk.MsgTypeURL(&MsgUnredeemVouchers{})
}

func (msg *MsgUnredeemVouchers) ValidateBasic() error {
	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidShares, sdk.Coins(msg.Shares).String())
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidShares, "shares is empty")
	}

	return nil
}
