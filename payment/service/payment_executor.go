package service

type PaymentGateways interface {
	Execute() bool
	Rollback() error
}

type PaymentContext struct {
	Client   string
	executor PaymentGateways
}

func (p *PaymentContext) SetPaymentMethod(method PaymentGateways) {
	p.executor = method
}

func (p *PaymentContext) ExecutePayment() (completed bool) {
	completed = p.executor.Execute()
	return completed
}

func (p *PaymentContext) RollbackPayment() (err error) {
	err = p.executor.Rollback()
	return err
}
