package dependencies

import (
	"api-gateway/clients"
	"api-gateway/proto/auth"
	"api-gateway/proto/inventory"
	"api-gateway/proto/order"
)

type Dependencies struct {
	AuthService auth.AuthServiceClient
	InventoryService inventory.InventoryServiceClient
	OrderService order.OrderServiceClient
}

func InitDependencies() *Dependencies {
	return &Dependencies{
		AuthService: clients.AuthServiceClient,
		InventoryService: clients.InventoryServiceClient,
		OrderService: clients.OrderServiceClient,
	}
}
