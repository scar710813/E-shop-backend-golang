package entity

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name, price := makeProductInputData()

	product, err := NewProduct(name, price)

	assert.Nil(t, err)
	assert.NotEmpty(t, product.Id.String())
	assert.Equal(t, product.Name, name)
	assert.Equal(t, product.Price, price)
}

func TestNewPriductShouldReturnErrorIfNameIsMissing(t *testing.T) {
	_, price := makeProductInputData()

	product, err := NewProduct("", price)

	assert.Nil(t, product)
	assert.Equal(t, err, ErrorNameIsRequired)
}

func TestNewPriductShouldReturnErrorIfPriceIsMissing(t *testing.T) {
	name, _ := makeProductInputData()

	product, err := NewProduct(name, 0)

	assert.Nil(t, product)
	assert.Equal(t, err, ErrorPriceIsRequired)
}

func TestNewPriductShouldReturnErrorIfPriceIsNegative(t *testing.T) {
	name, _ := makeProductInputData()

	product, err := NewProduct(name, -1)

	assert.Nil(t, product)
	assert.Equal(t, err, ErrorInvalidPrice)
}

func makeProductInputData() (name string, price int) {
	return uuid.New().String(), rand.Intn(1000)
}
