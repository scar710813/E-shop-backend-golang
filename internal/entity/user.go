package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func New(name, email, password string) (*User, error) {
	hashedPassword, hashPasswordError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashPasswordError != nil {
		return nil, hashPasswordError
	}

	newUserId := uuid.New().String()
	newUser := &User{
		Id:       newUserId,
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	return newUser, nil
}

func (u *User) ValidatePassword(password string) bool {
	comparePasswordError := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return comparePasswordError != nil
}
