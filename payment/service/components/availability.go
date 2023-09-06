package components

import(
	"fmt"
	"github.com/mroth/weightedrand"
	"distributed-systems-ghc/payment/entity"
)

func CheckAvailability(paymentMethod string) bool {

	chooser, err := weightedrand.NewChooser(entity.PaymentGatewaysConfig...)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	flakyPaymentMethod := chooser.Pick().(string)
	if paymentMethod == flakyPaymentMethod {
		fmt.Printf("Payment method %v is unavailable\n", paymentMethod)

		return false
	}
	fmt.Printf("Payment method %v is available\n", paymentMethod)
	return true

}