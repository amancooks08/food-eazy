package service

import (
	"auth-service/errors"
	"auth-service/models"
	"regexp"
)

func RegisterUser(name string, email string, password string, phoneNumber string) (err error) {
	err = validateUserDetails(email, password, phoneNumber)
	if err != nil {
		return err
	}
	newUser := models.User{
		Name:        name,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}

	err = models.RegisterUser(&newUser)
	return err
}

func validateUserDetails(email string, password string, phoneNumber string) (err error) {
	if len(email) == 0 || len(password) == 0 || len(phoneNumber) == 0 {
		return errors.ErrEmptyField
	}
	if !ValidateEmail(email) {
		return errors.ErrInvalidEmail
	}
	if !ValidatePhoneNumber(phoneNumber) {
		return errors.ErrInvalidPhoneNumber
	}
	if len(password) < 8 {
		return errors.ErrShortPassword
	}
	return nil
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidatePhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(`^[0-9]{10}$`)
	return re.MatchString(phoneNumber)
}
