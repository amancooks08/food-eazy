package service

import (
	"auth-service/errors"
	"auth-service/models"
	"auth-service/utils"
	"regexp"

	logger "github.com/sirupsen/logrus"
)

func RegisterUser(name string, email string, password string, phoneNumber string, role string) (err error) {
	err = validateUserDetails(email, password, phoneNumber)
	if err != nil {
		return err
	}
	newUser := models.User{
		Name:        name,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Role:  		 role,
	}

	newUser.Password, _ = utils.HashPassword(newUser.Password)

	err = models.RegisterUser(&newUser)
	return err
}

func LoginUser(email string, password string) (token string, err error) {
	if len(email) == 0 || len(password) == 0 {
		return "", errors.ErrEmptyField
	}
	if !ValidateEmail(email) {
		return "", errors.ErrInvalidEmail
	}
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.ErrInvalidPassword
	}
	token, err = utils.GenerateToken(user.Email, user.Role)
	if err != nil {
		logger.WithField("error", err).Error(errors.ErrTokenGeneration.Error())
		return "", errors.ErrTokenGeneration
	}
	return
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
