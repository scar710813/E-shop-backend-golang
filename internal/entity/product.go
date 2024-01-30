package entity

import (
	"errors"
	"time"

	"github.com/PaoloProdossimoLopes/goshop/pkg/entity"
)

var (
	ErrorInvalidPrice    = errors.New("invalid `price` (int) field")
	ErrorInvalidId       = errors.New("invalid `id` (string) field")
	ErrorIdIsRequired    = errors.New("missing `id` (string) field is required")
	ErrorNameIsRequired  = errors.New("missing `name` (string) field is required")
	ErrorPriceIsRequired = errors.New("missing `price` (int) field is required")
)

type Product struct {
	Id        entity.Id `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		Id:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if productValidationError := product.validate(); productValidationError != nil {
		return nil, productValidationError
	}

	return product, nil
}

func (p *Product) validate() error {
	if p.Id.String() == "" {
		return ErrorIdIsRequired
	}

	if _, err := entity.Parse(p.Id.String()); err != nil {
		return ErrorInvalidId
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}

	if p.Price == 0 {
		return ErrorPriceIsRequired
	}

	if p.Price < 0 {
		return ErrorInvalidPrice
	}

	return nil
}
