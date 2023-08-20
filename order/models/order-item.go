package models

// OrderItem contains information about each individual item within an order.
type OrderItem struct {
	ID       int     `db:"id"`
	OrderID  int     `db:"order_id"`
	Product  Product `db:"product"`
	Quantity int     `db:"quantity"`
	Price    float64 `db:"price"`
}
