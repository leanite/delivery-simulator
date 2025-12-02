package common

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

type Entity interface {
	ID() ID
}

type ValueObject[T any] interface {
	Compare(other T) bool
}

type DomainEvent interface {
	EventName() string
}
