package service

import (
	"errors"
	"log"
	"payment/entity"
	"payment/models"
	paymentMehtods "payment/service/payment-methods"
	"strings"
	"time"
)

var maxRetries = 3

type CircuitBreaker struct {
	mu   sync.Mutex
	open bool
}

var circuitBreakerMap = map[string]*CircuitBreaker{}

func InitiatePayment(payment models.Payment) (err error) {
	for _, paymentGateway := range entity.PaymentGateways {
		paymentGatewayClient := getPaymentMethod(paymentGateway)
		if paymentGatewayClient == nil {
			return errors.New("invalid payment gateway")
		}

		log.Printf("Initiating payment using %s gateway for OrderID %d \n", paymentGateway, payment.OrderID)

		paymentContext := &PaymentContext{}
		paymentContext.SetPaymentMethod(paymentGatewayClient)

		circuit, ok := circuitBreakerMap[paymentGateway]
		if !ok {
			// Create a new CircuitBreaker for the payment gateway
			circuit = &CircuitBreaker{}
			circuitBreakerMap[paymentGateway] = circuit
		}

		var completed bool

		for retry := 0; retry <= maxRetries; retry++ {

			if retry != 0 {
				// Log the retry and sleep before the next attempt
				log.Printf("Payment gateway %v is unavailable. Retrying payment, attempt %d", paymentGateway, retry)
				time.Sleep(time.Second * time.Duration(retry))
			}

			completed = circuit.ExecuteTransaction(func() bool {
				return paymentContext.ExecutePayment()
			}, retry+1, paymentGateway)
			if completed {
				break
			}
			// Log the retry and sleep before the next attempt
			log.Printf("Payment gateway %v is unavailable\n Retrying payment, attempt %d", paymentGateway, retry+1)
			time.Sleep(time.Second * time.Duration(retry+1))
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
	log.Printf("Rolling back payment with ID %d", payment.ID)

	return nil
}

func getPaymentMethod(paymentGateway string) PaymentGateways {
	switch paymentGateway {
	case entity.PaymentGatewayAmex:
		return &paymentMehtods.Amex{}
	case entity.PaymentGatewayPaypal:
		return &paymentMehtods.Paypal{}
	case entity.PaymentGatewayWeiss:
		return &paymentMehtods.Weiss{}
	case entity.PaymentGatewayZakpay:
		return &paymentMehtods.Zakpay{}
	default:
		return nil
	}
}

func (cb *CircuitBreaker) ExecuteTransaction(operation func() bool, consecutiveFails int, paymentGateway string) bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.open {
		log.Printf("%v is down, not retrying", paymentGateway)
		return false
	}

	completed := operation()

	if !completed {
		if consecutiveFails >= 3 {
			cb.open = true
			log.Printf("Circuit is open for %v", paymentGateway)
			go cb.ResetAfterDelay(paymentGateway)
		}
	}

	return completed
}

func (cb *CircuitBreaker) ResetAfterDelay(paymentGateway string) {
	// Schedule a reset of the circuit after a delay
	time.Sleep(30 * time.Second)
	cb.mu.Lock()
	cb.open = false
	cb.mu.Unlock()
	log.Printf("Circuit is reset for %v", paymentGateway)
}
