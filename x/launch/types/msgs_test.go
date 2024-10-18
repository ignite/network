package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
)

func TestMsgCreateChain_ValidateBasic(t *testing.T) {
	invalidGenesisHash := sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0)
	invalidInitialGenesis := types.NewGenesisURL("foo.com", "NoHash")
	invalidGenesisHash.InitialGenesis = invalidInitialGenesis

	invalidConfigFile := sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0)
	invalidInitialGenesis = types.NewGenesisConfig("")
	invalidConfigFile.InitialGenesis = invalidInitialGenesis

	invalidGenesisChainID := sample.MsgCreateChain(r, sample.Address(r), "", false, 0)
	invalidGenesisChainID.GenesisChainID = "invalid"

	msgInvalidCoins := sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0)
	msgInvalidCoins.AccountBalance = sdk.Coins{sdk.Coin{Denom: "invalid", Amount: sdkmath.NewInt(-1)}}

	for _, tc := range []struct {
		desc  string
		msg   types.MsgCreateChain
		valid bool
	}{
		{
			desc:  "should validate valid message",
			msg:   sample.MsgCreateChain(r, sample.Address(r), "", false, 0),
			valid: true,
		},
		{
			desc:  "should validate valid message with genesis URL",
			msg:   sample.MsgCreateChain(r, sample.Address(r), "foo.com", false, 0),
			valid: true,
		},
		{
			desc:  "should prevent validate message with invalid genesis hash for custom genesis",
			msg:   invalidGenesisHash,
			valid: false,
		},
		{
			desc:  "should prevent validate message with invalid file for GenesisConfig custom genesis",
			msg:   invalidConfigFile,
			valid: false,
		},
		{
			desc:  "should prevent validate message with invalid genesis chain ID",
			msg:   invalidGenesisChainID,
			valid: false,
		},
		{
			desc:  "should prevent chain with invalid coins structure",
			msg:   msgInvalidCoins,
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgEditChain_ValidateBasic(t *testing.T) {
	launchID := uint64(0)
	for _, tc := range []struct {
		desc  string
		msg   types.MsgEditChain
		valid bool
	}{
		{
			desc: "should validate valid message",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				true,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "should validate valid message with new metadata",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				false,
				0,
				true,
			),
			valid: true,
		},
		{
			desc: "should validate valid message with new chain ID",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				true,
				0,
				false,
			),
			valid: true,
		},
		{
			desc: "should prevent validate message with no value to edit",
			msg: sample.MsgEditChain(r,
				sample.Address(r),
				launchID,
				false,
				0,
				false,
			),
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestMsgSendRequest_ValidateBasic(t *testing.T) {
	launchID := sample.Uint64(r)

	tests := []struct {
		name string
		msg  types.MsgSendRequest
		err  error
	}{
		{
			name: "should validate valid message",
			msg: types.MsgSendRequest{
				Creator:  sample.Address(r),
				LaunchID: launchID,
				Content:  sample.RequestContent(r, launchID),
			},
		},
		{
			name: "should prevent validate message with invalid request content",
			msg: types.MsgSendRequest{
				Creator:  sample.Address(r),
				LaunchID: sample.Uint64(r),
				Content:  types.NewGenesisAccount(0, "", sdk.NewCoins()),
			},
			err: types.ErrInvalidRequestContent,
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

func TestMsgUpdateLaunchInformation_ValidateBasic(t *testing.T) {
	launchID := uint64(0)
	msgInvalidGenesisHash := sample.MsgUpdateLaunchInformation(r,
		sample.Address(r),
		launchID,
		false,
		true,
		false,
		false,
	)
	genesisURL := types.NewGenesisURL("foo.com", "NoHash")
	msgInvalidGenesisHash.InitialGenesis = &genesisURL

	msgInvalidGenesisChainID := sample.MsgUpdateLaunchInformation(r,
		sample.Address(r),
		launchID,
		false,
		true,
		false,
		false,
	)
	msgInvalidGenesisChainID.GenesisChainID = "invalid"

	for _, tc := range []struct {
		desc  string
		msg   types.MsgUpdateLaunchInformation
		valid bool
	}{
		{
			desc: "should validate valid message",
			msg: sample.MsgUpdateLaunchInformation(r,
				sample.Address(r),
				launchID,
				true,
				true,
				true,
				false,
			),
			valid: true,
		},
		{
			desc: "should validate valid message with new genesis chain ID",
			msg: sample.MsgUpdateLaunchInformation(r,
				sample.Address(r),
				launchID,
				true,
				false,
				false,
				false,
			),
			valid: true,
		},
		{
			desc: "should validate valid message with new source",
			msg: sample.MsgUpdateLaunchInformation(r,
				sample.Address(r),
				launchID,
				false,
				true,
				false,
				false,
			),
			valid: true,
		},
		{
			desc: "should validate valid message with new genesis",
			msg: sample.MsgUpdateLaunchInformation(r,
				sample.Address(r),
				launchID,
				false,
				false,
				true,
				false,
			),
			valid: true,
		},
		{
			desc: "should validate valid message with new genesis with a custom genesis url",
			msg: sample.MsgUpdateLaunchInformation(r,
				sample.Address(r),
				launchID,
				false,
				false,
				true,
				true,
			),
			valid: true,
		},
		{
			desc: "should prevent validate message with no value to edit",
			msg: sample.MsgUpdateLaunchInformation(r,
				sample.Address(r),
				launchID,
				false,
				false,
				false,
				false,
			),
			valid: false,
		},
		{
			desc:  "should prevent validate message with invalid initial genesis hash",
			msg:   msgInvalidGenesisHash,
			valid: false,
		},
		{
			desc:  "should prevent validate message with invalid initial genesis chain ID",
			msg:   msgInvalidGenesisChainID,
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
