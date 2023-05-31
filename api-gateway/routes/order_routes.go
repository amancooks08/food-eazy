package routes

import (
	"api-gateway/handlers/orderHandlers"

	"github.com/gorilla/mux"
	proto "api-gateway/proto/order"
)

func InitOrderRoutes(router *mux.Router, orderService proto.OrderServiceClient) {
	router.HandleFunc("/user/order", authMiddleware(orderHandlers.PlaceOrder(orderService))).Methods("POST")
	router.HandleFunc("/user/order", authMiddleware(orderHandlers.GetOrders(orderService))).Methods("GET")
}