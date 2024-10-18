package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEditChain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateLaunchInformation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSettleRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTriggerLaunch{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRevertLaunch{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
