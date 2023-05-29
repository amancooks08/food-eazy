package dependencies

import (
	"api-gateway/clients"
	"api-gateway/proto/auth"
	"api-gateway/proto/inventory"
)

type Dependencies struct {
	AuthService auth.AuthServiceClient
	InventoryService inventory.InventoryServiceClient
}

func InitDependencies() *Dependencies {
	return &Dependencies{
		AuthService: clients.AuthServiceClient,
		InventoryService: clients.InventoryServiceClient,
	}
}
