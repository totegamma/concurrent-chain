package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

func (k msgServer) CreateAddressbook(goCtx context.Context, msg *types.MsgCreateAddressbook) (*types.MsgCreateAddressbookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetAddressbook(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var addressbook = types.Addressbook{
		Creator: msg.Creator,
		Index:   msg.Index,
		Fqdn:    msg.Fqdn,
	}

	k.SetAddressbook(
		ctx,
		addressbook,
	)
	return &types.MsgCreateAddressbookResponse{}, nil
}

func (k msgServer) UpdateAddressbook(goCtx context.Context, msg *types.MsgUpdateAddressbook) (*types.MsgUpdateAddressbookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAddressbook(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var addressbook = types.Addressbook{
		Creator: msg.Creator,
		Index:   msg.Index,
		Fqdn:    msg.Fqdn,
	}

	k.SetAddressbook(ctx, addressbook)

	return &types.MsgUpdateAddressbookResponse{}, nil
}

func (k msgServer) DeleteAddressbook(goCtx context.Context, msg *types.MsgDeleteAddressbook) (*types.MsgDeleteAddressbookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetAddressbook(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAddressbook(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteAddressbookResponse{}, nil
}
