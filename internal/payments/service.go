package payments

import "errors"

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (s *PaymentService) ProcessPayment(p Payment) error {
	return errors.New("payment processing not implemented")
}
