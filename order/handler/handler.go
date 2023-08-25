package handler

import (
	"context"
	"distributed-systems-ghc/order/models"
	"distributed-systems-ghc/order/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func PlaceOrder(c *gin.Context) {
	var orderData models.OrderData
	if err := c.ShouldBindJSON(&orderData); err != nil {
		log.Printf("Error while parsing order data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
