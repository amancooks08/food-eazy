package service

import (
	"auth-service/errors"
	"auth-service/models"
	"auth-service/utils"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func RegisterUser(name string, email string, password string, phoneNumber string, role string) (err error) {
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
	token, err = utils.GenerateToken(user.Email, user.Role)
	if err != nil {
		logger.WithField("error", err).Error(errors.ErrTokenGeneration.Error())
		return "", errors.ErrTokenGeneration
	}
	return
}

func ValidateUser(token string) (int, string) {
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return http.StatusUnauthorized, errors.ErrInvalidToken.Error()
	}
	_, err = models.GetUserByEmail(claims["email"].(string))
	if err == nil {
		return http.StatusOK, ""
	}
	return http.StatusUnauthorized, errors.ErrInvalidToken.Error()
}