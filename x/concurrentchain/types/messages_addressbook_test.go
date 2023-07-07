package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/totegamma/concurrent-chain/testutil/sample"
)

func TestMsgCreateAddressbook_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateAddressbook
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateAddressbook{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateAddressbook{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateAddressbook_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAddressbook
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateAddressbook{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateAddressbook{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteAddressbook_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAddressbook
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteAddressbook{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteAddressbook{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
