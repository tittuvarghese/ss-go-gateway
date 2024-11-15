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
			ImageUrls:             request.ImageUrls,
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
		log.Error("Failed to retrieve the product", err)
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

func GetProducts(c *gin.Context) {
	// Grpc Request to Product Service
	service := client.ProductService

	getProductsReq := &proto.GetProductsRequest{
		Query: nil,
	}

	resp, err := service.GetProducts(context.Background(), getProductsReq)
	if err != nil {
		log.Error("Failed to retrieve the products", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	log.Info("Received response from the product service %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": resp.Message,
		"data":    resp.Products,
	})
}

func UpdateProduct(c *gin.Context) {
	productId := c.Param("productId")

	log.Info("Product Id: " + productId)

	var request models.Product

	// Bind the JSON data to the user struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
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

	updateProductReq := &proto.UpdateProductRequest{
		ProductId: productId,
		Product: &proto.Product{
			SellerId: userid,
		},
	}

	if request.Name != "" {
		updateProductReq.Product.Name = request.Name
	}
	if request.Quantity > 0 {
		updateProductReq.Product.Quantity = request.Quantity
	}
	if request.Type != "" {
		updateProductReq.Product.Type = request.Type
	}
	if request.Category != "" {
		updateProductReq.Product.Category = request.Category
	}
	if len(request.ImageUrls) > 0 {
		updateProductReq.Product.ImageUrls = request.ImageUrls
	}
	if request.Price > 0 {
		updateProductReq.Product.Price = request.Price
	}
	if request.Width > 0 {
		updateProductReq.Product.Size.Width = request.Width
	}
	if request.Height > 0 {
		updateProductReq.Product.Size.Height = request.Height
	}
	if request.ShippingBasePrice > 0 {
		updateProductReq.Product.ShippingBasePrice = request.ShippingBasePrice
	}
	if request.BaseDeliveryTimelines > 0 {
		updateProductReq.Product.BaseDeliveryTimelines = request.BaseDeliveryTimelines
	}

	resp, err := service.UpdateProduct(context.Background(), updateProductReq)
	if err != nil {
		log.Error("Failed to retrieve the product", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	log.Info("Received response from the product service %s", resp.Message)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": resp.Message,
	})
}
