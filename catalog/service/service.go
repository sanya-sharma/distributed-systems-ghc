package service

import (
	"catalog/models"
	"catalog/repository"
	"gorm.io/gorm"
)

func GetCatalogByCategory(db *gorm.DB, category int) (response *models.Catalog, err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository
	catalog, err := catalogRepo.GetCatalogByProductID(category)
	if err != nil {
		return &models.Catalog{}, err
	}

	return catalog, nil
}
func UpdateCatalogByCategory(db *gorm.DB, category int, qty int) (err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository

	err = catalogRepo.UpdateCatalog(&models.Catalog{ID: category, StockQty: qty})
	if err != nil {
		return err
	}
	return nil
}
