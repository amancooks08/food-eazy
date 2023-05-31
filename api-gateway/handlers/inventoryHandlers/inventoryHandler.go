package inventoryHandlers

import (
	"api-gateway/domain"
	proto "api-gateway/proto/inventory"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
			ID:          resp.Id,
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

func GetItem(inventoryService proto.InventoryServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var requestBody proto.GetItemRequest

		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("invalid request body: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(message)
			return
		}

		grpcRequest := proto.GetItemRequest{
			Id: requestBody.Id,
		}

		resp, err := inventoryService.GetItem(req.Context(), &grpcRequest)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		response := domain.GetItemResponse{
			ID:          resp.Id,
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

func GetAllItems(inventoryService proto.InventoryServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		resp, err := inventoryService.GetAllItems(req.Context(), &proto.GetAllItemsRequest{})
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		var response domain.GetAllItemsResponse

		for _, item := range resp.Items {
			response.Items = append(response.Items, domain.GetItemResponse{
				ID:          item.Id,
				Name:        item.Name,
				Description: item.Description,
				Quantity:    item.Quantity,
				Price:       item.Price,
			})
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

func AddQuantity(inventoryService proto.InventoryServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var requestBody domain.UpdateQuantityRequest

		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("invalid request body: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(message)
			return
		}

		grpcRequest := proto.AddQuantityRequest{
			Id:       requestBody.ID,
			Quantity: requestBody.Quantity,
		}

		resp, err := inventoryService.AddQuantity(req.Context(), &grpcRequest)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		response := domain.UpdateQuantityResponse{
			ID:       resp.Id,
			Quantity: resp.Quantity,
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

func LowerQuantity(inventoryService proto.InventoryServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var requestBody domain.UpdateQuantityRequest

		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("invalid request body: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(message)
			return
		}

		grpcRequest := proto.LowerQuantityRequest{
			Id:       requestBody.ID,
			Quantity: requestBody.Quantity,
		}

		resp, err := inventoryService.LowerQuantity(req.Context(), &grpcRequest)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		response := domain.UpdateQuantityResponse{
			ID:       resp.Id,
			Quantity: resp.Quantity,
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

func DeleteItem(inventoryService proto.InventoryServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		queryParams, err := url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			// Handle the error if parsing fails
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		id := queryParams.Get("id")
		itemID, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			// Handle the error if conversion fails
			http.Error(rw, "Invalid ID", http.StatusBadRequest)
			return
		}
		grpcRequest := &proto.DeleteItemRequest{
			Id: int32(itemID),
		}

		resp, err := inventoryService.DeleteItem(req.Context(), grpcRequest)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(int(resp.StatusCode))
			json.NewEncoder(rw).Encode(message)
			return
		}

		response := domain.Message{
			Message: resp.Message,
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
