package models

import (
	"time"
)

// Order represents an individual order.
type Order struct {
	ID         int       `db:"id"`
	Customer   Customer  `db:"customer"`
	OrderDate  time.Time `db:"order_date"`
	Status     string    `db:"status"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

type OrderData struct {
	CustomerID int
	ProductID  int
	Quantity   int
}
