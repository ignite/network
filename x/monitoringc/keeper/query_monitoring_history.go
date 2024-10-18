package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ignite/network/x/monitoringc/types"
)

func (q queryServer) GetMonitoringHistory(ctx context.Context, req *types.QueryGetMonitoringHistoryRequest) (*types.QueryGetMonitoringHistoryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.MonitoringHistory.Get(ctx, req.LaunchID)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetMonitoringHistoryResponse{MonitoringHistory: val}, nil
}
