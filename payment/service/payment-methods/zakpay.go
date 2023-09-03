package paymentmethods

import (
	"time"
)

type Zakpay struct{}

func (p *Zakpay) Execute() (error) {
	duration := 5 * time.Second
    time.Sleep(duration)
}

func (p *Zakpay) Cancel() (error) {
	duration := 2 * time.Second
    time.Sleep(duration)
}