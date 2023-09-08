package paymentmethods

import (
	"payment/service/components"
	"time"
)

type Amex struct{}

var amexClient = "Amex"

func (p *Amex) Execute() bool {
	duration := 2 * time.Second
	isClientAvailable := components.CheckAvailability(amexClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Amex) Rollback() error {
	duration := 5 * time.Second
	time.Sleep(duration)
	return nil
}
