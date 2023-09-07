package components

import (
	"distributed-systems-ghc/payment/entity"
	"fmt"
	"github.com/mroth/weightedrand"
)

func CheckAvailability(paymentMethod string) bool {

	chooser, err := weightedrand.NewChooser(entity.PaymentGatewaysConfig...)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	flakyPaymentMethod := chooser.Pick().(string)
	if paymentMethod == flakyPaymentMethod {
		return false
	}
	fmt.Printf("Payment method %v is available\n", paymentMethod)
	return true

}
