package models

// CatalogRequest represents the
type CatalogRequest struct {
	CategoryID int32
}

// CatalogResponse represents the
type CatalogResponse struct {
	CategoryID int
	Quantity   int
}
