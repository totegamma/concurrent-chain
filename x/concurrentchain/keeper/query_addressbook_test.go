package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/totegamma/concurrent-chain/testutil/keeper"
	"github.com/totegamma/concurrent-chain/testutil/nullify"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAddressbookQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ConcurrentchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAddressbook(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetAddressbookRequest
		response *types.QueryGetAddressbookResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAddressbookRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetAddressbookResponse{Addressbook: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAddressbookRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetAddressbookResponse{Addressbook: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAddressbookRequest{
				Index: strconv.Itoa(100000),
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
			response, err := keeper.Addressbook(wctx, tc.request)
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

func TestAddressbookQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ConcurrentchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAddressbook(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAddressbookRequest {
		return &types.QueryAllAddressbookRequest{
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
			resp, err := keeper.AddressbookAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Addressbook), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Addressbook),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AddressbookAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Addressbook), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Addressbook),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AddressbookAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Addressbook),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AddressbookAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
