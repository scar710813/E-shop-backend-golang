package database

import (
	"errors"

	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) (*entity.Product, error) {
	return product, p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	if err := p.DB.Offset((page - 1) * limit).Limit(limit).Order("created_at " + sort).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Update(product *entity.Product) (*entity.Product, error) {
	_, findByIdError := p.FindById(product.Id.String())
	if findByIdError != nil {
		return nil, errors.New("product not found")
	}

	return product, p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	return p.DB.Delete(&entity.Product{}, id).Error
}
