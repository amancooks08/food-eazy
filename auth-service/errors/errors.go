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
	ErrTokenGeneration = errors.New("failed to generate token")
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrInvalidEmail = errors.New("invalid email")
	ErrDuplicateEmail = errors.New("email already exists")
	ErrEmptyField = errors.New("empty field")
	ErrInvalidRole = errors.New("invalid role")
	ErrInvalidPassword = errors.New("invalid password")
	ErrHashPassword	= errors.New("failed to hash password")
	ErrShortPassword = errors.New("password must be atleast 8 characters long")

	ErrUnauthorized = errors.New("unauthorized")
)