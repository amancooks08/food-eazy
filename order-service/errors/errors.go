package errors

import "errors"

var (
	ErrFailedToConnectDB = errors.New("failed to connect to database")
	ErrFailedToCloseDB   = errors.New("failed to close database")
)
