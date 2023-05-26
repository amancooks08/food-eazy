package main

import (
	"auth-service/authServer"
	"auth-service/config"
	"auth-service/database"
	"auth-service/proto/authpb"
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

	authpb.RegisterAuthServiceServer(gRPCServer, &authServer.GRPCServer{})

	logger.Info("Starting gRPC server on port 33001")
	lis, err := net.Listen("tcp", ":33001")
	if err != nil {
		logger.WithField("error", err).Error("Error starting gRPC server")
		return
	}

	if err := gRPCServer.Serve(lis); err != nil {
		logger.WithField("error", err).Error("Error starting gRPC server")
		return
	}

	defer lis.Close()
	defer gRPCServer.Stop()
}
