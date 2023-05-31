package orderServer

import (
	"context"

	client "order-service/proto/inventorypb"
	proto "order-service/proto/orderpb"
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

func (s *GRPCServer) GetOrder(ctx context.Context, req *proto.GetOrderRequest) (*proto.GetOrderResponse, error) {
	statusCode, order, err := s.OrderService.GetOrder(req.OrderId)
	if err != nil {
		return &proto.GetOrderResponse{
			StatusCode: statusCode,
			Order:      nil,
		}, err
	}

	return &proto.GetOrderResponse{
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

func (s *GRPCServer) GetAllOrders(ctx context.Context, req *proto.GetAllOrdersRequest) (*proto.GetAllOrdersResponse, error) {
	statusCode, orders, err := s.OrderService.GetAllOrders(req.UserId)
	if err != nil {
		return &proto.GetAllOrdersResponse{
			StatusCode: statusCode,
			Orders:     nil,
		}, err
	}

	var ordersResponse []*proto.Order
	for _, order := range orders {
		ordersResponse = append(ordersResponse, &proto.Order{
			OrderId:  order.ID,
			UserId:   order.UserID,
			ItemId:   order.ItemID,
			Quantity: order.Quantity,
			Amount:   order.Amount,
		})
	}

	return &proto.GetAllOrdersResponse{
		StatusCode: statusCode,
		Orders:     ordersResponse,
	}, nil
}
