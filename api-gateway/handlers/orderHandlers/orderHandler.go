package orderHandlers

import (
	"api-gateway/domain"
	proto "api-gateway/proto/order"
	"encoding/json"
	"fmt"
	"net/http"
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
