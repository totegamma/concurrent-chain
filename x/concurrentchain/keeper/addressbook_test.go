package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/totegamma/concurrent-chain/testutil/keeper"
	"github.com/totegamma/concurrent-chain/testutil/nullify"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/keeper"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAddressbook(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Addressbook {
	items := make([]types.Addressbook, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAddressbook(ctx, items[i])
	}
	return items
}

func TestAddressbookGet(t *testing.T) {
	keeper, ctx := keepertest.ConcurrentchainKeeper(t)
	items := createNAddressbook(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAddressbook(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAddressbookRemove(t *testing.T) {
	keeper, ctx := keepertest.ConcurrentchainKeeper(t)
	items := createNAddressbook(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAddressbook(ctx,
			item.Index,
		)
		_, found := keeper.GetAddressbook(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAddressbookGetAll(t *testing.T) {
	keeper, ctx := keepertest.ConcurrentchainKeeper(t)
	items := createNAddressbook(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAddressbook(ctx)),
	)
}
