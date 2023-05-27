package routes

import (
	mux "github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	InitAuthRoutes(router)
	return router
}
