# Gateway Service for E-Commerce Application

[![GoDoc](https://pkg.go.dev/badge/github.com/tittuvarghese/ss-go-gateway)](https://pkg.go.dev/github.com/tittuvarghese/ss-go-gateway)
[![Build Status](https://travis-ci.org/tittuvarghese/ss-go-gateway.svg?branch=main)](https://travis-ci.org/tittuvarghese/ss-go-gateway)

The **Gateway Service** is a critical component in our e-commerce microservices architecture. It acts as a reverse proxy, accepting HTTP requests and routing them to the appropriate microservices (Customer, Product, and Order) via gRPC. This service exposes several REST endpoints to interact with customer data, product management, and order handling.

## Features

- **HTTP to gRPC Proxy**: Routes HTTP requests to backend microservices using gRPC.
- **Authentication**: All routes that require user authentication are protected via JWT tokens.
- **Microservices Integration**: Handles interactions with Customer, Product, and Order services.

## Endpoints

Here are the key endpoints exposed by the Gateway Service:

### **Status Endpoint**
- **Method**: `GET`
- **Path**: `/status`
- **Description**: Returns the status of the gateway service.

### **Customer Service Endpoints**

1. **Register Customer**
    - **Method**: `POST`
    - **Path**: `/register`
    - **Description**: Registers a new customer.

2. **Login Customer**
    - **Method**: `POST`
    - **Path**: `/login`
    - **Description**: Logs in a customer and returns a JWT token.

3. **Get Customer Profile**
    - **Method**: `GET`
    - **Path**: `/profile`
    - **Description**: Retrieves the customer profile details.
    - **Authorization**: Requires JWT token.

### **Product Service Endpoints**

1. **Create Product**
    - **Method**: `POST`
    - **Path**: `/create`
    - **Description**: Creates a new product.
    - **Authorization**: Requires JWT token.

2. **Get Product by ID**
    - **Method**: `GET`
    - **Path**: `/product/:productId`
    - **Description**: Retrieves product details by product ID.
    - **Authorization**: Requires JWT token.

3. **Get All Products**
    - **Method**: `GET`
    - **Path**: `/products`
    - **Description**: Lists all products.
    - **Authorization**: Requires JWT token.

4. **Update Product**
    - **Method**: `POST`
    - **Path**: `/product/:productId`
    - **Description**: Updates product details by product ID.
    - **Authorization**: Requires JWT token.

### **Order Service Endpoints**

1. **Get Orders**
    - **Method**: `GET`
    - **Path**: `/orders`
    - **Description**: Retrieves all orders for the authenticated user.
    - **Authorization**: Requires JWT token.

2. **Create Order**
    - **Method**: `POST`
    - **Path**: `/order`
    - **Description**: Creates a new order for the authenticated user.
    - **Authorization**: Requires JWT token.

3. **Get Order by ID**
    - **Method**: `GET`
    - **Path**: `/order/:orderId`
    - **Description**: Retrieves order details by order ID.
    - **Authorization**: Requires JWT token.

4. **Update Order**
    - **Method**: `POST`
    - **Path**: `/order/:orderId`
    - **Description**: Updates an existing order by order ID.
    - **Authorization**: Requires JWT token.

## Authentication

This service uses **JWT (JSON Web Tokens)** for user authentication. All endpoints, except for the `/status`, require the user to provide a valid JWT token in the `Authorization` header. The format of the header should be:
Authorization: Bearer <your_jwt_token>


### How to Obtain a JWT Token

1. Register and log in through the `/register` and `/login` endpoints of the Gateway service to obtain a JWT token.
2. Use this token for subsequent requests to authenticate.

## Dependencies

- Go 1.18 or higher
- gRPC
- JWT Authentication
- HTTP server library (e.g., `net/http` or custom server)

## Installation

To install and run the Gateway Service locally:

1. Clone the repository:

   ```bash
   git clone https://github.com/tittuvarghese/ss-go-gateway.git
   cd gateway-service
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the Gateway Service:
    ```bash
   go run cmd/main.go
   ```
   
**Note: The service will start on the default port (e.g., 8080).**

## Example Usage
#### Example: Registering a Customer
```bash
curl -X POST http://localhost:8080/api/v1/register \
   -d '{"username": "john_doe", "password": "secret123"}' \
   -H "Content-Type: application/json"
```
#### Example: Login
```bash
curl -X POST http://localhost:8080/api/v1/login \
   -d '{"username": "john_doe", "password": "secret123"}' \
   -H "Content-Type: application/json"
```
## Architecture
The Gateway Service acts as a reverse proxy and communicates with backend microservices via gRPC. It is designed to provide a unified HTTP interface for clients while abstracting away the complexity of direct gRPC communication with individual microservices.
- **Gateway:** Responsible for accepting all incoming request and forwards to the required services.
- **Customer Service:** Manages user registration, authentication, and profiles. 
- **Product Service:** Manages product listings, creation, and updates. 
- **Order Service:** Manages order creation, updates, and retrieval.