package entity

import "github.com/google/uuid"

type Id uuid.UUID

func New() Id {
	return Id(uuid.New())
}

func Parse(id string) (Id, error) {
	parsedId, parseError := uuid.Parse(id)
	return Id(parsedId), parseError
}

func (id Id) String() string {
	return uuid.UUID(id).String()
}
