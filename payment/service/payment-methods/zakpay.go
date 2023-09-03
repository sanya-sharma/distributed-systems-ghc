package paymentmethods

import (
	"time"
)

type Zakpay struct{}

func (p *Zakpay) Execute() (error) {
	duration := 5 * time.Second
    time.Sleep(duration)
}