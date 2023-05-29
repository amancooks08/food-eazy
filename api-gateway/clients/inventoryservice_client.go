package clients

import (
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	proto "api-gateway/proto/inventory"
	"google.golang.org/grpc/credentials/insecure"
)

var InventoryServiceClient proto.InventoryServiceClient

func initInventoryClient() {
	conn, err := grpc.Dial("localhost:33002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to connect to auth service")
		return
	}

	logger.Info("Connecting to auth service server...")
	InventoryServiceClient = proto.NewInventoryServiceClient(conn)
}
