package domain

import (
	"errors"
	"time"

	"github.com/leanite/delivery-simulator/internal/common"
)

type OrderStatus string

const (
	OrderStatusDraft     OrderStatus = "DRAFT"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
)

type Order struct {
	id         common.ID
	customerID common.ID
	items      []LineItem
	status     OrderStatus
	createdAt  time.Time
	changes    []common.DomainEvent
}

type LineItem struct {
	ProductID common.ID
	Name      string
	Quantity  int
	UnitPrice Money
}

func NewOrder(id common.ID, customerID common.ID) (*Order, error) {
	createdAt := time.Now()

	event := OrderCreatedEvent{
		OrderID:    id,
		CustomerID: customerID,
		CreatedAt:  createdAt,
	}
	order := Order{
		id:         id,
		customerID: customerID,
		status:     OrderStatusDraft,
		createdAt:  createdAt,
		items:      []LineItem{},
	}
	order.apply(event)

	return &order, nil
}

func (o *Order) AddItem(productID common.ID, name string, price Money, qty int) error {
	if o.status != OrderStatusDraft {
		return errors.New("cannot add items to a confirmed order")
	}
	if qty <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	event := NewOrderItemAddedEvent(o.id, productID, name, price.amount, qty)
	o.apply(event)

	return nil
}

func (o *Order) RemoveItem(productID common.ID) error {
	if o.status != OrderStatusDraft {
		return errors.New("cannot remove items to a confirmed order")
	}

	exists := false
	for _, item := range o.items {
		if item.ProductID == productID {
			exists = true
			break
		}
	}
	if !exists {
		return errors.New("item not found")
	}

	event := NewOrderItemRemovedEvent(o.id, productID)
	o.apply(event)

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

	event := NewOrderConfirmedEvent(o.id)
	o.apply(event)

	return nil
}

func (o *Order) apply(event common.DomainEvent) {
	o.changes = append(o.changes, event)

	switch e := event.(type) {
	case OrderConfirmedEvent:
		o.status = OrderStatusConfirmed
	case OrderCreatedEvent:
		o.id = e.OrderID
		o.customerID = e.CustomerID
		o.status = OrderStatusDraft
	case OrderItemAddedEvent:
		unitPrice, _ := NewMoney(e.Price, "BRL")
		item := LineItem{
			ProductID: e.ProductID,
			Name:      e.Name,
			UnitPrice: unitPrice,
			Quantity:  e.Quantity,
		}
		o.items = append(o.items, item)
	case OrderItemRemovedEvent:
		newItems := o.items[:0]
		for _, item := range o.items {
			if item.ProductID != e.ProductID {
				newItems = append(newItems, item)
			}
		}
		o.items = newItems
	}
}
