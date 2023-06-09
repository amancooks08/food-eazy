package service

import (
	"auth-service/errors"
	"auth-service/models"
	"auth-service/utils"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func RegisterUser(name string, email string, password string, phoneNumber string, role string) (err error) {
	if len(name) == 0 || len(email) == 0 || len(password) == 0 || len(phoneNumber) == 0 || len(role) == 0 {
		return errors.ErrEmptyField
	}
	err = utils.ValidateUserDetails(email, password, phoneNumber)
	if err != nil {
		return err
	}
	newUser := models.User{
		Name:        name,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Role:        role,
	}

	newUser.Password, _ = utils.HashPassword(newUser.Password)

	err = models.RegisterUser(&newUser)
	return err
}

func LoginUser(email string, password string) (token string, err error) {
	if len(email) == 0 || len(password) == 0 {
		return "", errors.ErrEmptyField
	}
	if !utils.ValidateEmail(email) {
		return "", errors.ErrInvalidEmail
	}
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.ErrInvalidPassword
	}
	token, err = utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		logger.WithField("error", err).Error(errors.ErrTokenGeneration.Error())
		return "", errors.ErrTokenGeneration
	}
	return
}

func ValidateUser(token string, role string) (int, error) {
	if len(token) == 0 || len(role) == 0 {
		return http.StatusBadRequest, errors.ErrEmptyField
	}
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return http.StatusForbidden, errors.ErrInvalidToken
	}
	if claims["role"] != role {
		return http.StatusForbidden, errors.ErrInvalidRole
	}
	return http.StatusOK, nil
}