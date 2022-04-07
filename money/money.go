package money

import (
	"errors"
)

type Money struct {
	Amount   uint32
	Currency string
}

func New(amount uint32, currency string) (*Money, error) {
	if currency == "" {
		return nil, errors.New("currency is required")
	}
	return &Money{amount, currency}, nil
}

func (m *Money) Add(n *Money) (*Money, error) {
	if n == nil {
		// panic("arugument must not be nil")
		return nil, errors.New("arugument must not be nil")

	}
	if m.Currency != n.Currency {
		// panic("currency must be same")
		return nil, errors.New("currency must be same")
	}
	return &Money{m.Amount + n.Amount, m.Currency}, nil
}
