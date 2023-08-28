package service

import (
	"distributed-systems-ghc/catalog/models"
	"distributed-systems-ghc/catalog/repository"
	"gorm.io/gorm"
)

func GetCatalogByCategory(db *gorm.DB, category int) (response *models.Inventory, err error) {
	inventoryRepo := &repository.InventoryRepository{DB: db} // Initialize with actual repository
	inventory, err := inventoryRepo.GetInventoryByProductID(category)
	if err != nil {
		return &models.Inventory{}, err
	}

	return inventory, nil
}
func UpdateCatalogByCategory(db *gorm.DB, category int, qty int) (err error) {
	inventoryRepo := &repository.InventoryRepository{DB: db} // Initialize with actual repository

	err = inventoryRepo.UpdateInventory(&models.Inventory{ProductID: category, StockQty: qty})
	if err != nil {
		return err
	}
	return nil
}
