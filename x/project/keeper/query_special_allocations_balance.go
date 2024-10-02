package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/project/types"
)

func (q queryServer) SpecialAllocationsBalance(ctx context.Context, req *types.QuerySpecialAllocationsBalanceRequest) (*types.QuerySpecialAllocationsBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// get the project
	totalShareNumber, err := q.k.TotalShares.Get(ctx)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can't get total shares %s", err.Error())
	}

	project, err := q.k.GetProject(ctx, req.ProjectID)
	if err != nil {
		return nil, status.Error(codes.NotFound, "project not found")
	}

	// calculate special allocations balance from total supply
	genesisDistribution, err := project.SpecialAllocations.GenesisDistribution.CoinsFromTotalSupply(
		project.TotalSupply,
		totalShareNumber,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "genesis distribution can't be calculated: %s", err.Error())
	}
	claimableAirdrop, err := project.SpecialAllocations.ClaimableAirdrop.CoinsFromTotalSupply(
		project.TotalSupply,
		totalShareNumber,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "claimable airdrop can't be calculated: %s", err.Error())
	}

	return &types.QuerySpecialAllocationsBalanceResponse{
		GenesisDistribution: genesisDistribution,
		ClaimableAirdrop:    claimableAirdrop,
	}, nil
}
