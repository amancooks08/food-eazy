package routes

import (
	"api-gateway/handlers/inventoryHandlers"
	inventoryproto "api-gateway/proto/inventory"

	"github.com/gorilla/mux"
)

//InitInventoryRoutes initializes the routes for the inventory service

func InitInventoryRoutes(router *mux.Router, inventoryService inventoryproto.InventoryServiceClient) {
	router.HandleFunc("/admin/inventory/item/add", authMiddleware(inventoryHandlers.AddItem(inventoryService))).Methods("POST")
	router.HandleFunc("/inventory/item", inventoryHandlers.GetItem(inventoryService)).Methods("POST")
	router.HandleFunc("/inventory/item/all", inventoryHandlers.GetAllItems(inventoryService)).Methods("POST")
	router.HandleFunc("/admin/inventory/item/quantity/add", authMiddleware(inventoryHandlers.AddQuantity(inventoryService))).Methods("POST")
	router.HandleFunc("/admin/inventory/item/quantity/remove", authMiddleware(inventoryHandlers.LowerQuantity(inventoryService))).Methods("POST")
	router.HandleFunc("/admin/inventory/item/remove", authMiddleware(inventoryHandlers.DeleteItem(inventoryService))).Methods("POST")
}
