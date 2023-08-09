package handler

import (
	"gorm.io/gorm"
	"net/http"

	"distributed-systems-ghc/service"
	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	var orderData struct {
		CustomerID int
		ProductID  int
		Quantity   int
	}

	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	// Call the service layer to handle the order placement
	order, err := service.PlaceOrder(db.(*gorm.DB), orderData.CustomerID, orderData.ProductID, orderData.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
