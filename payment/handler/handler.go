package handler

import (
	"distributed-systems-ghc/payment/entity"
	"distributed-systems-ghc/payment/models"
	"distributed-systems-ghc/payment/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mroth/weightedrand"
	"net/http"
)

func InitiatePayment(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.InitiatePayment(payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initiating the payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Payment initiated using %s method for OrderID %d.", payment.PaymentMethod, payment.OrderID),
		"payment": payment,
	})
}

func RollbackPayment(c *gin.Context) {
	var payment models.Payment

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
