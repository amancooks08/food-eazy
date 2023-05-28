package routes

import (
	"api-gateway/dependencies"

	mux "github.com/gorilla/mux"
)

func InitRouter(deps *dependencies.Dependencies) *mux.Router {
	router := mux.NewRouter()

	InitAuthRoutes(router, deps.AuthService)
	return router
}
