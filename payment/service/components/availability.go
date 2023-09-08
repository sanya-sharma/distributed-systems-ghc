package components

import (
	"fmt"
	"github.com/mroth/weightedrand"
	"payment/entity"
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
	return true

}
