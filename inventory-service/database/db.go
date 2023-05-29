package database

import (
	"inventory-service/errors"
	"inventory-service/models"

	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() error {
	var err error
	dbURI := "host=localhost user=amandeep password=Aman@123 dbname=testdb port=5432 sslmode=disable"
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
	models.InitInventoryModels(db)
}
