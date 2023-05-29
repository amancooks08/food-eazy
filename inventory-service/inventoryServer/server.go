package inventoryServer

import (
	"context"
	proto "inventory-service/proto/inventorypb"
	"inventory-service/service"
)

type GRPCServer struct {
	proto.UnimplementedInventoryServiceServer
}

func (s *GRPCServer) AddItem(ctx context.Context, req *proto.AddItemRequest) (*proto.AddItemResponse, error) {
	status, item, err := service.AddItem(req.Name, req.Description, req.Price, uint(req.Quantity))
	if err != nil {
		return &proto.AddItemResponse{
			StatusCode:  int32(status),
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}, err
	}

	return &proto.AddItemResponse{
		StatusCode:  int32(status),
		Id:          int32(item.ID),
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    uint32(item.Quantity),
	}, nil
}

func (s *GRPCServer) GetItem(ctx context.Context, req *proto.GetItemRequest) (*proto.GetItemResponse, error) {
	status, item, err := service.GetItem(uint(req.Id))
	if err != nil {
		return &proto.GetItemResponse{
			StatusCode:  int32(status),
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}, err
	}

	return &proto.GetItemResponse{
		StatusCode:  int32(status),
		Id:          int32(item.ID),
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Quantity:    uint32(item.Quantity),
	}, nil
}

func (s *GRPCServer) GetAllItems(ctx context.Context, req *proto.GetAllItemsRequest) (*proto.GetAllItemsResponse, error) {
	status, items, err := service.GetAllItems()
	if err != nil {
		return &proto.GetAllItemsResponse{
			StatusCode: int32(status),
			Items:      nil,
		}, err
	}

	var itemsResponse []*proto.GetItemResponse
	for _, item := range items {
		itemsResponse = append(itemsResponse, &proto.GetItemResponse{
			Id:          int32(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Quantity:    uint32(item.Quantity),
		})
	}

	return &proto.GetAllItemsResponse{
		StatusCode: int32(status),
		Items:      itemsResponse,
	}, nil
}

func (s *GRPCServer) AddQuantity(ctx context.Context, req *proto.AddQuantityRequest) (*proto.AddQuantityResponse, error) {
	status, item, err := service.AddQuantity(uint(req.Id), uint(req.Quantity))
	if err != nil {
		return &proto.AddQuantityResponse{
			StatusCode: int32(status),
			Id:         0,
			Quantity:   0,
		}, err
	}

	return &proto.AddQuantityResponse{
		StatusCode: int32(status),
		Id:         int32(item.ID),
		Quantity:   uint32(item.Quantity),
	}, nil
}

func (s *GRPCServer) LowerQuantity(ctx context.Context, req *proto.LowerQuantityRequest) (*proto.LowerQuantityResponse, error) {
	status, item, err := service.LowerQuantity(uint(req.Id), uint(req.Quantity))
	if err != nil {
		return &proto.LowerQuantityResponse{
			StatusCode: int32(status),
			Id:         0,
			Quantity:   0,
		}, err
	}

	return &proto.LowerQuantityResponse{
		StatusCode: int32(status),
		Id:         int32(item.ID),
		Quantity:   uint32(item.Quantity),
	}, nil
}

func (s *GRPCServer) DeleteItem(ctx context.Context, req *proto.DeleteItemRequest) (*proto.DeleteItemResponse, error) {
	status, err := service.DeleteItem(uint(req.Id))
	if err != nil {
		return &proto.DeleteItemResponse{
			StatusCode: int32(status),
		}, err
	}

	return &proto.DeleteItemResponse{
		StatusCode: int32(status),
	}, nil
}