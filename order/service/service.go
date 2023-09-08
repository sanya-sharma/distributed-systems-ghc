package service

import (
	"errors"
	"log"
	"order/gateways"
	"order/models"
	"order/repository"
	"time"

	"gorm.io/gorm"
)

var maxRetries = 3

func PlaceOrder(db *gorm.DB, customerID, productID, quantity int) (order models.Order, err error) {
	customerRepo := &repository.CustomerRepository{DB: db} // Initialize with actual repository
	catalogRepo := &repository.CatalogRepository{DB: db}   // Initialize with actual repository
	orderRepo := &repository.OrderRepository{DB: db}

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

	err = updateCatalog(catalogRepo, productID, quantity)
	if err != nil {
		return order, err
	}

	newOrder := &models.Order{
		CustomerID: uint(customer.ID),
		OrderDate:  time.Now(),
		Status:     "Placed",
		Created_at: time.Now(),
	}

	// Place the order and update inventory stock
	err = gateways.InitiatePayment(*newOrder)
	if err != nil {
		newOrder.Status = "Failed"
		log.Printf("Error in payment gateway: %v", err)
		return *newOrder, err
	}

	newOrder, err = orderRepo.CreateOrder(newOrder)
	if err != nil {
		return *newOrder, err
	}
	return *newOrder, nil

}

func updateCatalog(catalogRepo *repository.CatalogRepository, productID int, quantity int) (err error) {

	// Start a new transaction
	tx := catalogRepo.DB.Begin()
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
	if err := catalogRepo.UpdateCatalog(inventory); err != nil {
		tx.Rollback()
		return errors.New("failed to update catalog")
	}
	return
}
