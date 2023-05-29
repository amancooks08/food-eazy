package orderServer

import (
	"context"

	proto "order-service/proto/orderpb"
	"order-service/service"
)

type GRPCServer struct {
	proto.UnimplementedOrderServiceServer
}

func (s *GRPCServer) PlaceOrder(ctx context.Context, req *proto.PlaceOrderRequest) (*proto.PlaceOrderResponse, error) {
	statusCode, order, err := service.PlaceOrder(req.UserId, req.ItemId, req.Quantity)
	if err != nil {
		return nil, err
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
