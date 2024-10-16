package simulation

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/ignite/network/x/monitoringc/keeper"
)

// FindAccount find a specific address from an account list
func FindAccount(k keeper.Keeper, accs []simtypes.Account, address string) (simtypes.Account, error) {
	creator, err := k.AddressCodec().StringToBytes(address)
	if err != nil {
		return simtypes.Account{}, err
	}
	simAccount, found := simtypes.FindAccount(accs, sdk.AccAddress(creator))
	if !found {
		return simAccount, fmt.Errorf("address %s not found in the sim accounts", address)
	}
	return simAccount, nil
}
