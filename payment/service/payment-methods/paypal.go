package paymentmethods

import (
	"time"
)

type Paypal struct{}

func (p *Paypal) Execute() (error) {
	duration := 3 * time.Second
    time.Sleep(duration)
}