package models

import (
	"time"
)

// Order represents an individual order.
type Order struct {
	ID         int  `gorm:"column:id"`
	CustomerID uint `gorm:"column:customer_id"`
	OrderDate  time.Time
	Status     string
	Created_at time.Time
	Updated_at time.Time
}

type OrderData struct {
	CustomerID int `json:"customer_id"`
	ProductID  int `json:"product_id"`
	Quantity   int `json:"quantity"`
}
