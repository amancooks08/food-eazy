package service

import (
	"context"
	"net/http"
	"order-service/errors"
	grpc "order-service/inventoryClient"
	"order-service/models"
	"time"
)

func PlaceOrder(userID uint32, itemID uint32, quantity uint32) (uint32, *models.Order, error) {
	if userID == 0 || itemID == 0 || quantity == 0 {
		return http.StatusBadRequest, nil, errors.ErrEmptyField
	}

	itemResponse, err := grpc.GetItem(context.Background(), itemID)
	if err != nil {
		return itemResponse.StatusCode, nil, err
	}

	if itemResponse.Quantity < quantity {
		return http.StatusUnprocessableEntity, nil, errors.ErrLimitedSupplies
	}

	order := &models.Order{
		UserID:    userID,
		ItemID:    itemID,
		Quantity:  quantity,
		Amount:    itemResponse.Price * float32(quantity),
		OrderTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	status, err := models.CreateOrder(order)
	if err != nil {
		return status, nil, err
	}

	updateResponse, err := grpc.LowerQuantity(context.Background(), itemID, quantity)
	if err != nil {
		return updateResponse.StatusCode, nil, err
	}

	return status, order, nil
}
