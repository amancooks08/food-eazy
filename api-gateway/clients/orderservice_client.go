package clients


import (
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	proto "api-gateway/proto/order"
	"google.golang.org/grpc/credentials/insecure"
)

var OrderServiceClient proto.OrderServiceClient

func initOrderClient() {
	conn, err := grpc.Dial("localhost:33003", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to connect to auth service")
		return
	}

	logger.Info("Connecting to auth service server...")
	OrderServiceClient = proto.NewOrderServiceClient(conn)
}
