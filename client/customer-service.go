package client

import (
	"github.com/tittuvarghese/core/logger"
	"github.com/tittuvarghese/customer-service/proto"
	"github.com/tittuvarghese/gateway/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logger.NewLogger("customer-service-client")

var CustomerService proto.AuthServiceClient

func NewCustomerServiceClient() *proto.AuthServiceClient {
	// Establish a connection to the gRPC server
	conn, err := grpc.NewClient(constants.CustomerServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Failed to connect to gRPC server", err)
	}
	log.Info("Successfully connected to customer-service gRPC server")

	// Create a new client for the AuthService
	CustomerService = proto.NewAuthServiceClient(conn)

	// defer conn.Close()
	return &CustomerService
}
