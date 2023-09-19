package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mroth/weightedrand"
	"log"
	"net/http"
	"payment/entity"
	"payment/models"
	"payment/service"
	"time"
)

func InitiatePayment(c *gin.Context) {
	var payment models.Payment

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error reading request": err.Error()})
		return
	}

	log.Printf("Recieved request to initiate payment: %v", payment)

	if err := service.InitiatePayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error initiating the payment": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Payment initiated for OrderID %d, amount %f", payment.OrderID, payment.Amount),
		"status":  "success",
	})
}

func RollbackPayment(c *gin.Context) {
	var payment models.Payment

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.RollbackPayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error rolling back the payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Payment with ID %d has been rolled back.", payment.ID),
		"payment": payment,
	})
}

func GetAvailablePaymentMethods(c *gin.Context) {

	for _, paymentMethod := range entity.PaymentGatewaysConfig {
		go func(paymentMethod weightedrand.Choice) {
			chooser, err := weightedrand.NewChooser(entity.PaymentGatewaysConfig...)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			flakyPaymentMethod := chooser.Pick().(string)
			if paymentMethod.Item == flakyPaymentMethod {
				return
			}
			fmt.Printf(" %v\n", paymentMethod.Item)
		}(paymentMethod)

	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, "successfully derived list of available methods")
	return
}
