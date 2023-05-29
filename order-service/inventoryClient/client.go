package inventoryClient

import (
	"context"
	proto "order-service/proto/inventorypb"

	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var InventoryServiceClient proto.InventoryServiceClient

func InitGRPCClient(grpcServer proto.InventoryServiceClient) {
	conn, err := grpc.Dial("localhost:33002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		logger.WithField("error", err).Fatal("Failed to connect to auth service")
		return
	}

	logger.Info("Connecting to auth service server...")
	InventoryServiceClient = proto.NewInventoryServiceClient(conn)

}


func GetItem(ctx context.Context, itemID uint32) (*proto.GetItemResponse, error) {
	item, err := InventoryServiceClient.GetItem(ctx, &proto.GetItemRequest{
		Id : itemID,
	})
	if err != nil {
		return nil, err
	}
	
	return item, nil
}

func LowerQuantity(ctx context.Context, itemID uint32, quantity uint32) (*proto.LowerQuantityResponse, error) {
	item, err := InventoryServiceClient.LowerQuantity(ctx, &proto.LowerQuantityRequest{
		Id : itemID,
		Quantity : quantity,
	})
	if err != nil {
		return nil, err
	}
	return item, nil
}