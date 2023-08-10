package models

// Payment tracks the payment details for each order.
type Payment struct {
	ID            int     `db:"id"`
	OrderID       int     `db:"order_id"`
	PaymentMethod string  `db:"payment_method"`
	Amount        float64 `db:"amount"`
	Status        string  `db:"status"`
}
