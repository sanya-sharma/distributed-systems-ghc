package entity

import (
	"github.com/jmcvetta/randutil"
)

var (
	PaymentGateways = []string{
		"Amex",
		"Zakpay",
		//"Weiss",
		//"Paypal",

	}
	PaymentGatewaysConfig = []randutil.Choice{
		{Item: "Weiss", Weight: 1},
		{Item: "Zakpay", Weight: 80},
		{Item: "Paypal", Weight: 0},
		{Item: "Amex", Weight: 2},
	}
)
