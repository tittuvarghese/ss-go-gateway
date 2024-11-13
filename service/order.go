package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tittuvarghese/gateway/client"
	"github.com/tittuvarghese/gateway/models"
	"github.com/tittuvarghese/order-management-service/proto"
	"net/http"
)

func CreateOrder(c *gin.Context) {

	var request models.Order

	// Bind the JSON data to the user struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	// Validate input
	err := structValidator.Validate(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Invalid request format, " + err.Error()})
		return
	}

	// Extract userid
	userid, err := GetUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Unable to extract the user details"})
		return
	}

	// Grpc Request to Product Service
	service := client.OrderManagementService
	// Prepare the registration request
	orderReq := &proto.CreateOrderRequest{
		CustomerId: userid,
		Phone:      request.Phone,
		Address: &proto.Address{
			AddressLine1: request.Address.AddressLine1,
			AddressLine2: request.Address.AddressLine2,
			City:         request.Address.City,
			State:        request.Address.State,
			Country:      request.Address.Country,
			Zip:          request.Address.Zip,
		},
	}

	for _, item := range request.Items {
		orderedItem := &proto.OrderItem{
			Quantity:  item.Quantity,
			ProductId: item.ProductID,
			Price:     item.Price,
		}

		orderReq.Items = append(orderReq.Items, orderedItem)
	}

	resp, err := service.CreateOrder(context.Background(), orderReq)
	if err != nil {
		log.Error("Failed to create the order", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	log.Info("Received response from the order management service %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": resp.Message,
	})
}

func GetOrders(c *gin.Context) {

	// Extract userid
	userid, err := GetUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Unable to extract the user details"})
		return
	}

	// Grpc Request to Product Service
	service := client.OrderManagementService
	// Prepare the registration request
	ordersReq := &proto.GetOrdersRequest{
		CustomerId: userid,
	}

	resp, err := service.GetOrders(context.Background(), ordersReq)
	if err != nil {
		log.Error("Failed to get the orders", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	log.Info("Received response from the order management service %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": resp.Message,
		"data":    resp.Orders,
	})

}
