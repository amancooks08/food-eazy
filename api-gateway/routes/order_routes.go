package routes

import (
	"api-gateway/handlers/orderHandlers"

	"github.com/gorilla/mux"
	proto "api-gateway/proto/order"
)

func InitOrderRoutes(router *mux.Router, orderService proto.OrderServiceClient) {
	router.Handle("/order/place", authMiddleware(orderHandlers.PlaceOrder(orderService))).Methods("POST")
}