package repository

import (
	"distributed-systems-ghc/models"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	DB *gorm.DB
}

// CreateInventory inserts a new inventory record into the database.
func (repo *InventoryRepository) CreateInventory(inventory *models.Inventory) error {
	result := repo.DB.Create(inventory)
	return result.Error
}

// DeleteInventory deletes an inventory record from the database.
func (repo *InventoryRepository) DeleteInventory(inventory *models.Inventory) error {
	result := repo.DB.Delete(inventory)
	return result.Error
}

// UpdateInventory updates an inventory record in the database.
func (repo *InventoryRepository) UpdateInventory(inventory *models.Inventory) error {
	result := repo.DB.Save(inventory)
	return result.Error
}

// GetInventoryByProductID retrieves an inventory record by product ID from the database.
func (repo *InventoryRepository) GetInventoryByProductID(productID int) (*models.Inventory, error) {
	var inventory models.Inventory
	result := repo.DB.First(&inventory, "product_id = ?", productID)
	return &inventory, result.Error
}
