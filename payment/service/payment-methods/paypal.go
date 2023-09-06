package paymentmethods

import (
	"distributed-systems-ghc/payment/service"
	"time"
)

type Paypal struct{}

var paypalClient = "Paypal"

func (p *Paypal) Execute() bool {
	duration := 3 * time.Second
	isClientAvailable := service.CheckAvailability(paypalClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Paypal) Cancel() error {
	duration := 4 * time.Second
	time.Sleep(duration)
	return nil
}
