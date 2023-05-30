package models

import (
	"net/http"
	"order-service/errors"
	"gorm.io/gorm"
)

var db *gorm.DB

type Order struct {
	gorm.Model
	ID        uint32  `gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserID    uint32  `gorm:"column:user_id; not null"`
	ItemID    uint32  `gorm:"column:item_id; not null"`
	Quantity  uint32  `gorm:"column:quantity; not null"`
	Amount    float32 `gorm:"column:amount; not null"`
	OrderTime string  `gorm:"column:order_time; not null"`
}

func InitOrderModels(database *gorm.DB) {
	db = database
	db.AutoMigrate(&Order{})
}

func CreateOrder(order *Order) (uint32, error) {
	if order == nil {
		return http.StatusBadRequest, errors.ErrInvalidOrder
	}
	if err := db.Create(order).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func GetOrder(userID uint32) (uint32, *Order, error){
	if userID == 0 {
		return http.StatusBadRequest, nil, errors.ErrEmptyField
	}

	order := &Order{}
	if err := db.Where("user_id = ?", userID).First(order).Error; err != nil {
		return http.StatusInternalServerError, nil, err
	}
	
	return http.StatusOK, order, nil
}
