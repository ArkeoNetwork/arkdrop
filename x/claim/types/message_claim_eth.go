package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimEth = "claim_eth"

var _ sdk.Msg = &MsgClaimEth{}

func NewMsgClaimEth(creator string, ethAddress string, signature string) *MsgClaimEth {
	return &MsgClaimEth{
		Creator:    creator,
		EthAddress: ethAddress,
		Signature:  signature,
	}
}

func (msg *MsgClaimEth) Route() string {
	return RouterKey
}

func (msg *MsgClaimEth) Type() string {
	return TypeMsgClaimEth
}

func (msg *MsgClaimEth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimEth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimEth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if !IsValidEthAddress(msg.EthAddress) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid eth address (%s)", err)
	}

	return nil
}
