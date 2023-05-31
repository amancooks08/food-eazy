package service

import (
	"context"
	"net/http"
	"order-service/errors"
	"order-service/models"
	"time"
	grpc "order-service/inventoryClient"
	proto "order-service/proto/inventorypb"
)

type OrderService struct {
	client proto.InventoryServiceClient
}

func NewOrderService(Client proto.InventoryServiceClient) *OrderService {
	return &OrderService{
		client: grpc.InventoryServiceClient,
	}
}

func (service *OrderService) PlaceOrder(userID uint32, itemID uint32, quantity uint32) (uint32, *models.Order, error) {
	if userID == 0 || itemID == 0 || quantity == 0 {
		return http.StatusBadRequest, nil, errors.ErrEmptyField
	}

	itemResponse, err := service.client.GetItem(context.Background(), &proto.GetItemRequest{Id: itemID})
	if itemResponse == nil {
		return http.StatusBadGateway, nil, errors.ErrBadGateway
	}
	
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
		Amount:    float32(quantity) * float32(itemResponse.Price),
		OrderTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	status, err := models.CreateOrder(order)
	if err != nil {
		return status, nil, err
	}

	updateResponse, err := service.client.LowerQuantity(context.Background(), &proto.LowerQuantityRequest{Id: itemID, Quantity: quantity})
	if err != nil {
		return updateResponse.StatusCode, nil, err
	}

	return status, order, nil
}

func (service *OrderService) GetOrder(orderID uint32) (uint32, *models.Order, error) {
	if orderID == 0 {
		return http.StatusBadRequest, nil, errors.ErrEmptyField
	}

	status, order, err := models.GetOrder(orderID)
	if err != nil {
		return status, nil, err
	}

	return http.StatusOK, order, nil
}

func (service *OrderService) GetAllOrders(userID uint32) (uint32, []*models.Order, error) {
	status, orders, err := models.GetAllOrders(userID)
	if err != nil {
		return status, nil, err
	}

	if len(orders) == 0 {
		return http.StatusOK, orders, errors.ErrNoOrdersAvailable
	}

	return http.StatusOK, orders, nil
}