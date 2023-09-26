package components

import (
	"github.com/jmcvetta/randutil"
	"log"
	"payment/entity"
)

func CheckAvailability(paymentMethod string) bool {

	chooser, err := randutil.WeightedChoice(entity.PaymentGatewaysConfig)
	if err != nil {
		log.Println("Error: ", err)
		return false
	}
	if paymentMethod == chooser.Item {
		return false
	}
	return true

}
