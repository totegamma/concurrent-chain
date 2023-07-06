package keeper

import (
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

var _ types.QueryServer = Keeper{}
