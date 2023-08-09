package service

import (
	"distributed-systems-ghc/models"
	"distributed-systems-ghc/repository"
	"errors"
	"gorm.io/gorm"
)

func PlaceOrder(db *gorm.DB, customerID, productID, quantity int) (order models.Order, err error) {
	customerRepo := &repository.CustomerRepository{DB: db}   // Initialize with actual repository
	inventoryRepo := &repository.InventoryRepository{DB: db} // Initialize with actual repository

	customer, err := customerRepo.GetCustomerByID(customerID)
	if err != nil {
		return order, err
	}

	inventory, err := inventoryRepo.GetInventoryByProductID(productID)
	if err != nil {
		return order, err
	}

	// Check inventory stock
	if inventory.StockQty < quantity {
		return order, errors.New("insufficient stock")
	}

	newOrder := models.Order{
		Customer: *customer,
	}

	// Place the order and update inventory stock
	// Implement your order placement and inventory update logic here

	return newOrder, nil
}
