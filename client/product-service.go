package client

import (
	"github.com/tittuvarghese/gateway/constants"
	"github.com/tittuvarghese/product-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProductService proto.ProductServiceClient

func NewProductServiceClient() *proto.ProductServiceClient {
	// Establish a connection to the gRPC server
	conn, err := grpc.NewClient(constants.ProductServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Failed to connect to gRPC server", err)
	}
	log.Info("Successfully connected to product-service gRPC server")

	// Create a new client for the AuthService
	ProductService = proto.NewProductServiceClient(conn)

	// defer conn.Close()
	return &ProductService
}
