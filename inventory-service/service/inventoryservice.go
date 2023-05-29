package service

import (
	"inventory-service/errors"
	"inventory-service/models"
)

func AddItem(name string, description string, price float64, quantity uint) (*models.Item, error) {
	if name == "" || description == "" || price == 0 || quantity == 0 {
		return nil, errors.ErrEmptyField
	}
	newItem := &models.Item{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	newItem, err := models.CreateItem(newItem)
	if err != nil {
		return nil, err
	}
	return newItem, nil
}

func GetItem(id uint) (*models.Item, error) {
	item, err := models.GetItem(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}


func GetAllItems() ([]*models.Item, error) {
	items, err := models.GetAllItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}