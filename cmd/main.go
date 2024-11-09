package main

import (
	"github.com/tittuvarghese/core/config"
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

	// Client Connections
	client.NewCustomerServiceClient()

	server.Run(constants.HttpServerPort)
}
