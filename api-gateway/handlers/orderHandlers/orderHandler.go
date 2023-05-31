package orderHandlers

import (
	"api-gateway/domain"
	"api-gateway/errors"
	proto "api-gateway/proto/order"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func PlaceOrder(orderService proto.OrderServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var requestBody domain.PlaceOrderRequest

		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("invalid request body: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(message)
			return
		}
		if id := req.Context().Value("id"); id == nil {
			message := domain.Message{
				Message: "unauthorized access: invalid user id",
			}
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(message)
			return
		}
		id := req.Context().Value("id").(int)

		if id == 0 {
			message := domain.Message{
				Message: "unauthorized access: invalid user id",
			}
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(message)
			return
		}

		userID := uint32(id)

		grpcRequest := proto.PlaceOrderRequest{
			UserId:   userID,
			ItemId:   requestBody.ItemID,
			Quantity: requestBody.Quantity,
		}

		resp, err := orderService.PlaceOrder(req.Context(), &grpcRequest)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		response := domain.PlaceOrderResponse{
			OrderID:   resp.Order.OrderId,
			Amount:    resp.Order.Amount,
			OrderTime: resp.Order.OrderTime,
		}

		res, err := json.Marshal(response)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("error marshalling response: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(message)
			return
		}

		rw.WriteHeader(int(resp.StatusCode))
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(res)
	})
}

func GetOrders(orderService proto.OrderServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if id := req.Context().Value("id"); id == nil {
			message := domain.Message{
				Message: "unauthorized access: invalid user id",
			}
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(message)
			return
		}
		id := req.Context().Value("id").(int)

		if id == 0 {
			message := domain.Message{
				Message: "unauthorized access: invalid user id",
			}
			rw.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(rw).Encode(message)
			return
		}

		userID := uint32(id)

		// Check If a specific order is requested, if yes than return that order only
		// otherwise return all orders of the user
		orderID := req.URL.Query().Get("id")
		if orderID != "" {
			id, err := strconv.Atoi(orderID)
			if err != nil {
				message := domain.Message{
					Message: fmt.Sprintf("invalid order id: %s", err.Error()),
				}
				rw.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(rw).Encode(message)
				return
			}
			grpcRequest := proto.GetOrderRequest{
				OrderId: uint32(id),
			}

			resp, err := orderService.GetOrder(req.Context(), &grpcRequest)
			if err != nil {
				message := domain.Message{
					Message: fmt.Sprintf("grpc received error: %s", err.Error()),
				}
				rw.WriteHeader(int(resp.StatusCode))
				json.NewEncoder(rw).Encode(message)
				return
			}

			if resp.Order.UserId != userID {
				message := domain.Message{
					Message: "unauthorized access: invalid user id",
				}
				rw.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(rw).Encode(message)
				return
			}

			response := domain.GetOrderResponse{
				Order: &domain.Order{
					OrderID:   resp.Order.OrderId,
					UserID:    resp.Order.UserId,
					ItemID:    resp.Order.ItemId,
					Quantity:  resp.Order.Quantity,
					Amount:    resp.Order.Amount,
					OrderTime: resp.Order.OrderTime,
				},
			}

			res, err := json.Marshal(response)
			if err != nil {
				message := domain.Message{
					Message: fmt.Sprintf("error marshalling response: %s", err.Error()),
				}
				rw.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(rw).Encode(message)
				return
			}

			rw.WriteHeader(int(resp.StatusCode))
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(res)
			return
		}

		grpcRequest := proto.GetAllOrdersRequest{
			UserId: userID,
		}

		resp, err := orderService.GetAllOrders(req.Context(), &grpcRequest)
		if err == errors.ErrNoOrdersAvailable{
			message := domain.Message{
				Message: fmt.Sprintf("no orders available for user: %d", userID),
			}
			rw.WriteHeader(http.StatusNotFound)
			json.NewEncoder(rw).Encode(message)
			return
		} else if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		orders := make([]*domain.Order, 0)

		for _, order := range resp.Orders {
			orders = append(orders, &domain.Order{
				OrderID:   order.OrderId,
				UserID:    order.UserId,
				ItemID:    order.ItemId,
				Quantity:  order.Quantity,
				Amount:    order.Amount,
				OrderTime: order.OrderTime,
			})
		}
		response := domain.GetAllOrdersResponse{
			Orders: orders,
		}

		res, err := json.Marshal(response)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("error marshalling response: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(rw).Encode(message)
			return
		}

		rw.WriteHeader(int(resp.StatusCode))
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(res)
	})
}
