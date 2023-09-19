package service

import (
	"catalog/models"
	"catalog/repository"
	"gorm.io/gorm"
)

var maxRetries = 3

// GetCatalogs fetches the catalog from DB and returns it
func GetCatalog(db *gorm.DB) (response []models.Catalog, err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository

	for retry := 0; retry <= maxRetries; retry++ {
		response, err := catalogRepo.GetCatalog()
		if err == nil {
			break
		}

		// Log the retry and sleep before the next attempt
		log.Printf("Retrying GetCatalog from DB, attempt %d", retry+1)
		time.Sleep(time.Second * time.Duration(retry+1))
	}
	if err != nil {
		return response, err
	}

	return response, nil
}

// UpdateCatalogByCategory updates the catalog DB for given product ID and stockQty
func UpdateCatalogByCategory(db *gorm.DB, category int, qty int) (err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository

	for retry := 0; retry <= maxRetries; retry++ {
		err := catalogRepo.UpdateCatalog(&models.Catalog{ProductID: category, StockQty: qty})
		if err == nil {
			break
		}

		// Log the retry and sleep before the next attempt
		log.Printf("Retrying UpdateCatalog from DB, attempt %d", retry+1)
		time.Sleep(time.Second * time.Duration(retry+1))
	}
	if err != nil {
		return err
	}

	return nil
}
