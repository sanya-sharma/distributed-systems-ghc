package components

import (
	"fmt"
	"github.com/jmcvetta/randutil"
	"payment/entity"
)

func CheckAvailability(paymentMethod string) bool {

	chooser, err := randutil.WeightedChoice(entity.PaymentGatewaysConfig)
	//NewChooser(entity.PaymentGatewaysConfig...)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	//flakyPaymentMethod := chooser.
	if paymentMethod == chooser.Item {
		return false
	}
	return true

}
