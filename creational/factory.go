package creational

import (
	"errors"
	"fmt"
)

const (
	Cash       = 1
	CreditCard = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

type CashPaymentMethod struct{}
type CreditCardPaymentMethod struct{}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPaymentMethod), nil
	case CreditCard:
		return new(CreditCardPaymentMethod), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

func (c *CashPaymentMethod) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *CreditCardPaymentMethod) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using credit card\n", amount)
}
