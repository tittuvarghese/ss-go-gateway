package main

import (
	"github.com/tittuvarghese/ss-go-core/config"
	"github.com/tittuvarghese/ss-go-core/jwt"
	"github.com/tittuvarghese/ss-go-core/logger"
	"github.com/tittuvarghese/ss-go-core/otel"
	"github.com/tittuvarghese/ss-go-gateway/client"
	"github.com/tittuvarghese/ss-go-gateway/constants"
	"github.com/tittuvarghese/ss-go-gateway/core/handler"
	"github.com/tittuvarghese/ss-go-gateway/service"
)

func main() {

	log := logger.NewLogger(constants.ModuleName)
	log.Info("Initialising Gateway Module")

	// Config Management
	configManager := config.NewConfigManager(config.DEFAULT_CONFIG_PATH)
	configManager.Enable()

	if configManager.GetBool(constants.OtelEnableEnv) {
		serviceName := configManager.GetString(constants.OtelServiceNameEnv)
		collectorUrl := configManager.GetString(constants.OtelCollectorEnv)
		insecureMode := configManager.GetBool(constants.OtelInsecureModeEnv)
		otel.NewTraceProvider(serviceName, collectorUrl, insecureMode)
	}

	server := handler.NewHttpSerer()
	server.EnableLogger()
	server.EnableRecovery()
	server.EnableRateLimiter()
	server.EnableTelemetry()

	// Handlers
	server.AddHandler(constants.HttpGet, constants.GatewayServicePath, "/status", service.Status)
	server.AddHandler(constants.HttpPost, constants.CustomerServicePath, "/register", service.Register)
	server.AddHandler(constants.HttpPost, constants.CustomerServicePath, "/login", service.Login)
	server.AddHandler(constants.HttpGet, constants.CustomerServicePath, "/profile", jwt.Authorize(), service.GetProfile)
	server.AddHandler(constants.HttpPost, constants.ProductServicePath, "/create", jwt.Authorize(), service.CreateProduct)
	server.AddHandler(constants.HttpGet, constants.ProductServicePath, "/product/:productId", jwt.Authorize(), service.GetProduct)
	server.AddHandler(constants.HttpGet, constants.ProductServicePath, "/products", jwt.Authorize(), service.GetProducts)
	server.AddHandler(constants.HttpPost, constants.ProductServicePath, "/product/:productId", jwt.Authorize(), service.UpdateProduct)
	server.AddHandler(constants.HttpGet, constants.OrderServicePath, "/orders", jwt.Authorize(), service.GetOrders)
	server.AddHandler(constants.HttpPost, constants.OrderServicePath, "/order", jwt.Authorize(), service.CreateOrder)
	server.AddHandler(constants.HttpGet, constants.OrderServicePath, "/order/:orderId", jwt.Authorize(), service.GetOrder)
	server.AddHandler(constants.HttpPost, constants.OrderServicePath, "/order/:orderId", jwt.Authorize(), service.UpdateOrder)

	customerServiceEndpoint := configManager.GetString(constants.CustomerServiceAddressEnv)
	if customerServiceEndpoint == "" {
		customerServiceEndpoint = constants.CustomerServiceAddress
	}

	productServiceEndpoint := configManager.GetString(constants.ProductServiceAddressEnv)
	if productServiceEndpoint == "" {
		productServiceEndpoint = constants.ProductServiceAddress
	}

	orderServiceEndpoint := configManager.GetString(constants.OrderServiceAddressEnv)
	if orderServiceEndpoint == "" {
		orderServiceEndpoint = constants.OrderServiceAddress
	}

	// Client Connections
	client.NewCustomerServiceClient(customerServiceEndpoint)
	client.NewProductServiceClient(productServiceEndpoint)
	client.NewOrderManagementServiceClient(orderServiceEndpoint)

	server.Run(constants.HttpServerPort)
}
