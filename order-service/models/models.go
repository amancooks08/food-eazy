package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Order struct {
	gorm.Model
	OrderID   uint32  `gorm:"column:order_id; primaryKey; autoIncrement"`
	UserID    uint32  `gorm:"column:user_id; not null"`
	ItemID    uint32  `gorm:"column:item_id; not null"`
	Quantity  uint32  `gorm:"column:quantity; not null"`
	Amount    float32 `gorm:"column:amount; not null"`
	OrderTime 	string  `gorm:"column:order_time; not null"`
}

func InitOrderModels(database *gorm.DB) {
	db = database
	db.AutoMigrate(&Order{})
}
