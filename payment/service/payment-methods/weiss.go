package paymentmethods

import (
	"payment/service/components"
	"time"
)

type Weiss struct{}

var weissClient = "Weiss"

func (p *Weiss) Execute() bool {
	duration := 4 * time.Second
	isClientAvailable := components.CheckAvailability(weissClient)
	time.Sleep(duration)
	return isClientAvailable
}

func (p *Weiss) Rollback() error {
	duration := 3 * time.Second
	time.Sleep(duration)
	return nil
}
