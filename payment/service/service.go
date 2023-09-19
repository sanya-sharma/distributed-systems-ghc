package service

import (
	"errors"
	"fmt"
	"log"
	"payment/entity"
	"payment/models"
	paymentMehtods "payment/service/payment-methods"
	"strings"
	"time"
)

var maxRetries = 3

func InitiatePayment(payment models.Payment) (err error) {
	for _, paymentGateway := range entity.PaymentGateways {
		paymentGatewayClient := getPaymentMethod(paymentGateway)
		if paymentGatewayClient == nil {
			return errors.New("invalid payment gateway")
		}

		log.Printf("Initiating payment using %s gateway for OrderID %d...\n", paymentGateway, payment.OrderID)

		paymentContext := &PaymentContext{}
		paymentContext.SetPaymentMethod(paymentGatewayClient)
		for retry := 0; retry <= maxRetries; retry++ {
			completed := paymentContext.ExecutePayment()
			if !completed {
				break
			}

			// Log the retry and sleep before the next attempt
			log.Printf("Payment gateway %v is unavailable\n Retrying payment, attempt %d", paymentGateway, retry+1)
			time.Sleep(time.Second * time.Duration(retry+1))
		}
		completed := paymentContext.ExecutePayment()
		if !completed {
			log.Printf("Could not attempt payment via %v gateway\n", paymentGateway)
		} else {
			log.Printf("Payment successful\n")
			return nil
		}
	}

	log.Printf("Your payment could not be processed")
	return errors.New("payment was unsuccessful, please try again after some time")
}

func RollbackPayment(payment models.Payment) (err error) {
	fmt.Printf("Rolling back payment with ID %d...\n", payment.ID)
	//todo no payment method to pass
	//paymentMethod := getPaymentMethod(strings.ToLower(payment.PaymentMethod))
	//if paymentMethod == nil {
	//	return errors.New("Invalid payment method")
	//}
	//paymentContext := &PaymentContext{}
	//paymentContext.SetPaymentMethod(paymentMethod)
	//err = paymentContext.RollbackPayment()

	return nil
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
