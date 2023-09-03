package paymentmethods

import (
	"time"
)

type Amex struct{}

func (p *Amex) Execute() (error) {
	duration := 2 * time.Second
    time.Sleep(duration)
	return nil
}

func (p *Amex) Cancel() (error) {
	duration := 5 * time.Second
    time.Sleep(duration)
	return nil
}