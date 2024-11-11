package models

type Product struct {
	Name                  string   `gorm:"type:varchar(255);not null" json:"name" validate:"required,min=3,max=100"`
	Quantity              int32    `gorm:"not null" json:"quantity" validate:"required,gt=0"`
	Type                  string   `gorm:"type:varchar(20);not null" json:"type" validate:"required"`
	Category              string   `gorm:"type:varchar(100);not null" json:"category" validate:"required"`
	ImageURLs             []string `gorm:"type:json" json:"image_urls" validate:"required"`
	Price                 float64  `gorm:"type:decimal(10,2);not null" json:"price" validate:"required,gt=0"`
	Width                 float64  `gorm:"type:decimal(5,2)" json:"width" validate:"required,gt=0"`
	Height                float64  `gorm:"type:decimal(5,2)" json:"height" validate:"required,gt=0"`
	Weight                float64  `gorm:"type:decimal(5,2)" json:"weight" validate:"required,gt=0"`
	ShippingBasePrice     float64  `gorm:"type:decimal(10,2);not null" json:"shipping_base_price" validate:"required,gt=0"`
	BaseDeliveryTimelines int32    `gorm:"not null" json:"base_delivery_timelines" validate:"required,gt=0"`
}
