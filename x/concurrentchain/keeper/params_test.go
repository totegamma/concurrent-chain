package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/totegamma/concurrent-chain/testutil/keeper"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ConcurrentchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
