package repository

import (
	"catalog/models"
	"gorm.io/gorm"
)

type CatalogRepository struct {
	DB *gorm.DB
}

// CreateCatalog inserts a new catalog record into the database.
func (repo *CatalogRepository) CreateCatalog(catalog *models.Catalog) error {
	result := repo.DB.Create(catalog)
	return result.Error
}

// DeleteCatalog deletes an catalog record from the database.
func (repo *CatalogRepository) DeleteCatalog(catalog *models.Catalog) error {
	result := repo.DB.Delete(catalog)
	return result.Error
}

// UpdateCatalog updates an catalog record in the database.
func (repo *CatalogRepository) UpdateCatalog(catalog *models.Catalog) error {
	result := repo.DB.Save(catalog)
	return result.Error
}

// GetCatalogByProductID retrieves an catalog record by product ID from the database.
func (repo *CatalogRepository) GetCatalogByProductID(productID int) (*models.Catalog, error) {
	var catalog models.Catalog
	result := repo.DB.First(&catalog, "product_id = ?", productID)
	return &catalog, result.Error
}
