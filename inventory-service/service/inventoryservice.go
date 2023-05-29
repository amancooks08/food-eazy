package service

import (
	"inventory-service/errors"
	"inventory-service/models"
)

func AddItem(name string, description string, price float64, quantity uint) (*models.Item, error) {
	if name == "" || description == "" || price == 0 || quantity == 0 {
		return nil, errors.ErrEmptyField
	}
	newItem := models.Item{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	err := models.CreateItem(newItem)
	if err != nil {
		return nil, err
	}
	return &newItem, nil
}
