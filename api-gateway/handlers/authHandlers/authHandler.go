package authHandlers

import (
	"api-gateway/domain"
	proto "api-gateway/proto/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(authService proto.AuthServiceClient) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Println(req.Body)
		var requestBody domain.RegisterUserRequest

		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("invalid request body: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(message)
			return
		}

		grpcRequest := proto.RegisterUserRequest{
			Name:        requestBody.Name,
			Email:       requestBody.Email,
			Password:    requestBody.Password,
			PhoneNumber: requestBody.PhoneNumber,
			Role:        domain.RoleMap[requestBody.Role],
		} 

		resp, err := authService.RegisterUser(req.Context(), &grpcRequest)
		if err != nil {
			message := domain.Message{
				Message: fmt.Sprintf("grpc received error: %s", err.Error()),
			}
			rw.WriteHeader(http.StatusInternalServerError)
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

		rw.WriteHeader(http.StatusCreated)
		rw.Header().Set("Content-Type", "application/json")
		rw.Write(res)
	})
}