package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAddressbook = "create_addressbook"
	TypeMsgUpdateAddressbook = "update_addressbook"
	TypeMsgDeleteAddressbook = "delete_addressbook"
)

var _ sdk.Msg = &MsgCreateAddressbook{}

func NewMsgCreateAddressbook(
	creator string,
	index string,
	fqdn string,

) *MsgCreateAddressbook {
	return &MsgCreateAddressbook{
		Creator: creator,
		Index:   index,
		Fqdn:    fqdn,
	}
}

func (msg *MsgCreateAddressbook) Route() string {
	return RouterKey
}

func (msg *MsgCreateAddressbook) Type() string {
	return TypeMsgCreateAddressbook
}

func (msg *MsgCreateAddressbook) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAddressbook) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAddressbook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAddressbook{}

func NewMsgUpdateAddressbook(
	creator string,
	index string,
	fqdn string,

) *MsgUpdateAddressbook {
	return &MsgUpdateAddressbook{
		Creator: creator,
		Index:   index,
		Fqdn:    fqdn,
	}
}

func (msg *MsgUpdateAddressbook) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAddressbook) Type() string {
	return TypeMsgUpdateAddressbook
}

func (msg *MsgUpdateAddressbook) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAddressbook) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAddressbook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAddressbook{}

func NewMsgDeleteAddressbook(
	creator string,
	index string,

) *MsgDeleteAddressbook {
	return &MsgDeleteAddressbook{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteAddressbook) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAddressbook) Type() string {
	return TypeMsgDeleteAddressbook
}

func (msg *MsgDeleteAddressbook) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAddressbook) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAddressbook) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
