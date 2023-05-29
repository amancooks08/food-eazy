package main

import (
	"api-gateway/clients"
	"api-gateway/dependencies"
	"api-gateway/routes"

	"github.com/urfave/negroni"
)

func main() {
	clients.InitClients()

	deps := dependencies.InitDependencies()
	router := routes.InitRouter(deps)
	
	server := negroni.Classic()
	server.UseHandler(router)

	server.Run(":8002")
}
