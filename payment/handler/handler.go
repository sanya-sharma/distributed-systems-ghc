package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mroth/weightedrand"
	"net/http"
	"distributed-systems-ghc/payment/models"
	"distributed-systems-ghc/payment/service"
)

var (
	paymentMethods = []weightedrand.Choice{
		{Item: "Weiss", Weight: 1},
		{Item: "Zakpay", Weight: 5},
		{Item: "Paypal", Weight: 0},
		{Item: "Amex", Weight: 2},
	}
)

func InitiatePayment(c *gin.Context) {
    var payment models.Payment

    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found"})
        return
    }

    if err := service.InitiatePayment(db.(*gorm.DB), payment); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initiating the payment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Payment initiated using %s method for OrderID %d.", payment.PaymentMethod, payment.OrderID),
        "payment": payment,
    })
}


func RollbackPayment(c *gin.Context) {
}

func GetAvailablePaymentMethods(c *gin.Context) {

	for _, paymentMethod := range paymentMethods {
		go func(paymentMethod weightedrand.Choice) {
			chooser, err := weightedrand.NewChooser(paymentMethods...)
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
