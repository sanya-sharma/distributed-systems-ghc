package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mroth/weightedrand"
	"net/http"
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
