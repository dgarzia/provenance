package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/msgfees module sentinel errors
var (
	ErrEmptyMsgType        = sdkerrors.Register(ModuleName, 2, "msg type is empty")
	ErrInvalidCoinAmount   = sdkerrors.Register(ModuleName, 3, "invalid coin amount")
	ErrInvalidFee          = sdkerrors.Register(ModuleName, 4, "invalid fee amount")
	ErrMsgFeeAlreadyExists = sdkerrors.Register(ModuleName, 4, "fee for type already exists.")
)
