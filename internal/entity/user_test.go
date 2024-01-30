package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserShouldNotFail(t *testing.T) {
	name, email, password := makeUserInputData()

	_, err := NewUser(name, email, password)

	assert.Nil(t, err)
}

func TestCreateUserSetIdCorrecy(t *testing.T) {
	name, email, password := makeUserInputData()

	createdUser, _ := NewUser(name, email, password)

	assert.NotEmpty(t, createdUser.Id.String())
}

func TestUserCreationShouldSetNameCorrectly(t *testing.T) {
	name, email, password := makeUserInputData()

	createdUser, _ := NewUser(name, email, password)

	assert.Equal(t, createdUser.Name, name)
}

func TestUserCreationShouldSetEmailCorrectly(t *testing.T) {
	name, email, password := makeUserInputData()

	createdUser, _ := NewUser(name, email, password)

	assert.Equal(t, createdUser.Email, email)
}

func TestPasswordGeneration(t *testing.T) {
	name, email, password := makeUserInputData()

	createdUser, _ := NewUser(name, email, password)

	assert.NotEqual(t, createdUser.Password, password)
	assert.NotEmpty(t, createdUser.Password)
	assert.True(t, createdUser.ValidatePassword(password))
}

func makeUserInputData() (n string, e string, p string) {
	return "John Doe", "j@j.com", "123456"
}
