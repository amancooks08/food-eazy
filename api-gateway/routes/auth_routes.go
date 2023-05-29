package routes

import (
	"api-gateway/handlers/authHandlers"
	proto "api-gateway/proto/auth"
	"github.com/gorilla/mux"
)

func InitAuthRoutes(router *mux.Router, authService proto.AuthServiceClient) {
	router.HandleFunc("/register", authHandlers.RegisterUser(authService)).Methods("POST")
	router.HandleFunc("/login", authHandlers.LoginUser(authService)).Methods("POST")
}
