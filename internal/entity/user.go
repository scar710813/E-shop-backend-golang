package entity

import (
	"github.com/PaoloProdossimoLopes/goshop/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       entity.Id `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, hashPasswordError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashPasswordError != nil {
		return nil, hashPasswordError
	}

	newUser := &User{
		Id:       entity.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	return newUser, nil
}

func (u *User) ValidatePassword(password string) bool {
	comparePasswordError := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return comparePasswordError == nil
}
