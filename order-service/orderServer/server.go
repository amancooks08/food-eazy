package orderServer

import (
	"context"

	proto "order-service/proto/orderpb"
	client "order-service/proto/inventorypb"
	"order-service/service"
)

type GRPCServer struct {
	*service.OrderService
	proto.UnimplementedOrderServiceServer
}

func NewGRPCServer(Client client.InventoryServiceClient) *GRPCServer {
	service := service.NewOrderService(Client)
	return &GRPCServer{
		service,
		proto.UnimplementedOrderServiceServer{},
	}
}

func (s *GRPCServer) PlaceOrder(ctx context.Context, req *proto.PlaceOrderRequest) (*proto.PlaceOrderResponse, error) {
	statusCode, order, err := s.OrderService.PlaceOrder(req.UserId, req.ItemId, req.Quantity)
	if err != nil {
		return &proto.PlaceOrderResponse{
			StatusCode: statusCode,
			Order:      nil,
		}, err
	}

	return &proto.PlaceOrderResponse{
		StatusCode: statusCode,
		Order: &proto.Order{
			OrderId:  order.ID,
			UserId:   order.UserID,
			ItemId:   order.ItemID,
			Quantity: order.Quantity,
			Amount:   order.Amount,
		},
	}, nil
}