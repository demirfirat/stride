package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/stakeibc module sentinel errors
var (
	ErrInvalidVersion              = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrInvalidToken                = sdkerrors.Register(ModuleName, 1502, "invalid token denom")
	ErrInvalidHostZone             = sdkerrors.Register(ModuleName, 1503, "host zone not registered")
	ErrICAStake                    = sdkerrors.Register(ModuleName, 1504, "ICA stake failed")
	ErrEpochNotFound               = sdkerrors.Register(ModuleName, 1505, "epoch not found")
	ErrRecordNotFound              = sdkerrors.Register(ModuleName, 1506, "record not found")
	ErrInvalidAmount               = sdkerrors.Register(ModuleName, 1507, "invalid amount")
	ErrValidatorAlreadyExists      = sdkerrors.Register(ModuleName, 1508, "validator already exists")
	ErrNoValidatorWeights          = sdkerrors.Register(ModuleName, 1509, "no non-zero validator weights")
	ErrValidatorNotFound           = sdkerrors.Register(ModuleName, 1510, "validator not found")
	ErrWeightsNotDifferent         = sdkerrors.Register(ModuleName, 1511, "validator weights haven't changed")
	ErrValidatorDelegationChg      = sdkerrors.Register(ModuleName, 1512, "can't change delegation on validator")
	ErrAcctNotScopedForFunc        = sdkerrors.Register(ModuleName, 1513, "this account can't call this function")
	ErrInsufficientFunds           = sdkerrors.Register(ModuleName, 1514, "balance is insufficient")
	ErrInvalidUserRedemptionRecord = sdkerrors.Register(ModuleName, 1515, "user redemption record error")
	ErrRequiredFieldEmpty          = sdkerrors.Register(ModuleName, 1516, "required field is missing")
	ErrInvalidNumValidator         = sdkerrors.Register(ModuleName, 1517, "invalid number of validators")
	ErrValidatorNotRemoved         = sdkerrors.Register(ModuleName, 1518, "validator not removed")
	ErrHostZoneNotFound            = sdkerrors.Register(ModuleName, 1519, "host zone not found")
	ErrOutsideIcqWindow            = sdkerrors.Register(ModuleName, 1520, "outside time window that accepts icqs")
	ErrParamNotFound               = sdkerrors.Register(ModuleName, 1521, "param not found")
	ErrUnmarshalFailure            = sdkerrors.Register(ModuleName, 1522, "unable to unmarshal data structure")
	ErrMarshalFailure              = sdkerrors.Register(ModuleName, 1523, "unable to marshal data structure")
	ErrInvalidPacketCompletionTime = sdkerrors.Register(ModuleName, 1524, "invalid packet completion time")
	ErrIntCast                     = sdkerrors.Register(ModuleName, 1525, "unable to cast to safe cast int")
	ErrFeeAccountNotRegistered     = sdkerrors.Register(ModuleName, 1526, "fee account is not registered")
)
