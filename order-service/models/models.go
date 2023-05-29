package models

import (
	"net/http"

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
		return http.StatusBadRequest, nil
	}
	if err := db.Create(order).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
