package service

import (
	"distributed-systems-ghc/payment/models"
	"distributed-systems-ghc/payment/entity"
	paymentMehtods "distributed-systems-ghc/payment/service/payment-methods"
	"errors"
	"fmt"
	"strings"
)



func InitiatePayment(payment models.Payment) (err error) {

	for _, paymentGateway := range entity.PaymentGateways {
		paymentGatewayClient := getPaymentMethod(paymentGateway)
		if paymentGatewayClient == nil {
			return errors.New("invalid payment method")
		}

		fmt.Printf("Initiating payment using %s gateway for OrderID %d...\n", paymentGateway, payment.OrderID)

		paymentContext := &PaymentContext{}
		paymentContext.SetPaymentMethod(paymentGatewayClient)
		completed := paymentContext.ExecutePayment()
		if !completed {
			fmt.Printf("Payment gateway %v is unavailable\n", paymentGateway)
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
	paymentContext := &PaymentContext{}
	paymentContext.SetPaymentMethod(paymentMethod)
	err = paymentContext.RollbackPayment()

	return err
}

func getPaymentMethod(paymentGateway string) PaymentGateways {
	switch strings.ToLower(paymentGateway) {
	case "amex":
		return &paymentMehtods.Amex{}
	case "paypal":
		return &paymentMehtods.Paypal{}
	case "weiss":
		return &paymentMehtods.Weiss{}
	case "zakpay":
		return &paymentMehtods.Zakpay{}
	default:
		return nil
	}
}
