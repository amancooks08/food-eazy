package main

import (
	"net"
	"order-service/config"
	"order-service/database"
	"order-service/inventoryClient"
	"order-service/proto/orderpb"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"order-service/orderServer"
)

func main() {
	inventoryClient.InitGRPCClient()
	config.Load()

	if err := database.InitDB(); err != nil {
		logger.WithField("error", err).Error("Error connecting to database")
		return
	}
	defer database.Close()

	gRPCServer := grpc.NewServer()

	orderpb.RegisterOrderServiceServer(gRPCServer, orderServer.NewGRPCServer(inventoryClient.InventoryServiceClient))

	logger.Info("Starting gRPC server on port 33003")
	lis, err := net.Listen("tcp", ":33003")
	if err != nil {
		logger.WithField("error", err).Error("Error starting inventory gRPC server")
		return
	}

	if err := gRPCServer.Serve(lis); err != nil {
		logger.WithField("error", err).Error("Error starting inventory gRPC server")
		return
	}

	defer lis.Close()
	defer gRPCServer.Stop()
}
