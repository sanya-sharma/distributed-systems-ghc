package service

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"distributed-systems-ghc/models"
	"distributed-systems-ghc/repository"
)

func PlaceOrder(db *gorm.DB, customerID, productID, quantity int) (order models.Order, err error) {
	customerRepo := &repository.CustomerRepository{DB: db}   // Initialize with actual repository
	inventoryRepo := &repository.InventoryRepository{DB: db} // Initialize with actual repository
	maxRetries := 3 // maximum number of downstream retries

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

	var inventory *models.Inventory
	for retry := 0; retry <= maxRetries; retry++ {
		inventory, err = inventoryRepo.GetInventoryByProductID(productID)
		if err == nil {
			break
		}

		// Log the retry and sleep before the next attempt
		log.Printf("Retrying GetInventoryByProductID, attempt %d", retry+1)
		time.Sleep(time.Second * time.Duration(retry+1))
	}
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
