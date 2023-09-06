package service

import (
	"distributed-systems-ghc/payment/models"
	"errors"
	"fmt"
	"github.com/mroth/weightedrand"
	"strings"
)

var (
	paymentGatewaysConfig = []weightedrand.Choice{
		{Item: "Weiss", Weight: 1},
		{Item: "Zakpay", Weight: 5},
		{Item: "Paypal", Weight: 0},
		{Item: "Amex", Weight: 2},
	}
	paymentGateways = []string{
		"Zakpay",
		"Amex",
		"Weiss",
		"Paypal",
	}
)

func InitiatePayment(payment models.Payment) (err error) {

	for _, paymentGateway := range paymentGateways {
		paymentGatewayClient := getPaymentMethod(paymentGateway)
		if paymentGatewayClient == nil {
			return errors.New("invalid payment method")
		}

		fmt.Printf("Initiating payment using %s gateway for OrderID %d...\n", payment.PaymentMethod, payment.OrderID)

		paymentContext := SetPaymentMethod(paymentGatewayClient)
		err = paymentContext.ExecutePayment()
		if err != nil {
			fmt.Printf("Payment gateway %v\n is %v", paymentGateway, err.Error())
		}
	}

	return err
}

func RollbackPayment(payment models.Payment) (err error) {
	fmt.Printf("Rolling back payment with ID %d...\n", payment.ID)

	paymentMethod := getPaymentMethod(strings.ToLower(payment.PaymentMethod))
	if paymentMethod == nil {
		return errors.New("Invalid payment method")
	}
	paymentContext := SetPaymentMethod(paymentMethod)
	err := paymentContext.RollbackPayment()

	return err
}

func getPaymentMethod(paymentGateway string) PaymentGateways {
	switch strings.ToLower(paymentGateway) {
	case "amex":
		return &paymentGatewaysConfig.Amex{}
	case "paypal":
		return &paymentGatewaysConfig.Paypal{}
	case "weiss":
		return &paymentGatewaysConfig.Weiss{}
	case "zakpay":
		return &paymentGatewaysConfig.Zakpay{}
	default:
		return nil
	}
}
func CheckAvailability(paymentMethod string) bool {

	chooser, err := weightedrand.NewChooser(paymentGatewaysConfig...)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	flakyPaymentMethod := chooser.Pick().(string)
	if paymentMethod == flakyPaymentMethod {
		fmt.Printf("Payment method %v\n is unavailable", paymentMethod)

		return false
	}
	fmt.Printf("Payment method %v\n is available", paymentMethod)
	return true

}
