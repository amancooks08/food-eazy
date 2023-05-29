package errors

import "errors"

var (
	ErrFailedToConnectDB = errors.New("failed to connect to database")
	ErrFailedToCloseDB   = errors.New("failed to close database")

	ErrEmptyField = errors.New("empty field(s)")
	ErrLimitedSupplies = errors.New("limited supplies")
)
