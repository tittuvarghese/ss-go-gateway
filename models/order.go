package models

// OrderItem model
type OrderItem struct {
	ProductID string  `json:"product_id" validate:"required"`
	Quantity  int32   `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}

type Order struct {
	Items   []OrderItem `json:"items" validate:"required"`
	Phone   string      `json:"phone" validate:"required"`
	Address Address     `json:"address" validate:"required"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1" validate:"required"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required"`
	Zip          string `json:"zip" validate:"required"`
	Country      string `json:"country"`
}

type OrderStatusUpdate struct {
	Status string `json:"status" validate:"required"`
}
