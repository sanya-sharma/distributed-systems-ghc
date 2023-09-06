package paymentmethods

import (
	"distributed-systems-ghc/payment/service"
	"time"
)

type Zakpay struct{}

var zakpayClient = "Zakpay"

func (p *Zakpay) Execute() bool {
	duration := 5 * time.Second
	isClientAvailable := service.CheckAvailability(zakpayClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Zakpay) Cancel() error {
	duration := 2 * time.Second
	time.Sleep(duration)
	return nil
}
