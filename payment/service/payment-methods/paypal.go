package paymentmethods

import (
	"time"
)

type Paypal struct{}

func (p *Paypal) Execute() (error) {
	duration := 3 * time.Second
    time.Sleep(duration)
}

func (p *Paypal) Cancel() (error) {
	duration := 4 * time.Second
    time.Sleep(duration)
}