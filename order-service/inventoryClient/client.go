package inventoryClient

import (
	proto "order-service/proto/inventorypb"

	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var InventoryServiceClient proto.InventoryServiceClient

func InitGRPCClient() {
	conn, err := grpc.Dial("localhost:33002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to connect to inventory service")
		return
	}

	logger.Info("Connecting to inventory service server...")
	InventoryServiceClient = proto.NewInventoryServiceClient(conn)
}

