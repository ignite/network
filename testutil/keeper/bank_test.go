package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
)

func TestTestKeepers_Mint(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	address := sample.Address(r)
	coins, otherCoins := sample.Coins(r), sample.Coins(r)

	getBalances := func(address string) sdk.Coins {
		res, err := tk.BankKeeper.AllBalances(ctx, &banktypes.QueryAllBalancesRequest{
			Address: address,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		return res.Balances
	}

	// should create the account
	tk.Mint(ctx, address, coins)
	require.True(t, getBalances(address).Equal(coins))

	// should add the minted coins in the balance
	previousBalance := getBalances(address)
	tk.Mint(ctx, address, otherCoins)
	require.True(t, getBalances(address).Equal(previousBalance.Add(otherCoins...)))
}
