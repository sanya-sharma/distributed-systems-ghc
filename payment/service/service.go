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
	var completed bool
	for _, paymentGateway := range entity.PaymentGateways {
		paymentGatewayClient := getPaymentMethod(paymentGateway)
		if paymentGatewayClient == nil {
			return errors.New("invalid payment gateway")
		}

		log.Printf("Initiating payment using %s gateway for OrderID %d \n", paymentGateway, payment.OrderID)

		paymentContext := &PaymentContext{}
		paymentContext.SetPaymentMethod(paymentGatewayClient)

		for retry := 0; retry <= maxRetries; retry++ {
			if retry != 0  {
				// Log the retry and sleep before the next attempt
				log.Printf("Payment gateway %v is unavailable. Retrying payment, attempt %d", paymentGateway, retry+1)
				time.Sleep(time.Second * time.Duration(retry))
			}
			completed = paymentContext.ExecutePayment()
			if completed {
				break
			}
		}

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
