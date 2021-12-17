package nalo

import (
	"errors"
)

var (
	ErrInvalidURL                 = errors.New("invalid url")
	ErrInvalidCredentials         = errors.New("invalid value in 'username' or 'password' field")
	ErrInvalidType                = errors.New("invalid value in 'type' field")
	ErrInvalidMessage             = errors.New("invalid message")
	ErrInvalidDestination         = errors.New("invalid destination")
	ErrInvalidSource              = errors.New("invalid source")
	ErrInvalidDLR                 = errors.New("invalid value for 'dlr' field")
	ErrInvalidUserValidation      = errors.New("user validation failed")
	ErrInternal                   = errors.New("internal error")
	ErrInsufficientCreditUser     = errors.New("insufficient credit user")
	ErrInsufficientCreditReseller = errors.New("insufficient credit reseller")
	ErrUnknown                    = errors.New("unknown error")
)
