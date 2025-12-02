package common

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

type DomainEvent interface {
	EventName() string
}
