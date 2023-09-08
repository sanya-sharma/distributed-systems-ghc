package repository

import (
	"gorm.io/gorm"
	"order/models"
)

type OrderRepository struct {
	DB *gorm.DB
}

// CreateOrder inserts a new order record into the database.
func (repo *OrderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
	result := repo.DB.Create(order)
	return order, result.Error
}

// DeleteOrder deletes a order record from the database.
func (repo *OrderRepository) DeleteOrder(order *models.Order) error {
	result := repo.DB.Delete(order)
	return result.Error
}

// UpdateOrder updates a order record in the database.
func (repo *OrderRepository) UpdateOrder(order *models.Order) error {
	result := repo.DB.Save(order)
	return result.Error
}

// GetOrderByID retrieves a order record by ID from the database.
func (repo *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
	var order models.Order
	result := repo.DB.First(&order, id)
	return &order, result.Error
}
