package service

import (
	"catalog/models"
	"catalog/repository"
	"gorm.io/gorm"
)

func GetCatalog(db *gorm.DB) (response []models.Catalog, err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository
	catalog, err := catalogRepo.GetCatalog()
	if err != nil {
		return nil, err
	}

	return catalog, nil
}
func UpdateCatalogByCategory(db *gorm.DB, category int, qty int) (err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository

	err = catalogRepo.UpdateCatalog(&models.Catalog{ProductID: category, StockQty: qty})
	if err != nil {
		return err
	}
	return nil
}
