package client

import (
	"github.com/tittuvarghese/ss-go-core/logger"
	"github.com/tittuvarghese/ss-go-customer-service/proto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logger.NewLogger("gateway-client")

var CustomerService proto.AuthServiceClient

func NewCustomerServiceClient(endpoint string) *proto.AuthServiceClient {
	// Establish a connection to the gRPC server
	conn, err := grpc.NewClient(
		endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Error("Failed to connect to gRPC server", err)
	}
	log.Info("Successfully connected to customer-service gRPC server")

	// Create a new client for the AuthService
	CustomerService = proto.NewAuthServiceClient(conn)
	// defer conn.Close()
	return &CustomerService
}
