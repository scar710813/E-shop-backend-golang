package database

import "github.com/PaoloProdossimoLopes/goshop/internal/entity"

type ProductRespository interface {
	Create(product *entity.Product) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id string) error
}

type UserRespository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
