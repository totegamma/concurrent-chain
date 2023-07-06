package concurrentchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/totegamma/concurrent-chain/testutil/keeper"
	"github.com/totegamma/concurrent-chain/testutil/nullify"
	"github.com/totegamma/concurrent-chain/x/concurrentchain"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ConcurrentchainKeeper(t)
	concurrentchain.InitGenesis(ctx, *k, genesisState)
	got := concurrentchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
