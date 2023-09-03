package paymentmethods

import (
	"time"
)

type Weiss struct{}

func (p *Weiss) Execute() (error) {
	duration := 4 * time.Second
    time.Sleep(duration)
}

func (p *Weiss) Cancel() (error) {
	duration := 3 * time.Second
    time.Sleep(duration)
}