package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	fundraisingtypes "github.com/ignite/modules/x/fundraising/types"
)

// StakingKeeper defines the expected interface for the Staking module.
type StakingKeeper interface {
	GetDelegatorDelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) ([]stakingtypes.Delegation, error)
	ValidatorByConsAddr(ctx context.Context, addr sdk.ConsAddress) (stakingtypes.ValidatorI, error)
}

type FundraisingKeeper interface {
	Auctions(ctx context.Context) ([]fundraisingtypes.AuctionI, error)
	GetAuction(ctx context.Context, auctionID uint64) (fundraisingtypes.AuctionI, error)
	AddAllowedBidders(ctx context.Context, auctionID uint64, allowedBidders []fundraisingtypes.AllowedBidder) error
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
