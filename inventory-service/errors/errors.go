package errors

import(
	"errors"
)

var (
	ErrFailedToConnectDB = errors.New("failed to connect to database")
	ErrFailedToCloseDB = errors.New("failed to close database")
	ErrInsufficientQuantity = errors.New("insufficient quantity")
	ErrItemNotFound = errors.New("item not found")
	ErrInvalidItem = errors.New("invalid item")
	ErrAddItem = errors.New("failed to add item")
	ErrItemExists = errors.New("item already exists")
	ErrEmptyField = errors.New("empty field")
)
