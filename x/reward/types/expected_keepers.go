package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	launchtypes "github.com/ignite/network/x/launch/types"
	profiletypes "github.com/ignite/network/x/profile/types"
)

type ProfileKeeper interface {
	GetValidator(ctx context.Context, address string) (profiletypes.Validator, error)
	GetValidatorByOperatorAddress(ctx context.Context, operatorAddress string) (profiletypes.ValidatorByOperatorAddress, error)
	GetCoordinator(ctx context.Context, coordinatorID uint64) (profiletypes.Coordinator, error)
	CoordinatorIDFromAddress(ctx context.Context, address sdk.AccAddress) (uint64, error)
}

type LaunchKeeper interface {
	GetChain(ctx context.Context, launchID uint64) (launchtypes.Chain, error)
	Chains(ctx context.Context) ([]launchtypes.Chain, error)
}

// AccountKeeper defines the expected account keeper used for simulations
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
