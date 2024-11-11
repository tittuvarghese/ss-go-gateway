package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tittuvarghese/core/jwt"
	"github.com/tittuvarghese/core/logger"
	"github.com/tittuvarghese/core/validator"
	"github.com/tittuvarghese/customer-service/proto"
	"github.com/tittuvarghese/gateway/client"
	"github.com/tittuvarghese/gateway/models"
	"net/http"
)

var log = logger.NewLogger("gateway-service")

var structValidator = validator.NewStructValidator()

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
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	// Validate input
	if request.Username == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Username and password are required"})
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
		Type:      request.Type,
	}

	resp, err := service.Register(context.Background(), registerReq)
	if err != nil {
		log.Error("Failed to register user", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	// Validate input
	if request.Username == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Username and password are required"})
		return
	}

	// Grpc Request to Customer Service
	service := client.CustomerService

	// Prepare the login request
	loginReq := &proto.LoginRequest{
		Username: request.Username,
		Password: request.Password,
	}

	resp, err := service.Login(context.Background(), loginReq)
	if err != nil {
		log.Error("Failed to authenticate the user", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
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
	claims, _ := jwt.GetClaims(c)

	// Type assert the `Data` field to map[string]interface{}
	request, ok := claims.Data.(map[string]interface{})
	if !ok {
		log.Error("Failed to retrieve the user", fmt.Errorf("invalid map"))
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "failed to retrieve the user"})
	}

	service := client.CustomerService

	// Prepare the get profile request
	userid, ok := request["userid"].(string)
	if !ok {
		log.Error("Failed to retrieve the user", fmt.Errorf("userid is not a string"))
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "failed to retrieve the user"})
		return
	}

	getProfileReq := &proto.GetProfileRequest{
		Userid: userid,
	}

	resp, err := service.GetProfile(context.Background(), getProfileReq)
	if err != nil {
		log.Error("Failed to authenticate the user", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	log.Info("Received response from the customer service %s", resp.Userid)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Successfully retrieved the user information",
		"data":    resp,
	})
}

func GetUser(c *gin.Context) (string, error) {
	claims, _ := jwt.GetClaims(c)

	// Type assert the `Data` field to map[string]interface{}
	request, ok := claims.Data.(map[string]interface{})
	if !ok {
		log.Error("Failed to retrieve the user", fmt.Errorf("invalid map"))
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "failed to retrieve the user"})
	}

	// Prepare the get profile request
	userid, ok := request["userid"].(string)
	if !ok {
		log.Error("Failed to retrieve the user", fmt.Errorf("userid is not a string"))
		return "", fmt.Errorf("unable to get the userid")
	}

	return userid, nil
}
