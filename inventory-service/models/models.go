package models

import (
	"inventory-service/errors"
	"net/http"

	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var db *gorm.DB

type Item struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey; column:id; autoIncrement; not null"`
	Name        string  `gorm:"column:name; unique; not null"`
	Description string  `gorm:"column:description; not null"`
	Price       float32 `gorm:"column:price; not null"`
	Quantity    uint    `gorm:"column:quantity; not null"`
}

func InitInventoryModels(database *gorm.DB) {
	db = database
	db.AutoMigrate(&Item{})
}

func CreateItem(item *Item) (uint32, *Item, error) {
	if item == nil {
		logger.WithField("error", errors.ErrInvalidItem.Error()).Error(errors.ErrInvalidItem.Error())
		return http.StatusBadRequest, nil, errors.ErrInvalidItem
	}
	err := db.Create(item).Error
	if err != nil && err.Error() == "UNIQUE constraint failed: items.name" {
		return http.StatusUnprocessableEntity, nil, errors.ErrItemExists
	} else if err != nil {
		logger.WithField("error", err.Error()).Error(err.Error())
		return http.StatusBadRequest, nil, err
	}
	return http.StatusCreated, item, nil
}

func GetItem(id uint) (uint32, *Item, error) {
	item := &Item{}
	err := db.Where("id = ?", id).First(item).Error
	if err != nil && err.Error() == "record not found" {
		return http.StatusNotFound, nil, errors.ErrItemNotFound
	} else if err != nil {
		logger.WithField("error", err.Error()).Error(err.Error())
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, item, nil
}

func GetAllItems() (uint32, []*Item, error) {
	items := []*Item{}
	err := db.Find(&items).Error
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, items, nil
}

func UpdateItemQuantity(id uint, quantity uint) (uint32, error) {
	status, item, err := GetItem(id)
	if err != nil {
		return status, err
	}
	item.Quantity = quantity
	err = db.Save(item).Error
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func DeleteItem(id uint) (uint32, error) {
	status, item, err := GetItem(id)
	if err != nil {
		return status, err
	}
	err = db.Delete(item).Error
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}