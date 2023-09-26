package entity

import (
	"github.com/jmcvetta/randutil"
)

var (
	PaymentGatewayAmex   = "Amex"
	PaymentGatewayZakpay = "Zakpay"
	PaymentGatewayWeiss  = "Weiss"
	PaymentGatewayPaypal = "Paypal"

	PaymentGateways = []string{
		PaymentGatewayAmex,
		PaymentGatewayZakpay,
		//PaymentGatewayWeiss,
		//PaymentGatewayPaypal,

	}
	PaymentGatewaysConfig = []randutil.Choice{
		{Item: PaymentGatewayWeiss, Weight: 1},
		{Item: PaymentGatewayZakpay, Weight: 80},
		{Item: PaymentGatewayPaypal, Weight: 0},
		{Item: PaymentGatewayAmex, Weight: 2},
	}
)
