// Package constructor provides constructors to easily initialize objects for test purpose with automatic error handling
package constructor

import (
	"testing"

	"cosmossdk.io/core/comet"
	sdkmath "cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	prototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ignite/network/pkg/types"
	monitoringptypes "github.com/ignite/network/x/monitoringp/types"
	projecttypes "github.com/ignite/network/x/project/types"
)

// Vote is a simplified type for abci.VoteInfo for testing purpose
type Vote struct {
	Address []byte
	BlockID prototypes.BlockIDFlag
}

// LastCommitInfo creates a ABCI LastCommitInfo object for test purpose from a list of vote
func LastCommitInfo(votes ...Vote) comet.CommitInfo {
	var lci CommitInfo

	// add votes
	for _, vote := range votes {
		lci.CommitInfo.Votes = append(lci.CommitInfo.Votes, abci.VoteInfo{
			Validator: abci.Validator{
				Address: vote.Address,
			},
			BlockIdFlag: vote.BlockID,
		})
	}
	return lci
}

// Coin returns a sdk.Coin from a string
func Coin(t testing.TB, str string) sdk.Coin {
	coin, err := sdk.ParseCoinNormalized(str)
	require.NoError(t, err)
	return coin
}

// Coins returns a sdk.Coins from a string
func Coins(t testing.TB, str string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(str)
	require.NoError(t, err)
	return coins
}

// Dec returns a sdk.Dec from a string
func Dec(t testing.TB, str string) sdkmath.LegacyDec {
	dec, err := sdkmath.LegacyNewDecFromStr(str)
	require.NoError(t, err)
	return dec
}

// SignatureCount returns a signature count object for test from a operator address and a decimal string for relative signatures
func SignatureCount(t testing.TB, opAddr string, relSig string) types.SignatureCount {
	return types.SignatureCount{
		OpAddress:          opAddr,
		RelativeSignatures: Dec(t, relSig),
	}
}

// SignatureCounts returns a signature counts object for tests from a a block count and list of signature counts
func SignatureCounts(blockCount uint64, sc ...types.SignatureCount) types.SignatureCounts {
	return types.SignatureCounts{
		BlockCount: blockCount,
		Counts:     sc,
	}
}

// MonitoringInfo returns a monitoring info object for tests from a a block count and list of signature counts
func MonitoringInfo(blockCount uint64, sc ...types.SignatureCount) (mi monitoringptypes.MonitoringInfo) {
	mi.SignatureCounts = SignatureCounts(blockCount, sc...)
	return
}

// Shares returns a Shares object from a string of coin inputs
func Shares(t testing.TB, coinStr string) projecttypes.Shares {
	shares := projecttypes.NewSharesFromCoins(Coins(t, coinStr))
	return shares
}

// Vouchers returns a Vouchers object from a string of coin inputs
func Vouchers(t testing.TB, coinStr string, projectID uint64) sdk.Coins {
	coins := Coins(t, coinStr)
	vouchers := make(sdk.Coins, len(coins))
	for i, coin := range coins {
		coin.Denom = projecttypes.VoucherDenom(projectID, coin.Denom)
		vouchers[i] = coin
	}
	return vouchers
}
