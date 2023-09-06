package paymentmethods

import (
	"distributed-systems-ghc/payment/service/components"
	"time"
)

type Paypal struct{}

var paypalClient = "Paypal"

func (p *Paypal) Execute() bool {
	duration := 3 * time.Second
	isClientAvailable := components.CheckAvailability(paypalClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Paypal) Rollback() error {
	duration := 4 * time.Second
	time.Sleep(duration)
	return nil
}
