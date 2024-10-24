package keeper_test

import (
	"context"
	"fmt"
	"testing"

	"cosmossdk.io/collections"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tc "github.com/ignite/network/testutil/constructor"
	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/project/keeper"
	"github.com/ignite/network/x/project/types"
)

func createNMainnetAccountForProjectIDWithTotalSupply(
	t *testing.T,
	keeper keeper.Keeper,
	ctx context.Context,
	n int,
	projectID uint64,
) []types.MainnetAccountBalance {
	totalSupply := tc.Coins(t, "100000foo,200000bar")
	totalShares := uint64(100000)

	// create and set project
	project := sample.Project(r, projectID)
	project.TotalSupply = totalSupply
	err := keeper.Project.Set(ctx, projectID, project)
	require.NoError(t, err)
	err = keeper.TotalShares.Set(ctx, totalShares)
	require.NoError(t, err)

	// set account and create n account balance
	// shares of accounts are foo and bar shares with random share number
	items := make([]types.MainnetAccountBalance, n)
	for i := range items {
		addr := sample.AccAddress(r)
		acc := sample.MainnetAccount(r, projectID, addr.String())
		fooShares := r.Intn(int(totalShares))
		barShares := r.Intn(int(totalShares))
		acc.Shares = tc.Shares(t, fmt.Sprintf("%dfoo,%dbar", fooShares, barShares))
		err = keeper.MainnetAccount.Set(ctx, collections.Join(projectID, addr), acc)
		require.NoError(t, err)

		balance, err := acc.Shares.CoinsFromTotalSupply(totalSupply, totalShares)
		require.NoError(t, err)
		items[i] = types.MainnetAccountBalance{
			ProjectId: projectID,
			Address:   acc.Address,
			Coins:     balance,
		}
	}
	return items
}

func TestMainnetAccountBalanceQuerySingle(t *testing.T) {
	var (
		projectID = uint64(5)
		k, ctx, _ = keepertest.ProjectKeeper(t)
		qs        = keeper.NewQueryServerImpl(k)
		msgs      = createNMainnetAccountForProjectIDWithTotalSupply(t, k, ctx, 5, projectID)
	)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryMainnetAccountBalanceRequest
		response *types.QueryMainnetAccountBalanceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryMainnetAccountBalanceRequest{
				ProjectId: msgs[0].ProjectId,
				Address:   msgs[0].Address,
			},
			response: &types.QueryMainnetAccountBalanceResponse{MainnetAccountBalance: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryMainnetAccountBalanceRequest{
				ProjectId: msgs[1].ProjectId,
				Address:   msgs[1].Address,
			},
			response: &types.QueryMainnetAccountBalanceResponse{MainnetAccountBalance: msgs[1]},
		},
		{
			desc: "project not found",
			request: &types.QueryMainnetAccountBalanceRequest{
				ProjectId: 10000,
				Address:   sample.Address(r),
			},
			err: status.Error(codes.NotFound, "project not found"),
		},
		{
			desc: "account not found",
			request: &types.QueryMainnetAccountBalanceRequest{
				ProjectId: projectID,
				Address:   sample.Address(r),
			},
			err: status.Error(codes.NotFound, "account not found"),
		},
		{
			desc: "invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.MainnetAccountBalance(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestMainnetAccountBalanceQueryPaginated(t *testing.T) {
	var (
		projectID = uint64(5)
		k, ctx, _ = keepertest.ProjectKeeper(t)
		qs        = keeper.NewQueryServerImpl(k)
		msgs      = createNMainnetAccountForProjectIDWithTotalSupply(t, k, ctx, 5, projectID)
	)
	request := func(projectID uint64, next []byte, offset, limit uint64, total bool) *types.QueryListMainnetAccountBalanceRequest {
		return &types.QueryListMainnetAccountBalanceRequest{
			ProjectId: projectID,
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListMainnetAccountBalance(ctx, request(projectID, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccountBalance), step)
			require.Subset(t, msgs, resp.MainnetAccountBalance)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListMainnetAccountBalance(ctx, request(projectID, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MainnetAccountBalance), step)
			require.Subset(t, msgs, resp.MainnetAccountBalance)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListMainnetAccountBalance(ctx, request(projectID, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t, msgs, resp.MainnetAccountBalance)
	})
	t.Run("invalid request", func(t *testing.T) {
		_, err := qs.ListMainnetAccountBalance(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
	t.Run("project not found", func(t *testing.T) {
		_, err := qs.ListMainnetAccountBalance(ctx, request(10000, nil, 0, 0, true))
		require.ErrorIs(t, err, status.Error(codes.NotFound, "project not found"))
	})
}

func TestMainnetAccountBalanceAll(t *testing.T) {
	var (
		k, ctx, _ = keepertest.ProjectKeeper(t)
		qs        = keeper.NewQueryServerImpl(k)

		projectID   = uint64(5)
		totalSupply = tc.Coins(t, "1000foo,1000bar")
		totalShares = uint64(100)
		addr1       = sample.AccAddress(r)
		addr2       = sample.AccAddress(r)
		addr3       = sample.AccAddress(r)
		project     = sample.Project(r, projectID)
	)

	// set project and sample accounts
	project.TotalSupply = totalSupply
	err := k.Project.Set(ctx, projectID, project)
	require.NoError(t, err)
	err = k.TotalShares.Set(ctx, totalShares)
	require.NoError(t, err)
	err = k.MainnetAccount.Set(ctx, collections.Join(projectID, addr1), types.MainnetAccount{
		ProjectId: projectID,
		Address:   addr1.String(),
		Shares:    tc.Shares(t, "100foo"),
	})
	require.NoError(t, err)
	err = k.MainnetAccount.Set(ctx, collections.Join(projectID, addr2), types.MainnetAccount{
		ProjectId: projectID,
		Address:   addr2.String(),
		Shares:    tc.Shares(t, "100bar"),
	})
	require.NoError(t, err)
	err = k.MainnetAccount.Set(ctx, collections.Join(projectID, addr3), types.MainnetAccount{
		ProjectId: projectID,
		Address:   addr3.String(),
		Shares:    tc.Shares(t, "100baz"),
	})
	require.NoError(t, err)

	t.Run("accounts with empty balance are skipped", func(t *testing.T) {
		accountBalances, err := qs.ListMainnetAccountBalance(ctx, &types.QueryListMainnetAccountBalanceRequest{
			ProjectId: projectID,
			Pagination: &query.PageRequest{
				CountTotal: true,
			},
		})
		require.NoError(t, err)

		// Account 3 must not be included in balances since the total supply doesn't contains baz tokens
		balances := accountBalances.MainnetAccountBalance
		require.Len(t, balances, 2)
		require.Contains(t, balances, types.MainnetAccountBalance{
			ProjectId: projectID,
			Address:   addr1.String(),
			Coins:     tc.Coins(t, "1000foo"),
		})
		require.Contains(t, balances, types.MainnetAccountBalance{
			ProjectId: projectID,
			Address:   addr2.String(),
			Coins:     tc.Coins(t, "1000bar"),
		})
	})
}
