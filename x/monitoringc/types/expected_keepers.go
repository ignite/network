package types

import (
	"context"

	tmtypes "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ignite/network/pkg/types"
	launchtypes "github.com/ignite/network/x/launch/types"
)

type LaunchKeeper interface {
	GetChain(ctx context.Context, launchID uint64) (launchtypes.Chain, error)
	EnableMonitoringConnection(ctx context.Context, launchID uint64) error
	CheckValidatorSet(
		ctx context.Context,
		launchID uint64,
		chainID string,
		validatorSet tmtypes.ValidatorSet,
	) error
}

type RewardKeeper interface {
	DistributeRewards(
		ctx context.Context,
		launchID uint64,
		signatureCounts types.SignatureCounts,
		lastBlockHeight int64,
		closeRewardPool bool,
	) error
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
