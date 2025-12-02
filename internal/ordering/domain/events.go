package domain

import (
	"time"

	"github.com/leanite/delivery-simulator/internal/common"
)

type OrderCreatedEvent struct {
	OrderID    common.ID
	CustomerID common.ID
	CreatedAt  time.Time
}

func NewOrderCreatedEvent(orderID common.ID, customerID common.ID) common.DomainEvent {
	return &OrderCreatedEvent{
		OrderID:    orderID,
		CustomerID: customerID,
		CreatedAt:  time.Now(),
	}
}
func (e OrderCreatedEvent) EventName() string { return "OrderCreated" }

type OrderItemAddedEvent struct {
	OrderID   common.ID
	ProductID common.ID
	Name      string
	Price     int64 // Value Object serializado para primitivo
	Quantity  int
}

func NewOrderItemAddedEvent(orderID common.ID, productID common.ID, name string, price int64, quantity int) common.DomainEvent {
	return &OrderItemAddedEvent{
		OrderID:   orderID,
		ProductID: productID,
		Name:      name,
		Price:     price,
		Quantity:  quantity,
	}
}
func (e OrderItemAddedEvent) EventName() string { return "OrderItemAdded" }

type OrderItemRemovedEvent struct {
	OrderID   common.ID
	ProductID common.ID
}

func NewOrderItemRemovedEvent(orderID common.ID, productID common.ID) common.DomainEvent {
	return &OrderItemRemovedEvent{
		OrderID:   orderID,
		ProductID: productID,
	}
}
func (e OrderItemRemovedEvent) EventName() string { return "OrderItemRemoved" }

type OrderConfirmedEvent struct {
	OrderID     common.ID
	ConfirmedAt time.Time
}

func NewOrderConfirmedEvent(orderID common.ID) common.DomainEvent {
	return &OrderConfirmedEvent{
		OrderID:     orderID,
		ConfirmedAt: time.Now(),
	}
}
func (e OrderConfirmedEvent) EventName() string { return "OrderConfirmed" }
