package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tittuvarghese/gateway/client"
	"github.com/tittuvarghese/gateway/models"
	"github.com/tittuvarghese/product-service/proto"
	"net/http"
)

func CreateProduct(c *gin.Context) {

	var request models.Product

	// Bind the JSON data to the user struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	// Validate input
	err := structValidator.Validate(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Invalid request format"})
		return
	}

	// Extract userid
	userid, err := GetUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": "Unable to extract the user details"})
		return
	}

	// Grpc Request to Product Service
	service := client.ProductService
	// Prepare the registration request
	registerReq := &proto.CreateProductRequest{
		Product: &proto.Product{
			Name:                  request.Name,
			Quantity:              request.Quantity,
			Type:                  request.Type,
			Category:              request.Category,
			ImageUrls:             request.ImageURLs,
			Price:                 request.Price,
			Size:                  &proto.Product_Size{Width: request.Width, Height: request.Height},
			Weight:                request.Weight,
			ShippingBasePrice:     request.ShippingBasePrice,
			BaseDeliveryTimelines: request.BaseDeliveryTimelines,
			SellerId:              userid,
		},
	}
	resp, err := service.CreateProduct(context.Background(), registerReq)
	if err != nil {
		log.Error("Failed to create the product", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	log.Info("Received response from the product service %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": resp.Message,
	})
}

func GetProduct(c *gin.Context) {
	productId := c.Param("productId")

	// Grpc Request to Product Service
	service := client.ProductService

	getProductReq := &proto.GetProductRequest{
		ProductId: productId,
	}

	resp, err := service.GetProduct(context.Background(), getProductReq)
	if err != nil {
		log.Error("Failed to create the product", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	log.Info("Received response from the product service %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": resp.Message,
		"data":    resp.Product,
	})
}

func GetProducts(c *gin.Context) {}
