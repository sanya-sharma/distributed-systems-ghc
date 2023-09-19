package service

import (
	"catalog/models"
	"catalog/repository"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

var maxRetries = 3

// GetCatalogs fetches the catalog from DB and returns it
func GetCatalog(db *gorm.DB) (response []models.Catalog, err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository

	for retry := 0; retry <= maxRetries; retry++ {
		response, err = catalogRepo.GetCatalog()
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

// GetCatalogByProductID fetches the catalog from DB and returns it
func GetCatalogByProductID(db *gorm.DB, productID int) (response *models.Catalog, err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository

	for retry := 0; retry <= maxRetries; retry++ {
		response, err = catalogRepo.GetCatalogByProductID(productID)
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

// UpdateCatalog updates the catalog DB for given product ID and stockQty
func UpdateCatalog(db *gorm.DB, productID int, quantity int) (err error) {
	catalogRepo := &repository.CatalogRepository{DB: db} // Initialize with actual repository
	// Start a new transaction
	tx := catalogRepo.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			log.Printf("error updating the catalog with productID: %v, quantity: %v with err: %v ", productID, quantity, err)
		} else {
			tx.Commit()
		}
	}()

	var inventory *models.Catalog
	inventory, err = catalogRepo.GetCatalogByProductID(productID)
	if err != nil {
		return err
	}

	// Check inventory stock
	if inventory.StockQty < quantity {
		return errors.New("insufficient stock")
	}

	inventory.StockQty -= quantity

	err = catalogRepo.UpdateCatalog(inventory)
	if err != nil {
		return errors.New("failed to update catalog")
	}

	return
}
