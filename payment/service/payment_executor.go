package service

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
	return err
}

func (p *PaymentContext) RollbackPayment() (err error) {
	err = p.executor.Cancel()
	return err
}