package clients

import (
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	proto "api-gateway/proto/auth"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthServiceClient proto.AuthServiceClient

func initAuthClient() {
	conn, err := grpc.Dial("localhost:33001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to connect to auth service")
		return
	}

	logger.Info("Connecting to auth service server...")
	AuthServiceClient = proto.NewAuthServiceClient(conn)
}
