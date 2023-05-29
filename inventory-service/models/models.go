package models

import "gorm.io/gorm"

var db *gorm.DB

type Item struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey; column:id; autoIncrement; not null"`
	Name        string `gorm:"column:name; not null"`
	Description string `gorm:"column:description; not null"`
	Price       float64   `gorm:"column:price; not null"`
	Quantity    uint   `gorm:"column:quantity; not null"`
}

func InitInventoryModels(database *gorm.DB) {
	db = database
	db.AutoMigrate(&Item{})
}

func CreateItem(item Item) error {
	err := db.Create(&item).Error
	if err != nil {
		return err
	}
	return nil
}
