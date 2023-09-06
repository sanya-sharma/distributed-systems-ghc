package paymentmethods

import (
	"distributed-systems-ghc/payment/service"
	"time"
)

type Weiss struct{}

var weissClient = "Weiss"

func (p *Weiss) Execute() bool {
	duration := 4 * time.Second
	isClientAvailable := service.CheckAvailability(weissClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Weiss) Cancel() error {
	duration := 3 * time.Second
	time.Sleep(duration)
	return nil
}
