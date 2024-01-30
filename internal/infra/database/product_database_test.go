package database

import (
	"fmt"
	"testing"

	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10)
	prouctDatabase := NewProduct(db)

	_, createProductDatabaseError := prouctDatabase.Create(product)
	assert.NoError(t, createProductDatabaseError)

	assert.NotEmpty(t, product.Id)
}

func TestFindAllProduct(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	db.AutoMigrate(&entity.Product{})

	for i := 0; i < 24; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %v", i), i)
		db.Create(product)
	}

	prouctDatabase := NewProduct(db)
	products, createProductDatabaseError := prouctDatabase.FindAll(1, 10, "asc")
	assert.NoError(t, createProductDatabaseError)
	assert.Equal(t, len(products), 10)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, createProductDatabaseError = prouctDatabase.FindAll(2, 10, "asc")
	assert.NoError(t, createProductDatabaseError)
	assert.Equal(t, len(products), 10)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")

	products, createProductDatabaseError = prouctDatabase.FindAll(3, 10, "asc")
	assert.NoError(t, createProductDatabaseError)
	assert.Equal(t, len(products), 3)
	assert.Equal(t, products[0].Name, "Product 21")
	assert.Equal(t, products[2].Name, "Product 23")
}
