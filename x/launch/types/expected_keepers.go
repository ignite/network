package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	profiletypes "github.com/ignite/network/x/profile/types"
	projecttypes "github.com/ignite/network/x/project/types"
)

type ProjectKeeper interface {
	AddChainToProject(ctx context.Context, projectID, launchID uint64) error
	GetProject(ctx context.Context, projectID uint64) (projecttypes.Project, error)
	Projects(ctx context.Context) ([]projecttypes.Project, error)
	GetProjectChains(ctx context.Context, projectID uint64) (projecttypes.ProjectChains, error)
	ListMainnetAccountBalance(ctx context.Context, projectID uint64) ([]projecttypes.MainnetAccountBalance, error)
	MainnetAccountBalance(ctx context.Context, projectID uint64, address sdk.AccAddress) (projecttypes.MainnetAccountBalance, error)
}

type MonitoringConsumerKeeper interface {
	ClearVerifiedClientIDs(ctx context.Context, launchID uint64) error
}

type ProfileKeeper interface {
	CoordinatorIDFromAddress(ctx context.Context, address sdk.AccAddress) (uint64, error)
	GetCoordinator(ctx context.Context, coordinatorID uint64) (profiletypes.Coordinator, error)
}

type DistributionKeeper interface {
	FundCommunityPool(ctx context.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	ValidateBalance(ctx context.Context, addr sdk.AccAddress) error
	HasBalance(ctx context.Context, addr sdk.AccAddress, amt sdk.Coin) bool
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx context.Context) []banktypes.Balance
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
