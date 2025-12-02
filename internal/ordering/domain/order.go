package domain

import (
	"errors"
	"time"
)

type OrderStatus string

const (
	OrderStatusDraft     OrderStatus = "DRAFT"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
)

type Order struct {
	id         string
	customerId string
	items      []LineItem
	status     OrderStatus
	createdAt  time.Time
}

type LineItem struct {
	ProductID string
	Name      string
	Quantity  int
	UnitPrice Money
}

func NewOrder(id string, customerId string) (*Order, error) {
	if id == "" {
		return nil, errors.New("order ID is required")
	}
	if customerId == "" {
		return nil, errors.New("customer ID is required")
	}

	return &Order{
		id:         id,
		customerId: customerId,
		status:     OrderStatusDraft,
		createdAt:  time.Now(),
		items:      []LineItem{},
	}, nil
}

func (o *Order) AddItem(productID, name string, price Money, qty int) error {
	if o.status != OrderStatusDraft {
		return errors.New("cannot add items to a confirmed order")
	}
	if qty <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	item := LineItem{
		ProductID: productID,
		Name:      name,
		UnitPrice: price,
		Quantity:  qty,
	}

	o.items = append(o.items, item)
	return nil
}

func (o *Order) TotalPrice() (Money, error) {
	total, _ := NewMoney(0, "BRL") //TODO: enum de currency?

	for _, item := range o.items {
		itemTotalAmount := item.UnitPrice.amount * int64(item.Quantity)
		itemTotal, _ := NewMoney(itemTotalAmount, item.UnitPrice.currency)

		var err error
		total, err = total.Add(itemTotal)
		if err != nil {
			return Money{}, err
		}
	}
	return total, nil
}

func (o *Order) Confirm() error {
	if len(o.items) <= 0 {
		return errors.New("cannot confirm empty order")
	}

	o.status = OrderStatusConfirmed
	return nil
}
