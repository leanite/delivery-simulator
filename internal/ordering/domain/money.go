package domain

import (
	"errors"
	"fmt"

	"github.com/leanite/delivery-simulator/internal/common"
)

type Money struct {
	amount   int64  // valor em centavos
	currency string // ex: "BRL", "USD"
}

var _ common.ValueObject[Money] = (*Money)(nil)

func NewMoney(amount int64, currency string) (Money, error) {
	if currency == "" {
		return Money{}, errors.New("currency is required")
	}
	if amount < 0 {
		return Money{}, errors.New("amount cannot be negative")
	}

	return Money{amount: amount, currency: currency}, nil
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("cannot add different currencies: %s and %s", m.currency, other.currency)
	}
	return Money{
		amount:   m.amount + other.amount,
		currency: m.currency,
	}, nil
}

func (m Money) Sub(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("cannot subtract different currencies: %s and %s", m.currency, other.currency)
	}
	return Money{
		amount:   m.amount - other.amount,
		currency: m.currency,
	}, nil
}

func (m Money) String() string {
	return fmt.Sprintf("%s %.2f", m.currency, float64(m.amount)/100)
}

func (m *Money) Compare(other Money) bool {
	return (m.amount == other.amount) && (m.currency == other.currency)
}
