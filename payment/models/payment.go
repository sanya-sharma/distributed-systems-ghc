package models

// Payment tracks the payment details for each order.
type Payment struct {
	ID            int     `json:"id"`
	OrderID       int     `json:"order_id"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
}
