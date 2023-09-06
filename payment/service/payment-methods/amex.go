package paymentmethods

import (
	"distributed-systems-ghc/payment/service"
	"time"
)

type Amex struct{}

var amexClient = "Amex"

func (p *Amex) Execute() bool {
	duration := 2 * time.Second
	isClientAvailable := service.CheckAvailability(amexClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Amex) Cancel() error {
	duration := 5 * time.Second
	time.Sleep(duration)
	return nil
}
