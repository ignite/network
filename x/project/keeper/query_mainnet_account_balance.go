package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) MainnetAccountBalance(ctx context.Context, req *types.QueryMainnetAccountBalanceRequest) (*types.QueryMainnetAccountBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	address, err := q.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	// get project and share information
	totalShareNumber, err := q.k.TotalShares.Get(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	project, err := q.k.GetProject(ctx, req.ProjectID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "project not found")
	}

	// get account balance
	acc, err := q.k.GetMainnetAccount(ctx, req.ProjectID, address)
	if err != nil {
		return nil, status.Error(codes.NotFound, "account not found")
	}

	balance, err := acc.Shares.CoinsFromTotalSupply(project.TotalSupply, totalShareNumber)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "balance can't be calculated: %s", err.Error())
	}

	mainnetAccountBalance := types.MainnetAccountBalance{
		ProjectID: acc.ProjectID,
		Address:   acc.Address,
		Coins:     balance,
	}

	return &types.QueryMainnetAccountBalanceResponse{MainnetAccountBalance: mainnetAccountBalance}, nil
}

func (q queryServer) ListMainnetAccountBalance(ctx context.Context, req *types.QueryListMainnetAccountBalanceRequest) (*types.QueryListMainnetAccountBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// get project and share information
	totalShareNumber, err := q.k.TotalShares.Get(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	project, err := q.k.GetProject(ctx, req.ProjectID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "project not found")
	}

	// iterate accounts
	mainnetAccountBalances := make([]types.MainnetAccountBalance, 0)
	_, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.MainnetAccount,
		req.Pagination,
		func(_ collections.Pair[uint64, sdk.AccAddress], acc types.MainnetAccount) (types.MainnetAccount, error) {
			balance, err := acc.Shares.CoinsFromTotalSupply(project.TotalSupply, totalShareNumber)
			if err != nil {
				return acc, fmt.Errorf("balance can't be calculated for account %s: %s", acc.Address, err.Error())
			}
			// add the balance if not zero
			if !balance.IsZero() {
				mainnetAccountBalance := types.MainnetAccountBalance{
					ProjectID: acc.ProjectID,
					Address:   acc.Address,
					Coins:     balance,
				}
				mainnetAccountBalances = append(mainnetAccountBalances, mainnetAccountBalance)
			}
			return acc, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListMainnetAccountBalanceResponse{MainnetAccountBalance: mainnetAccountBalances, Pagination: pageRes}, nil
}
