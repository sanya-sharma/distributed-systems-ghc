package models

// Catalog represents a product that can be ordered.
type Catalog struct {
	ID          int `gorm:"column:id"`
	Name        string
	Description string
	Price       float64
	StockQty    int
}

// CatalogRequest represents the
type CatalogRequest struct {
	CategoryID int32
}

// CatalogResponse represents the
type CatalogResponse struct {
	CategoryID int
	Quantity   int
}
