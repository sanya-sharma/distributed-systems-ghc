package models

// Product represents a product that can be ordered.
type Product struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	StockQty    int     `db:"stock_qty"`
}
