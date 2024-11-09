package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tittuvarghese/core/logger"
	"github.com/tittuvarghese/customer-service/proto"
	"github.com/tittuvarghese/gateway/client"
	"github.com/tittuvarghese/gateway/models"
	"net/http"
)

var log = logger.NewLogger("gateway-service")

// Status service that returns server status
func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Server is running smoothly",
	})
}

func Register(c *gin.Context) {
	// Initialize user struct
	var request models.RegisterRequest

	// Bind the JSON data to the user struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if request.Username == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Grpc Request to Customer Service

	service := client.CustomerService
	// Prepare the registration request
	registerReq := &proto.RegisterRequest{
		Username:  request.Username,
		Password:  request.Password,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
	}

	resp, err := service.Register(context.Background(), registerReq)
	if err != nil {
		log.Error("Failed to register user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Info("Received response from the customer service %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": resp.Message,
	})
}

func Login(c *gin.Context) {

	var request models.LoginRequest

	// Bind the JSON data to the user struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if request.Username == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Grpc Request to Customer Service

	service := client.CustomerService
	// Prepare the registration request
	loginReq := &proto.LoginRequest{
		Username: request.Username,
		Password: request.Password,
	}

	resp, err := service.Login(context.Background(), loginReq)
	if err != nil {
		log.Error("Failed to authenticate the user", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Info("Received response from the customer service %s", resp.Token)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully authenticated the user",
		"token":   resp.Token,
	})

}

func GetProfile(c *gin.Context) {
	claims, _ := c.Get("claims")

	// Do something with the claims (e.g., display the username)
	c.JSON(200, gin.H{
		"claims": claims,
	})
}
