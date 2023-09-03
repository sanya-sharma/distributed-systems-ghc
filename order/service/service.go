package service

import (
	"distributed-systems-ghc/order/models"
	"distributed-systems-ghc/order/repository"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

var maxRetries = 3

func PlaceOrder(db *gorm.DB, customerID, productID, quantity int) (order models.Order, err error) {
	customerRepo := &repository.CustomerRepository{DB: db}   // Initialize with actual repository
	inventoryRepo := &repository.InventoryRepository{DB: db} // Initialize with actual repository
	// maximum number of downstream retries

	var customer *models.Customer
	for retry := 0; retry <= maxRetries; retry++ {
		customer, err = customerRepo.GetCustomerByID(customerID)
		if err == nil {
			break
		}

		// Log the retry and sleep before the next attempt
		log.Printf("Retrying GetCustomerByID, attempt %d", retry+1)
		time.Sleep(time.Second * time.Duration(retry+1))
	}
	if err != nil {
		return order, err
	}

	newOrder := models.Order{
		Customer: *customer,
	}

	err = updateCatalog(inventoryRepo, productID, quantity)
	if err != nil {
		return order, err
	}

	// Place the order and update inventory stock
	// Implement your order placement and inventory update logic here

	return newOrder, nil
}

func updateCatalog(inventoryRepo *repository.InventoryRepository, productID int, quantity int) (err error) {

	// Start a new transaction
	tx := inventoryRepo.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			err = errors.New("failed to update catalog")
		} else {
			tx.Commit()
		}
	}()

	var inventory *models.Inventory
	inventory, err = inventoryRepo.GetInventoryByProductID(productID)
	if err != nil {
		return err
	}

	// Check inventory stock
	if inventory.StockQty < quantity {
		return errors.New("insufficient stock")
	}

	inventory.StockQty -= quantity
	if err := inventoryRepo.UpdateInventory(inventory); err != nil {
		tx.Rollback()
		return errors.New("failed to update catalog")
	}
	return
}
