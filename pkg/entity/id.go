package entity

import "github.com/google/uuid"

type Id = uuid.UUID

func NewId() Id {
	return Id(uuid.New())
}

func Parse(id string) (Id, error) {
	parsedId, parseError := uuid.Parse(id)
	return Id(parsedId), parseError
}
