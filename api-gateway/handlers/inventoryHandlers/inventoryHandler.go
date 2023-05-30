package inventoryHandlers

import (
	"api-gateway/domain"
	proto "api-gateway/proto/inventory"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddItem(inventoryService proto.InventoryServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var requestBody proto.AddItemRequest

		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("invalid request body: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(message)
			return
		}

		grpcRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		resp, err := inventoryService.AddItem(req.Context(), &grpcRequest)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		response := domain.AddItemResponse{
			ID:		  resp.Id,
			Name:        resp.Name,
			Description: resp.Description,
			Quantity:    resp.Quantity,
			Price:       resp.Price,
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
