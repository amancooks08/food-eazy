package errors


import(
	"errors"
)

var (
	ErrFailedToConnectDB = errors.New("failed to connect to database")
	ErrFailedToCloseDB = errors.New("failed to close database")
	
	ErrInvalidUser = errors.New("invalid user")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrInvalidEmail = errors.New("invalid email")
	ErrDuplicateEmail = errors.New("email already exists")
	ErrShortPassword = errors.New("password must be atleast 8 characters long")
	ErrEmptyField = errors.New("empty field")

	ErrCreateUser = errors.New("failed to create user")
)