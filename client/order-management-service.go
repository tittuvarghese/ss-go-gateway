package client

import (
	"github.com/tittuvarghese/gateway/constants"
	"github.com/tittuvarghese/order-management-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var OrderManagementService proto.OrderServiceClient

func NewOrderManagementServiceClient() *proto.OrderServiceClient {
	// Establish a connection to the gRPC server
	conn, err := grpc.NewClient(constants.OrderServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Failed to connect to gRPC server", err)
	}
	log.Info("Successfully connected to product-service gRPC server")

	// Create a new client for the AuthService
	OrderManagementService = proto.NewOrderServiceClient(conn)

	// defer conn.Close()
	return &OrderManagementService
}
