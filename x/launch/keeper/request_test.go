package keeper_test

import (
	"context"
	"testing"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	ignterrors "github.com/ignite/network/pkg/errors"
	testkeeper "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/launch/keeper"
	"github.com/ignite/network/x/launch/types"
)

type RequestSample struct {
	Content types.RequestContent
	Creator string
	Status  types.Request_Status
}

func createRequestsFromSamples(
	k *keeper.Keeper,
	ctx context.Context,
	launchID uint64,
	samples []RequestSample,
) []types.Request {
	items := make([]types.Request, len(samples))
	for i, s := range samples {
		items[i] = sample.RequestWithContentAndCreator(r, launchID, s.Content, s.Creator)
		items[i].Status = s.Status
		items[i].RequestId, _ = k.AppendRequest(ctx, items[i])
	}
	return items
}

func TestCheckAccount(t *testing.T) {
	var (
		genesisAcc = sample.AccAddress(r)
		vestingAcc = sample.AccAddress(r)
		dupAcc     = sample.AccAddress(r)
		notFound   = sample.AccAddress(r)
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		launchID   = uint64(10)
	)

	ga := sample.GenesisAccount(
		r,
		launchID,
		genesisAcc.String(),
	)
	err := tk.LaunchKeeper.GenesisAccount.Set(ctx, collections.Join(launchID, genesisAcc), ga)
	require.NoError(t, err)

	va := sample.VestingAccount(
		r,
		launchID,
		vestingAcc.String(),
	)
	err = tk.LaunchKeeper.VestingAccount.Set(ctx, collections.Join(launchID, vestingAcc), va)
	require.NoError(t, err)

	// set duplicated entries
	ga.Address = dupAcc.String()
	va.Address = dupAcc.String()
	err = tk.LaunchKeeper.GenesisAccount.Set(ctx, collections.Join(launchID, dupAcc), ga)
	require.NoError(t, err)
	err = tk.LaunchKeeper.VestingAccount.Set(ctx, collections.Join(launchID, dupAcc), va)
	require.NoError(t, err)

	tests := []struct {
		name  string
		addr  sdk.AccAddress
		found bool
		err   error
	}{
		{
			name:  "should return false if genesis or vesting account is not found",
			addr:  notFound,
			found: false,
		}, {
			name:  "should return true if genesis account found",
			addr:  genesisAcc,
			found: true,
		}, {
			name:  "vesting return true if account found",
			addr:  vestingAcc,
			found: true,
		}, {
			name: "should return critical error if duplicated genesis and vesting accounts",
			addr: dupAcc,
			err:  ignterrors.ErrCritical,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, err := keeper.CheckAccount(ctx, tk.LaunchKeeper, launchID, tt.addr.String())
			if tt.err != nil {
				require.Error(t, err)
				require.ErrorIs(t, tt.err, err)
				return
			}

			require.Equal(t, found, tt.found)
		})
	}
}

func TestApplyRequest(t *testing.T) {
	var (
		coord          = sample.Coordinator(r, sample.Address(r))
		coordID        = uint64(3)
		genesisAcc     = sample.Address(r)
		vestingAcc     = sample.Address(r)
		validatorAcc   = sample.Address(r)
		ctx, tk, _     = testkeeper.NewTestSetup(t)
		launchID       = uint64(10)
		contents       = sample.AllRequestContents(r, launchID, genesisAcc, vestingAcc, validatorAcc)
		invalidContent = types.NewGenesisAccount(launchID, "", sdk.NewCoins())
	)

	coord.CoordinatorId = coordID
	err := tk.ProfileKeeper.Coordinator.Set(ctx, coord.CoordinatorId, coord)
	require.NoError(t, err)
	chain := sample.Chain(r, launchID, coordID)
	err = tk.LaunchKeeper.Chain.Set(ctx, chain.LaunchId, chain)
	require.NoError(t, err)

	tests := []struct {
		name    string
		request types.Request
		wantErr bool
	}{
		{
			name:    "should allow applying GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
		},
		{
			name:    "should prevent applying duplicated GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
			wantErr: true,
		},
		{
			name:    "should allow applying genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
		},
		{
			name:    "should prevent applying AccountRemoval when account not found",
			request: sample.RequestWithContent(r, launchID, contents[1]),
			wantErr: true,
		},
		{
			name:    "should allow applying VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
		},
		{
			name:    "should prevent applying duplicated VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
			wantErr: true,
		},
		{
			name:    "should allow applying vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
		},
		{
			name:    "should prevent applying vesting AccountRemoval content when account not found",
			request: sample.RequestWithContent(r, launchID, contents[3]),
			wantErr: true,
		},
		{
			name:    "should allow applying GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
		},
		{
			name:    "should prevent applying duplicated GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
			wantErr: true,
		},
		{
			name:    "should allow applying ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
		},
		{
			name:    "should prevent applying ValidatorRemoval when validator not found",
			request: sample.RequestWithContent(r, launchID, contents[5]),
			wantErr: true,
		},
		{
			name:    "should prevent applying invalid request content",
			request: sample.RequestWithContent(r, launchID, invalidContent),
			wantErr: true,
		},
		{
			name: "should prevent applying empty request content",
			request: types.Request{
				Content: types.RequestContent{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.ApplyRequest(ctx, tk.LaunchKeeper, chain, tt.request, coord)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			switch requestContent := tt.request.Content.Content.(type) {
			case *types.RequestContent_GenesisAccount:
				ga := requestContent.GenesisAccount
				address, err := tk.LaunchKeeper.AddressCodec().StringToBytes(ga.Address)
				require.NoError(t, err)
				_, err = tk.LaunchKeeper.GenesisAccount.Get(ctx, collections.Join(launchID, sdk.AccAddress(address)))
				require.NoError(t, err, "genesis account not found")
			case *types.RequestContent_VestingAccount:
				va := requestContent.VestingAccount
				address, err := tk.LaunchKeeper.AddressCodec().StringToBytes(va.Address)
				require.NoError(t, err)
				_, err = tk.LaunchKeeper.VestingAccount.Get(ctx, collections.Join(launchID, sdk.AccAddress(address)))
				require.NoError(t, err, "vesting account not found")
			case *types.RequestContent_AccountRemoval:
				ar := requestContent.AccountRemoval
				address, err := tk.LaunchKeeper.AddressCodec().StringToBytes(ar.Address)
				require.NoError(t, err)
				_, err = tk.LaunchKeeper.GenesisAccount.Get(ctx, collections.Join(launchID, sdk.AccAddress(address)))
				require.Error(t, err, "genesis account not removed")
				_, err = tk.LaunchKeeper.VestingAccount.Get(ctx, collections.Join(launchID, sdk.AccAddress(address)))
				require.Error(t, err, "vesting account not removed")
			case *types.RequestContent_GenesisValidator:
				ga := requestContent.GenesisValidator
				address, err := tk.LaunchKeeper.AddressCodec().StringToBytes(ga.Address)
				require.NoError(t, err)
				_, err = tk.LaunchKeeper.GenesisValidator.Get(ctx, collections.Join(launchID, sdk.AccAddress(address)))
				require.NoError(t, err, "genesis validator not found")
			case *types.RequestContent_ValidatorRemoval:
				vr := requestContent.ValidatorRemoval
				address, err := tk.LaunchKeeper.AddressCodec().StringToBytes(vr.ValAddress)
				require.NoError(t, err)
				_, err = tk.LaunchKeeper.GenesisValidator.Get(ctx, collections.Join(launchID, sdk.AccAddress(address)))
				require.Error(t, err, "genesis validator not removed")
			}
		})
	}
}

func TestCheckRequest(t *testing.T) {
	var (
		coord                           = sample.Coordinator(r, sample.Address(r))
		coordID                         = uint64(3)
		genesisAcc                      = sample.AccAddress(r)
		vestingAcc                      = sample.AccAddress(r)
		validatorAcc                    = sample.AccAddress(r)
		duplicatedAcc                   = sample.AccAddress(r)
		ctx, tk, _                      = testkeeper.NewTestSetup(t)
		launchID                        = uint64(10)
		contents                        = sample.AllRequestContents(r, launchID, genesisAcc.String(), vestingAcc.String(), validatorAcc.String())
		invalidContent                  = types.NewGenesisAccount(launchID, "", sdk.NewCoins())
		duplicatedRequestGenesisContent = types.NewGenesisAccount(launchID, duplicatedAcc.String(), sample.Coins(r))
		duplicatedRequestVestingContent = types.NewVestingAccount(launchID, duplicatedAcc.String(), sample.VestingOptions(r))
		duplicatedRequestRemovalContent = types.NewAccountRemoval(duplicatedAcc.String())
	)

	coord.CoordinatorId = coordID
	err := tk.ProfileKeeper.Coordinator.Set(ctx, coord.CoordinatorId, coord)
	require.NoError(t, err)
	chain := sample.Chain(r, launchID, coordID)
	err = tk.LaunchKeeper.Chain.Set(ctx, chain.LaunchId, chain)
	require.NoError(t, err)

	err = tk.LaunchKeeper.GenesisAccount.Set(ctx, collections.Join(launchID, duplicatedAcc), types.GenesisAccount{
		LaunchId: launchID,
		Address:  duplicatedAcc.String(),
		Coins:    nil,
	})
	require.NoError(t, err)

	err = tk.LaunchKeeper.VestingAccount.Set(ctx, collections.Join(launchID, duplicatedAcc), types.VestingAccount{
		LaunchId:       launchID,
		Address:        duplicatedAcc.String(),
		VestingOptions: types.VestingOptions{},
	})
	require.NoError(t, err)

	tests := []struct {
		name    string
		request types.Request
		err     error
	}{
		{
			name:    "should validate valid GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
		},
		{
			name:    "should prevent validate duplicated GenesisAccount content",
			request: sample.RequestWithContent(r, launchID, contents[0]),
			err:     types.ErrAccountAlreadyExist,
		},
		{
			name:    "should validate valid genesis AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[1]),
		},
		{
			name:    "should prevent validate AccountRemoval with no account",
			request: sample.RequestWithContent(r, launchID, contents[1]),
			err:     types.ErrAccountNotFound,
		},
		{
			name:    "should validate valid VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
		},
		{
			name:    "should prevent validate duplicated VestingAccount content",
			request: sample.RequestWithContent(r, launchID, contents[2]),
			err:     types.ErrAccountAlreadyExist,
		},
		{
			name:    "should validate valid vesting AccountRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[3]),
		},
		{
			name:    "should validate vesting AccountRemoval content with vesting account not found",
			request: sample.RequestWithContent(r, launchID, contents[3]),
			err:     types.ErrAccountNotFound,
		},
		{
			name:    "should validate valid GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
		},
		{
			name:    "should prevent validate duplicated GenesisValidator content",
			request: sample.RequestWithContent(r, launchID, contents[4]),
			err:     types.ErrValidatorAlreadyExist,
		},
		{
			name:    "should validate valid ValidatorRemoval content",
			request: sample.RequestWithContent(r, launchID, contents[5]),
		},
		{
			name:    "should prevent validate ValidatorRemoval content with no validator to remove",
			request: sample.RequestWithContent(r, launchID, contents[5]),
			err:     types.ErrValidatorNotFound,
		},
		{
			name:    "should prevent validate request content with invalid parameters",
			request: sample.RequestWithContent(r, launchID, invalidContent),
			err:     ignterrors.ErrCritical,
		},
		{
			name:    "should prevent validate with critical error genesis account request content with genesis and vesting account",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestGenesisContent),
			err:     ignterrors.ErrCritical,
		},
		{
			name:    "should prevent validate with critical error vesting account request content with genesis and vesting account",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestVestingContent),
			err:     ignterrors.ErrCritical,
		},
		{
			name:    "should prevent validate with critical error account removal request content with genesis and vesting account",
			request: sample.RequestWithContent(r, launchID, duplicatedRequestRemovalContent),
			err:     ignterrors.ErrCritical,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := keeper.CheckRequest(ctx, tk.LaunchKeeper, launchID, tt.request)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
			} else {
				err := keeper.ApplyRequest(ctx, tk.LaunchKeeper, chain, tt.request, coord)
				require.NoError(t, err)
			}
		})
	}
}
