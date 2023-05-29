package service

import (
	"inventory-service/errors"
	"inventory-service/models"
	"net/http"
)

func AddItem(name string, description string, price float32, quantity uint) (uint32, *models.Item, error) {
	if name == "" || description == "" || price <= 0 || quantity == 0 {
		return http.StatusBadRequest, nil, errors.ErrEmptyField
	}
	newItem := &models.Item{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
	status, newItem, err := models.CreateItem(newItem)
	if err != nil {
		return status, nil, err
	}
	return status, newItem, nil
}

func GetItem(id uint) (uint32, *models.Item, error) {
	status, item, err := models.GetItem(id)
	if err != nil {
		return status, nil, err
	}
	return status, item, nil
}

func GetAllItems() (uint32, []*models.Item, error) {
	status, items, err := models.GetAllItems()
	if err != nil {
		return status, nil, err
	}
	return status, items, nil
}

func AddQuantity(id uint, quantity uint) (uint32, *models.Item, error) {
	status, item, err := models.GetItem(uint(id))
	if err != nil {
		return status, nil, err
	}
	item.Quantity += quantity
	status, err = models.UpdateItemQuantity(item.ID, item.Quantity)
	if err != nil {
		return status, nil, err
	}
	return http.StatusOK, item , nil
}

func LowerQuantity(id uint, quantity uint) (uint32, *models.Item, error) {
	status, item, err := models.GetItem(uint(id))
	if err != nil {
		return status, nil, err
	}
	if item.Quantity < quantity {
		return http.StatusConflict, nil, errors.ErrInsufficientQuantity
	}
	item.Quantity -= quantity
	status, err = models.UpdateItemQuantity(item.ID, item.Quantity)
	if err != nil {
		return status, nil, err
	}
	return status, item , nil
}

func DeleteItem(id uint) (uint32, error) {
	status, err := models.DeleteItem(id)
	if err != nil {
		return status, err
	}
	return status, nil
}