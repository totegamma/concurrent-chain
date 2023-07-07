package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

// SetAddressbook set a specific addressbook in the store from its index
func (k Keeper) SetAddressbook(ctx sdk.Context, addressbook types.Addressbook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressbookKeyPrefix))
	b := k.cdc.MustMarshal(&addressbook)
	store.Set(types.AddressbookKey(
		addressbook.Index,
	), b)
}

// GetAddressbook returns a addressbook from its index
func (k Keeper) GetAddressbook(
	ctx sdk.Context,
	index string,

) (val types.Addressbook, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressbookKeyPrefix))

	b := store.Get(types.AddressbookKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAddressbook removes a addressbook from the store
func (k Keeper) RemoveAddressbook(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressbookKeyPrefix))
	store.Delete(types.AddressbookKey(
		index,
	))
}

// GetAllAddressbook returns all addressbook
func (k Keeper) GetAllAddressbook(ctx sdk.Context) (list []types.Addressbook) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddressbookKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Addressbook
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
