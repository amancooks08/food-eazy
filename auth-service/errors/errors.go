package errors


import(
	"errors"
)

var (
	ErrFailedToConnectDB = errors.New("failed to connect to database")
	ErrFailedToCloseDB = errors.New("failed to close database")
	
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidUser = errors.New("invalid user")
	ErrCreateUser = errors.New("failed to create user")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrInvalidEmail = errors.New("invalid email")
	ErrDuplicateEmail = errors.New("email already exists")
	ErrEmptyField = errors.New("empty field")

	ErrInvalidPassword = errors.New("invalid password")
	ErrHashPassword	= errors.New("failed to hash password")
	ErrShortPassword = errors.New("password must be atleast 8 characters long")

)