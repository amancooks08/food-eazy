package models

import "gorm.io/gorm"

var db *gorm.DB

type User struct {
	gorm.Model
	Name string 			`gorm:"column:name; unique; primarykey; AUTO_INCREMENT;"`
	Email string 			`gorm:"column:email; unique; not null"`
	Password string 		`gorm:"not null"`
	PhoneNumber string 		`gorm:"unique; not null"`
}

func InitAuthModels(database *gorm.DB) {
	db = database
	db.AutoMigrate(&User{})
}

