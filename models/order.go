package models

import (
	"time"
)

// Order represents an individual order.
type Order struct {
	ID         int         `db:"id"`
	Customer   Customer    `db:"customer"`
	OrderDate  time.Time   `db:"order_date"`
	Status     string      `db:"status"`
	Items      []OrderItem `db:"items"`
	Created_at time.Time   `db:"created_at"`
	Updated_at time.Time   `db:"updated_at"`
}
