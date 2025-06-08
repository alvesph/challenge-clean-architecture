package entity

import "gorm.io/gorm"

type Order struct {
	CustomerID string  `json:"customer_id"`
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"` // e.g., "pending", "shipped", "delivered"
	gorm.Model
}
