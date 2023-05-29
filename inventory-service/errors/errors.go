package errors

import(
	"errors"
)

var (
	ErrFailedToConnectDB = errors.New("failed to connect to database")
	ErrFailedToCloseDB = errors.New("failed to close database")
	
	ErrItemNotFound = errors.New("item not found")
	ErrInvalidItem = errors.New("invalid item")
	ErrAddItem = errors.New("failed to add item")
	
	ErrEmptyField = errors.New("empty field")
)
