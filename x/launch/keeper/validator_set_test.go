package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	sdkmath "cosmossdk.io/math"
	"github.com/cometbft/cometbft/crypto"
	tmtypes "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/types"
)

func TestKeeper_CheckValidatorSet(t *testing.T) {
	var (
		ctx, tk, _           = testkeeper.NewTestSetup(t)
		validators           = []crypto.PubKey{sample.PubKey(r), sample.PubKey(r), sample.PubKey(r)}
		validatorSet         = tmtypes.ValidatorSet{}
		validatorNotFoundSet = tmtypes.ValidatorSet{}
		invalidValidatorSet  = tmtypes.ValidatorSet{}
	)
	notTriggeredLaunchID, err := tk.LaunchKeeper.AppendChain(ctx, types.Chain{
		CoordinatorId:   0,
		LaunchTriggered: false,
		GenesisChainId:  "spn-1",
	})
	require.NoError(t, err)
	invalidChainIDLaunchID, err := tk.LaunchKeeper.AppendChain(ctx, types.Chain{
		CoordinatorId:   0,
		LaunchTriggered: true,
		GenesisChainId:  "spn-10",
	})
	require.NoError(t, err)
	monitoringConnectedLaunchID, err := tk.LaunchKeeper.AppendChain(ctx, types.Chain{
		CoordinatorId:       0,
		LaunchTriggered:     true,
		GenesisChainId:      "spn-1",
		MonitoringConnected: true,
	})
	require.NoError(t, err)
	launchID, err := tk.LaunchKeeper.AppendChain(ctx, types.Chain{
		CoordinatorId:   0,
		LaunchTriggered: true,
		GenesisChainId:  "spn-1",
	})
	require.NoError(t, err)

	for _, validator := range validators {
		addr := sdk.AccAddress(validator.Address().Bytes())
		err = tk.LaunchKeeper.GenesisValidator.Set(ctx, collections.Join(launchID, addr), types.GenesisValidator{
			LaunchId:       launchID,
			Address:        addr.String(),
			ConsPubKey:     validator.Bytes(),
			SelfDelegation: sdk.NewCoin("spn", sdkmath.NewInt(1000)),
		})
		require.NoError(t, err)
		validatorSet.Validators = append(validatorSet.Validators,
			tmtypes.NewValidator(validator, 0),
		)
	}
	validatorNotFoundSet.Validators = append(
		validatorSet.Validators,
		tmtypes.NewValidator(sample.PubKey(r), 0),
	)
	invalidValidatorSet.Validators = validatorSet.Validators[:1]
	type args struct {
		launchID     uint64
		chainID      string
		validatorSet tmtypes.ValidatorSet
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should prevent validate validator set for non existing chain",
			args: args{
				launchID:     999,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent validate validator set for chain with launch not triggered",
			args: args{
				launchID:     notTriggeredLaunchID,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
			err: types.ErrNotTriggeredLaunch,
		},
		{
			name: "should prevent validate validator set for chain with invalid genesis chain id",
			args: args{
				launchID:     invalidChainIDLaunchID,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
			err: types.ErrInvalidGenesisChainID,
		},
		{
			name: "should prevent validate validator set for chain with monitoring already connected",
			args: args{
				launchID:     monitoringConnectedLaunchID,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
			err: types.ErrChainMonitoringConnected,
		},
		{
			name: "should prevent validate validator set if a validator is not found",
			args: args{
				launchID:     launchID,
				chainID:      "spn-1",
				validatorSet: validatorNotFoundSet,
			},
			err: types.ErrValidatorNotFound,
		},
		{
			name: "should prevent validate validator set if the minimum self delegation total is not reached",
			args: args{
				launchID:     launchID,
				chainID:      "spn-1",
				validatorSet: invalidValidatorSet,
			},
			err: types.ErrMinSelfDelegationNotReached,
		},
		{
			name: "should allow validate valid validator set",
			args: args{
				launchID:     launchID,
				chainID:      "spn-1",
				validatorSet: validatorSet,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tk.LaunchKeeper.CheckValidatorSet(ctx, tt.args.launchID, tt.args.chainID, tt.args.validatorSet)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
