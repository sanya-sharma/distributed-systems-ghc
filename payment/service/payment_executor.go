package service

import (

)

type PaymentMethods interface {
	Execute()
}

type PaymentContext struct {
	executor PaymentMethods
}

func (p *PaymentContext) SetPaymentMethod(method PaymentMethods) {
	p.executor = method
}

func (p *PaymentContext) ExecutePayment() (err error) {

	err = p.executor.Execute()
	if err != nil {
		return err
	}
	
	return nil
}