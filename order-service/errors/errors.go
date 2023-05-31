package errors

import "errors"

var (
	ErrFailedToConnectDB = errors.New("failed to connect to database")
	ErrFailedToCloseDB   = errors.New("failed to close database")
	ErrNoOrdersAvailable = errors.New("no orders available")
	ErrInvalidOrder	= errors.New("invalid order")
	ErrEmptyField = errors.New("empty field(s)")
	ErrLimitedSupplies = errors.New("limited supplies")
	ErrBadGateway = errors.New("bad gateway while connecting to inventory service")
)
