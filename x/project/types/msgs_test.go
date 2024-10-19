package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/types"
)

func TestMsgBurnVouchers_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}
	tests := []struct {
		name string
		msg  types.MsgBurnVouchers
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers:  sample.Vouchers(r, 0),
			},
		},
		{
			name: "should prevent validation of msg with invalid vouchers",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers:  invalidCoins,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "should prevent validation of msg with empty vouchers",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers:  sdk.Coins{},
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "should prevent validation of msg with vouchers not matching project",
			msg: types.MsgBurnVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Vouchers: sdk.NewCoins(
					sdk.NewCoin("invalid/foo", sdkmath.NewInt(100)),
				),
			},
			err: types.ErrNoMatchVouchers,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgCreateProject_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgCreateProject
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgCreateProject{
				Coordinator: sample.Address(r),
				ProjectName: sample.ProjectName(r),
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should prevent validation of msg with invalid project name",
			msg: types.MsgCreateProject{
				Coordinator: sample.Address(r),
				ProjectName: invalidProjectName,
				TotalSupply: sample.TotalSupply(r),
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrInvalidProjectName,
		},
		{
			name: "should prevent validation of msg with invalid total supply",
			msg: types.MsgCreateProject{
				Coordinator: sample.Address(r),
				ProjectName: sample.ProjectName(r),
				TotalSupply: invalidCoins,
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgEditProject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgEditProject
		err  error
	}{
		{
			name: "should allow validation of msg with both name and metadata modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        sample.ProjectName(r),
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should allow validation of msg with name modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        sample.ProjectName(r),
				Metadata:    []byte{},
			},
		},
		{
			name: "should allow validation of msg with metadata modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        "",
				Metadata:    sample.Metadata(r, 20),
			},
		},
		{
			name: "should prevent validation of msg with invalid project name",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        invalidProjectName,
				Metadata:    sample.Metadata(r, 20),
			},
			err: types.ErrInvalidProjectName,
		},
		{
			name: "should prevent validation of msg with no fields modified",
			msg: types.MsgEditProject{
				ProjectID:   0,
				Coordinator: sample.Address(r),
				Name:        "",
				Metadata:    []byte{},
			},
			err: types.ErrCannotUpdateProject,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgInitializeMainnet_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgInitializeMainnet
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				ProjectID:      sample.Uint64(r),
				SourceURL:      sample.String(r, 30),
				SourceHash:     sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
		},
		{
			name: "should prevent validation of msg with empty source URL",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				ProjectID:      sample.Uint64(r),
				SourceURL:      "",
				SourceHash:     sample.String(r, 20),
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrInvalidMainnetInfo,
		},
		{
			name: "should prevent validation of msg with empty source hash",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				ProjectID:      sample.Uint64(r),
				SourceURL:      sample.String(r, 30),
				SourceHash:     "",
				MainnetChainID: sample.GenesisChainID(r),
			},
			err: types.ErrInvalidMainnetInfo,
		},
		{
			name: "should prevent validation of msg with invalid chain id",
			msg: types.MsgInitializeMainnet{
				Coordinator:    sample.Address(r),
				ProjectID:      sample.Uint64(r),
				SourceURL:      sample.String(r, 30),
				SourceHash:     sample.String(r, 20),
				MainnetChainID: "invalid_chain_id",
			},
			err: types.ErrInvalidMainnetInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgMintVouchers_ValidateBasic(t *testing.T) {
	invalidShares := types.Shares{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgMintVouchers
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectID:   0,
				Shares:      sample.Shares(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid shares",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectID:   0,
				Shares:      invalidShares,
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "should prevent validation of msg with empty shares",
			msg: types.MsgMintVouchers{
				Coordinator: sample.Address(r),
				ProjectID:   0,
				Shares:      types.EmptyShares(),
			},
			err: types.ErrInvalidShares,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgRedeemVouchers_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	addr := sample.Address(r)
	tests := []struct {
		name string
		msg  types.MsgRedeemVouchers
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgRedeemVouchers{
				Sender:    sample.Address(r),
				Account:   sample.Address(r),
				Vouchers:  sample.Vouchers(r, 0),
				ProjectID: 0,
			},
		},
		{
			name: "should allow validation of valid msg with same account and sender",
			msg: types.MsgRedeemVouchers{
				Sender:    addr,
				Account:   addr,
				Vouchers:  sample.Vouchers(r, 0),
				ProjectID: 0,
			},
		},
		{
			name: "should prevent validation of msg with invalid coin voucher",
			msg: types.MsgRedeemVouchers{
				Sender:    sample.Address(r),
				Account:   sample.Address(r),
				Vouchers:  invalidCoins,
				ProjectID: 0,
			},
			err: types.ErrInvalidVouchers,
		},
		{
			name: "should prevent validation of msg with vouchers not matching project",
			msg: types.MsgRedeemVouchers{
				Sender:    sample.Address(r),
				Account:   sample.Address(r),
				Vouchers:  sample.Vouchers(r, 10),
				ProjectID: 0,
			},
			err: types.ErrNoMatchVouchers,
		},
		{
			name: "should prevent validation of msg with invalid voucher prefix",
			msg: types.MsgRedeemVouchers{
				Sender:  sample.Address(r),
				Account: sample.Address(r),
				Vouchers: sdk.NewCoins(
					sdk.NewCoin("invalid/foo", sdkmath.NewInt(100)),
				),
				ProjectID: 0,
			},
			err: types.ErrNoMatchVouchers,
		},
		{
			name: "should prevent validation of msg with empty vouchers",
			msg: types.MsgRedeemVouchers{
				Sender:    sample.Address(r),
				Account:   sample.Address(r),
				Vouchers:  sdk.Coins{},
				ProjectID: 0,
			},
			err: types.ErrInvalidVouchers,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUnredeemVouchers_ValidateBasic(t *testing.T) {
	invalidShares := types.Shares{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgUnredeemVouchers
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgUnredeemVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Shares:    sample.Shares(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid shares",
			msg: types.MsgUnredeemVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Shares:    invalidShares,
			},
			err: types.ErrInvalidShares,
		},
		{
			name: "should prevent validation of msg with empty shares",
			msg: types.MsgUnredeemVouchers{
				Sender:    sample.Address(r),
				ProjectID: 0,
				Shares:    types.EmptyShares(),
			},
			err: types.ErrInvalidShares,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateSpecialAllocations_ValidateBasic(t *testing.T) {
	invalidShares := types.Shares{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgUpdateSpecialAllocations
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgUpdateSpecialAllocations{
				Coordinator:        sample.Address(r),
				ProjectID:          1,
				SpecialAllocations: sample.SpecialAllocations(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid special allocations",
			msg: types.MsgUpdateSpecialAllocations{
				Coordinator:        sample.Address(r),
				ProjectID:          1,
				SpecialAllocations: types.NewSpecialAllocations(invalidShares, sample.Shares(r)),
			},
			err: types.ErrInvalidSpecialAllocations,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateTotalSupply_ValidateBasic(t *testing.T) {
	invalidCoins := sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	tests := []struct {
		name string
		msg  types.MsgUpdateTotalSupply
		err  error
	}{
		{
			name: "should allow validation of valid msg",
			msg: types.MsgUpdateTotalSupply{
				Coordinator:       sample.Address(r),
				ProjectID:         0,
				TotalSupplyUpdate: sample.TotalSupply(r),
			},
		},
		{
			name: "should prevent validation of msg with invalid total supply",
			msg: types.MsgUpdateTotalSupply{
				Coordinator:       sample.Address(r),
				ProjectID:         0,
				TotalSupplyUpdate: invalidCoins,
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "should prevent validation of msg with empty total supply",
			msg: types.MsgUpdateTotalSupply{
				Coordinator:       sample.Address(r),
				ProjectID:         0,
				TotalSupplyUpdate: sdk.NewCoins(),
			},
			err: types.ErrInvalidTotalSupply,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
