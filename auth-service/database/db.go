package database

import (
	"auth-service/errors"
	"fmt"
	"os"
	"auth-service/models"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() error {
	var err error
	dbURI := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		logger.WithField("error", err).Fatal("failed to connect database")
		return errors.ErrFailedToConnectDB
	}

	// models to be intialized here
	initModels()

	logger.Info("Database connection established")

	return nil
}

func Close() error {
	sqlDB, err := db.DB()
	if err != nil {
		logger.WithField("error", err).Error("failed to close database")
		return errors.ErrFailedToCloseDB
	}

	sqlDB.Close()
	logger.Info("Database connection closed")
	return nil
}

func initModels() {
	logger.Info("Initializing auth-service models")
	models.InitAuthModels(db)
}
