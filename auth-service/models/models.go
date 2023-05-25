package models

import (
	"auth-service/errors"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	ID 		    uint   `gorm:"primaryKey; autoIncrement; not null"`
	Name        string `gorm:"column:name;"`
	Email       string `gorm:"column:email; unique; not null"`
	Password    string `gorm:"column:password; not null"`
	PhoneNumber string `gorm:"column:phoneNumber; unique; not null"`
}

func InitAuthModels(database *gorm.DB) {
	db = database
	db.AutoMigrate(&User{})
}

func RegisterUser(user *User) error {
	if user == nil {
		return errors.ErrInvalidUser
	}
	if err := db.Create(user).Error; err != nil {
		logger.WithField("error", err).Error(errors.ErrCreateUser.Error())
		return errors.ErrCreateUser
	}
	return nil
}
