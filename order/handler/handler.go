package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"order/models"
	"order/service"
	"context"
	"time"
)

func PlaceOrder(c *gin.Context) {

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	var orderData models.OrderData
	if err := c.ShouldBindJSON(&orderData); err != nil {
		log.Printf("Error while parsing order data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	// Call the service layer to handle the order placement
	order, err := service.PlaceOrder(db.(*gorm.DB), orderData.CustomerID, orderData.ProductID, orderData.Quantity)
	if err != nil {
		log.Printf("Error while placing order: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
