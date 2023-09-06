package repository

import (
	"gorm.io/gorm"
	"order/models"
)

type CustomerRepository struct {
	DB *gorm.DB
}

// CreateCustomer inserts a new customer record into the database.
func (repo *CustomerRepository) CreateCustomer(customer *models.Customer) error {
	result := repo.DB.Create(customer)
	return result.Error
}

// DeleteCustomer deletes a customer record from the database.
func (repo *CustomerRepository) DeleteCustomer(customer *models.Customer) error {
	result := repo.DB.Delete(customer)
	return result.Error
}

// UpdateCustomer updates a customer record in the database.
func (repo *CustomerRepository) UpdateCustomer(customer *models.Customer) error {
	result := repo.DB.Save(customer)
	return result.Error
}

// GetCustomerByID retrieves a customer record by ID from the database.
func (repo *CustomerRepository) GetCustomerByID(id int) (*models.Customer, error) {
	var customer models.Customer
	result := repo.DB.First(&customer, id)
	return &customer, result.Error
}
