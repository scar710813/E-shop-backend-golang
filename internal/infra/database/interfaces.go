package database

import "github.com/PaoloProdossimoLopes/goshop/internal/entity"

type ProductRespository interface {
}

type UserRespository interface {
	Create(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}
