package utils

import (
	"auth-service/errors"
	logger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		logger.WithField("error", err).Error(errors.ErrHashPassword.Error())
		return "", errors.ErrHashPassword
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		logger.WithField("error", err).Error(errors.ErrInvalidPassword.Error())
		return false
	}
	return err == nil
}
