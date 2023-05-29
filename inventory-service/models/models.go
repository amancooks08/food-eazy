package models

import (
	"inventory-service/errors"

	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

type Item struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey; column:id; autoIncrement; not null"`
	Name        string  `gorm:"column:name; unique; not null"`
	Description string  `gorm:"column:description; not null"`
	Price       float64 `gorm:"column:price; not null"`
	Quantity    uint    `gorm:"column:quantity; not null"`
}

func InitInventoryModels(database *gorm.DB) {
	db = database
	db.AutoMigrate(&Item{})
}

func CreateItem(item *Item) (*Item, error) {
	if item == nil {
		logger.WithField("error", errors.ErrInvalidItem.Error()).Error(errors.ErrInvalidItem.Error())
		return nil, errors.ErrInvalidItem
	}
	err := db.Create(item).Error
	if err != nil && err.Error() == "UNIQUE constraint failed: items.name" {
		return nil, errors.ErrItemExists
	} else if err != nil {
		logger.WithField("error", err.Error()).Error(err.Error())
		return nil, err
	}
	return item, nil
}

func GetItem(id uint) (*Item, error) {
	item := &Item{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil && err.Error() == "record not found" {
		return nil, errors.ErrItemNotFound
	} else if err != nil {
		logger.WithField("error", err.Error()).Error(err.Error())
		return nil, err
	}
	return item, nil
}


func GetAllItems() ([]*Item, error) {
	items := []*Item{}
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}