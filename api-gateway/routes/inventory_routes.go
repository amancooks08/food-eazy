package routes

import (
	"api-gateway/handlers/inventoryHandlers"
	inventoryproto "api-gateway/proto/inventory"

	"github.com/gorilla/mux"
)

//InitInventoryRoutes initializes the routes for the inventory service

func InitInventoryRoutes(router *mux.Router, inventoryService inventoryproto.InventoryServiceClient) {
	router.HandleFunc("/admin/inventory/item/add", authMiddleware(inventoryHandlers.AddItem(inventoryService))).Methods("POST")
}
