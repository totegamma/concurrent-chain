package concurrentchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/keeper"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the addressbook
	for _, elem := range genState.AddressbookList {
		k.SetAddressbook(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AddressbookList = k.GetAllAddressbook(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
