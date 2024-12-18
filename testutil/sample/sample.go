// Package sample provides methods to initialize sample object of various types for test purposes
package sample

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cosmosed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibctypes "github.com/cosmos/ibc-go/v8/modules/core/types"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	claimtypes "github.com/ignite/modules/x/claim/types"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"
	minttypes "github.com/ignite/modules/x/mint/types"
	"github.com/stretchr/testify/require"

	launch "github.com/ignite/network/x/launch/types"
	monitoringc "github.com/ignite/network/x/monitoringc/types"
	monitoringp "github.com/ignite/network/x/monitoringp/types"
	participation "github.com/ignite/network/x/participation/types"
	profile "github.com/ignite/network/x/profile/types"
	project "github.com/ignite/network/x/project/types"
	reward "github.com/ignite/network/x/reward/types"
)

func InterfaceRegistry() codectypes.InterfaceRegistry {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	cryptocodec.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)
	authz.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	ibctypes.RegisterInterfaces(interfaceRegistry)
	ibcfeetypes.RegisterInterfaces(interfaceRegistry)
	ibcclienttypes.RegisterInterfaces(interfaceRegistry)
	ibctm.RegisterInterfaces(interfaceRegistry)
	ibctransfertypes.RegisterInterfaces(interfaceRegistry)
	icacontrollertypes.RegisterInterfaces(interfaceRegistry)
	icatypes.RegisterInterfaces(interfaceRegistry)
	channeltypes.RegisterInterfaces(interfaceRegistry)
	consensusparamtypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	distrtypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	claimtypes.RegisterInterfaces(interfaceRegistry)
	feegrant.RegisterInterfaces(interfaceRegistry)
	// govtypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)

	launch.RegisterInterfaces(interfaceRegistry)
	profile.RegisterInterfaces(interfaceRegistry)
	project.RegisterInterfaces(interfaceRegistry)
	monitoringc.RegisterInterfaces(interfaceRegistry)
	monitoringp.RegisterInterfaces(interfaceRegistry)
	reward.RegisterInterfaces(interfaceRegistry)
	participation.RegisterInterfaces(interfaceRegistry)
	minttypes.RegisterInterfaces(interfaceRegistry)
	fundraisingtypes.RegisterInterfaces(interfaceRegistry)

	return interfaceRegistry
}

// Codec returns a codec with preregistered interfaces
func Codec() codec.Codec {
	return codec.NewProtoCodec(InterfaceRegistry())
}

// Bool returns randomly true or false
func Bool(r *rand.Rand) bool {
	b := r.Intn(100)
	return b < 50
}

// Bytes returns a random array of bytes
func Bytes(r *rand.Rand, n int) []byte {
	return []byte(String(r, n))
}

// Uint64 returns a random uint64
func Uint64(r *rand.Rand) uint64 {
	return uint64(r.Intn(10000))
}

// String returns a random string of length n
func String(r *rand.Rand, n int) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[r.Intn(len(letter))]
	}
	return string(randomString)
}

// AlphaString returns a random string with lowercase alpha char of length n
func AlphaString(r *rand.Rand, n int) string {
	letter := []rune("abcdefghijklmnopqrstuvwxyz")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[r.Intn(len(letter))]
	}
	return string(randomString)
}

// NonAlphaString returns a random string with non alpha char of length n
func NonAlphaString(r *rand.Rand, n int) string {
	letter := []rune("0123456789!@#$%^&*()_+")

	randomString := make([]rune, n)
	for i := range randomString {
		randomString[i] = letter[r.Intn(len(letter))]
	}
	return string(randomString)
}

// PubKey returns a sample account PubKey
func PubKey(r *rand.Rand) crypto.PubKey {
	seed := []byte(strconv.Itoa(r.Int()))
	return ed25519.GenPrivKeyFromSecret(seed).PubKey()
}

// ConsAddress returns a sample consensus address
func ConsAddress(r *rand.Rand) sdk.ConsAddress {
	return sdk.ConsAddress(PubKey(r).Address())
}

// AccAddress returns a sample account address
func AccAddress(r *rand.Rand) sdk.AccAddress {
	addr := PubKey(r).Address()
	return sdk.AccAddress(addr)
}

// Address returns a sample string account address
func Address(r *rand.Rand) string {
	return AccAddress(r).String()
}

// ValAddress returns a sample validator operator address
func ValAddress(r *rand.Rand) sdk.ValAddress {
	return sdk.ValAddress(PubKey(r).Address())
}

// OperatorAddress returns a sample string validator operator address
func OperatorAddress(r *rand.Rand) string {
	return ValAddress(r).String()
}

// Validator returns a sample staking validator
func Validator(t testing.TB, r *rand.Rand) stakingtypes.Validator {
	seed := []byte(strconv.Itoa(r.Int()))
	val, err := stakingtypes.NewValidator(
		ValAddress(r).String(),
		cosmosed25519.GenPrivKeyFromSecret(seed).PubKey(),
		stakingtypes.Description{})
	require.NoError(t, err)
	return val
}

// Delegation returns staking delegation with the given address
func Delegation(t testing.TB, r *rand.Rand, addr string) stakingtypes.Delegation {
	delAcc, err := sdk.AccAddressFromBech32(addr)
	require.NoError(t, err)

	return stakingtypes.NewDelegation(
		delAcc.String(),
		ValAddress(r).String(),
		sdkmath.LegacyNewDec(int64(r.Intn(10000))),
	)
}

// Coin returns a sample coin structure
func Coin(r *rand.Rand) sdk.Coin {
	return sdk.NewCoin(AlphaString(r, 5), sdkmath.NewInt(r.Int63n(10000)+1))
}

// CoinWithRange returns a sample coin structure where the amount is a random number between provided min and max values
// with a random denom
func CoinWithRange(r *rand.Rand, min, max int64) sdk.Coin {
	return sdk.NewCoin(AlphaString(r, 5), sdkmath.NewInt(r.Int63n(max-min)+min))
}

// CoinWithRangeAmount returns a sample coin structure where the amount is a random number between provided min and max values
// with a given denom
func CoinWithRangeAmount(r *rand.Rand, denom string, min, max int64) sdk.Coin {
	return sdk.NewCoin(denom, sdkmath.NewInt(r.Int63n(max-min)+min))
}

// Coins returns a sample coins structure
func Coins(r *rand.Rand) sdk.Coins {
	return sdk.NewCoins(Coin(r), Coin(r), Coin(r))
}

// CoinsWithRange returns a sample coins structure where the amount is a random number between provided min and max values
func CoinsWithRange(r *rand.Rand, min, max int64) sdk.Coins {
	return sdk.NewCoins(CoinWithRange(r, min, max), CoinWithRange(r, min, max), CoinWithRange(r, min, max))
}

// CoinsWithRangeAmount returns a sample coins structure where the amount is a random number between provided min and max values
// with a set of given denoms
func CoinsWithRangeAmount(r *rand.Rand, denom1, denom2, denom3 string, min, max int64) sdk.Coins {
	return sdk.NewCoins(CoinWithRangeAmount(r, denom1, min, max), CoinWithRangeAmount(r, denom2, min, max), CoinWithRangeAmount(r, denom3, min, max))
}

// TotalSupply returns a sample coins structure where each denom's total supply is within the default
// allowed supply range
func TotalSupply(r *rand.Rand) sdk.Coins {
	return CoinsWithRange(r, project.DefaultMinTotalSupply.Int64(), project.DefaultMaxTotalSupply.Int64())
}

// Duration returns a sample time.Duration between a second and 21 days
func Duration(r *rand.Rand) time.Duration {
	return time.Duration(r.Int63n(int64(time.Hour*24*21-time.Second))) + time.Second
}

// DurationFromRange returns a sample time.Duration between the min and max values provided
func DurationFromRange(r *rand.Rand, min, max time.Duration) time.Duration {
	return time.Duration(r.Int63n(int64(max-min))) + min
}

// Int returns a sample sdkmath.Int
func Int(r *rand.Rand) sdkmath.Int {
	return sdkmath.NewInt(r.Int63())
}

// Time returns a sample time
func Time(r *rand.Rand) time.Time {
	return time.UnixMilli(r.Int63n(1000) + 1).UTC()
}

// ZeroTime returns time.Time that represents 0
func ZeroTime() time.Time {
	return time.UnixMilli(0).UTC()
}
