package paymentmethods

import (
	"payment/service/components"
	"time"
)

type Zakpay struct{}

var zakpayClient = "Zakpay"

func (p *Zakpay) Execute() bool {
	duration := 2 * time.Second
	isClientAvailable := components.CheckAvailability(zakpayClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Zakpay) Rollback() error {
	duration := 2 * time.Second
	time.Sleep(duration)
	return nil
}
