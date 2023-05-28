package dependencies

import (
	"api-gateway/clients"
	proto "api-gateway/proto/auth"
)

type Dependencies struct {
	AuthService proto.AuthServiceClient
}

func InitDependencies() *Dependencies {
	return &Dependencies{
		AuthService: clients.AuthServiceClient,
	}
}
