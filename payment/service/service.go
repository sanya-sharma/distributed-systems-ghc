package service

import (
	"gorm.io/gorm"
	"strings"
	"errors"
	"fmt"

	"distributed-systems-ghc/payment/models"
	"distributed-systems-ghc/payment/repository"
	paymentMethods "distributed-systems-ghc/payment/service/payment-methods"
)

func InitiatePayment(db *gorm.DB, payment models.Payment) (err error) {
    fmt.Printf("Initiating payment using %s method for OrderID %d...\n", payment.PaymentMethod, payment.OrderID)

	paymentMethod := getPaymentMethod(strings.ToLower(payment.PaymentMethod))
	if paymentMethod == nil {
		return errors.New("Invalid payment method")
	}
	paymentContext := SetPaymentMethod(paymentMethod)
	err := paymentContext.ExecutePayment()

	return err
}

func getPaymentMethod(paymentMethod string) PaymentMethods{
	switch paymentMethod {
	case "amex":
		return &paymentMethods.Amex{}
	case "paypal":
		return &paymentMethods.Paypal{}
	case "weiss":
		return &paymentMethods.Weiss{}
	case "zakpay":
		return &paymentMethods.Zakpay{}
	default:
		return nil
	}
}
