package routes

import (
	"api-gateway/handlers/inventoryHandlers"
	proto "api-gateway/proto/auth"
	"github.com/gorilla/mux"
)


//InitInventoryRoutes initializes the routes for the inventory service

func InitInventoryRoutes(router *mux.Router, authService proto.AuthServiceClient) {
	router.HandleFunc("/admin/inventory/item/add", authMiddleware(inventoryHandlers.AddItem(authService))).Methods("POST")
}