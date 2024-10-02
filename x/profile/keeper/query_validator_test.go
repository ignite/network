package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/ignite/network/testutil/keeper"
	"github.com/ignite/network/testutil/nullify"
	"github.com/ignite/network/testutil/sample"
	"github.com/ignite/network/x/profile/keeper"
	"github.com/ignite/network/x/profile/types"
)

func createNValidator(keeper keeper.Keeper, ctx context.Context, n int) []types.Validator {
	items := make([]types.Validator, n)
	for i := range items {
		items[i].Address = sample.Address(r)
		_ = keeper.Validator.Set(ctx, items[i].Address, items[i])
	}
	return items
}

func TestValidatorQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ProfileKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNValidator(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetValidatorRequest
		response *types.QueryGetValidatorResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetValidatorRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetValidatorResponse{Validator: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetValidatorRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetValidatorResponse{Validator: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetValidatorRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetValidator(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestValidatorQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.ProfileKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNValidator(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllValidatorRequest {
		return &types.QueryAllValidatorRequest{
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
			resp, err := qs.ListValidator(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Validator),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListValidator(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Validator), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Validator),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListValidator(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Validator),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListValidator(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

func createNValidatorByOperatorAddress(keeper keeper.Keeper, ctx context.Context, n int) []types.Validator {
	valByOp := make([]types.ValidatorByOperatorAddress, n)
	validators := make([]types.Validator, n)
	for i := range valByOp {
		valAddress := sample.Address(r)
		opAddress := sample.Address(r)
		valByOp[i].ValidatorAddress = valAddress
		valByOp[i].OperatorAddress = opAddress
		_ = keeper.ValidatorByOperatorAddress.Set(ctx, opAddress, valByOp[i])

		validators[i].Address = valAddress
		_ = keeper.Validator.Set(ctx, valAddress, validators[i])
	}
	return validators
}

func TestValidatorByOperatorAddressQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.ProfileKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNValidatorByOperatorAddress(k, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorByOperatorAddressRequest
		response *types.QueryGetValidatorByOperatorAddressResponse
		err      error
	}{
		{
			desc: "should allow querying first validator by operator address",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetValidatorByOperatorAddressResponse{Validator: msgs[0]},
		},
		{
			desc: "should allow querying second validator by operator address",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetValidatorByOperatorAddressResponse{Validator: msgs[1]},
		},
		{
			desc: "should prevent querying non existing validator by operator address",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				Address: sample.Address(r),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "should prevent querying with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetValidatorByOperatorAddress(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
