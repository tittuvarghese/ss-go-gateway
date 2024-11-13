package main

import (
	"github.com/tittuvarghese/core/config"
	"github.com/tittuvarghese/core/jwt"
	"github.com/tittuvarghese/core/logger"
	"github.com/tittuvarghese/gateway/client"
	"github.com/tittuvarghese/gateway/constants"
	"github.com/tittuvarghese/gateway/core/handler"
	"github.com/tittuvarghese/gateway/service"
)

func main() {
	log := logger.NewLogger(constants.ModuleName)
	log.Info("Initialising Gateway Module")

	// Config Management
	configManager := config.NewConfigManager(config.DEFAULT_CONFIG_PATH)
	configManager.Enable()

	server := handler.NewHttpSerer()
	server.EnableLogger()
	server.EnableRecovery()
	server.EnableRateLimiter()
	// Handlers
	server.AddHandler(constants.HttpGet, constants.GatewayServicePath, "/status", service.Status)
	server.AddHandler(constants.HttpPost, constants.CustomerServicePath, "/register", service.Register)
	server.AddHandler(constants.HttpPost, constants.CustomerServicePath, "/login", service.Login)
	server.AddHandler(constants.HttpGet, constants.CustomerServicePath, "/profile", jwt.Authorize(), service.GetProfile)
	server.AddHandler(constants.HttpPost, constants.ProductServicePath, "/create", jwt.Authorize(), service.CreateProduct)
	server.AddHandler(constants.HttpGet, constants.ProductServicePath, "/product/:productId", jwt.Authorize(), service.GetProduct)
	server.AddHandler(constants.HttpGet, constants.ProductServicePath, "/products", jwt.Authorize(), service.GetProducts)
	server.AddHandler(constants.HttpPost, constants.ProductServicePath, "/product/:productId", jwt.Authorize(), service.UpdateProduct)
	server.AddHandler(constants.HttpPost, constants.OrderServicePath, "/order", jwt.Authorize(), service.CreateOrder)
	server.AddHandler(constants.HttpGet, constants.OrderServicePath, "/orders", jwt.Authorize(), service.GetOrders)

	// Client Connections
	client.NewCustomerServiceClient()
	client.NewProductServiceClient()
	client.NewOrderManagementServiceClient()

	server.Run(constants.HttpServerPort)
}
