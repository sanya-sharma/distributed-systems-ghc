package entity

import (
	"github.com/mroth/weightedrand"
)

var (
	PaymentGateways = []string{
		"Zakpay",
		"Amex",
		"Weiss",
		"Paypal",
	}
	PaymentGatewaysConfig = []weightedrand.Choice{
		{Item: "Weiss", Weight: 1},
		{Item: "Zakpay", Weight: 9},
		{Item: "Paypal", Weight: 0},
		{Item: "Amex", Weight: 2},
	}
)
