package models

// Inventory represents the stock of products available.
type Inventory struct {
	ProductID int `db:"product_id"`
	StockQty  int `db:"stock_qty"`
}
