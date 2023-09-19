package models

// Catalog represents a product that can be ordered.
type Catalog struct {
	ProductID   int `gorm:"column:id"`
	Name        string
	Description string
	Price       float64
	StockQty    int
}
