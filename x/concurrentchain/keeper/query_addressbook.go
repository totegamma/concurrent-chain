package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressbookAll(goCtx context.Context, req *types.QueryAllAddressbookRequest) (*types.QueryAllAddressbookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addressbooks []types.Addressbook
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	addressbookStore := prefix.NewStore(store, types.KeyPrefix(types.AddressbookKeyPrefix))

	pageRes, err := query.Paginate(addressbookStore, req.Pagination, func(key []byte, value []byte) error {
		var addressbook types.Addressbook
		if err := k.cdc.Unmarshal(value, &addressbook); err != nil {
			return err
		}

		addressbooks = append(addressbooks, addressbook)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressbookResponse{Addressbook: addressbooks, Pagination: pageRes}, nil
}

func (k Keeper) Addressbook(goCtx context.Context, req *types.QueryGetAddressbookRequest) (*types.QueryGetAddressbookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetAddressbook(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAddressbookResponse{Addressbook: val}, nil
}
