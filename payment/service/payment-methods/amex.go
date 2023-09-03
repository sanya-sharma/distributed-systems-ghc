package paymentmethods

import (
	"time"
)

type Amex struct{}

func (p *Amex) Execute() (error) {
	duration := 2 * time.Second
    time.Sleep(duration)
}