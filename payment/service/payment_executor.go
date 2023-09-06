package service

type PaymentGateways interface {
	Execute()
}

type PaymentContext struct {
	Client   string
	executor PaymentGateways
}

func (p *PaymentContext) SetPaymentMethod(method PaymentGateways) {
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
