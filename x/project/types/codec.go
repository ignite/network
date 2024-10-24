package types

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEditProject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateTotalSupply{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateSpecialAllocations{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitializeMainnet{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintVouchers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnVouchers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRedeemVouchers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnredeemVouchers{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
