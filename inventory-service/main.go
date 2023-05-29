package main

import (
	"inventory-service/inventoryServer"
	"inventory-service/config"
	"inventory-service/database"
	"inventory-service/proto/inventorypb"
	"net"

	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	config.Load()

	if err := database.InitDB(); err != nil {
		logger.WithField("error", err).Error("Error connecting to database")
		return
	}
	defer database.Close()

	gRPCServer := grpc.NewServer()

	inventorypb.RegisterInventoryServiceServer(gRPCServer, &inventoryServer.GRPCServer{})

	logger.Info("Starting gRPC server on port 33002")
	lis, err := net.Listen("tcp", ":33002")
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
