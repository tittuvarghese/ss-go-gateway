package client

import (
	"github.com/tittuvarghese/ss-go-product-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ProductService proto.ProductServiceClient

func NewProductServiceClient(endpoint string) *proto.ProductServiceClient {
	// Establish a connection to the gRPC server
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("Failed to connect to gRPC server", err)
	}
	log.Info("Successfully connected to product-service gRPC server")

	// Create a new client for the AuthService
	ProductService = proto.NewProductServiceClient(conn)

	// defer conn.Close()
	return &ProductService
}
