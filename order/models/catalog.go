package models

// Catalog represents the stock of products available.
type Catalog struct {
	ProductID   int `gorm:"column:id"`
	Name        string
	Description string
	Price       float64
	StockQty    int
}
