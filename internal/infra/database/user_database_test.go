package database

import (
	"testing"

	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open database: %v", err)
	}
	db.AutoMigrate(&entity.User{})

	domainUser, _ := entity.NewUser("Jhomn Doe", "j@j.com", "123456")
	databaseUser := NewUser(db)

	_, createUserInDatabaseError := databaseUser.Create(domainUser)
	assert.Nil(t, createUserInDatabaseError)

	var foundedUser entity.User
	findUserError := db.Find(&foundedUser, "id = ?", domainUser.Id).Error
	assert.Nil(t, findUserError)
	assert.Equal(t, domainUser.Id, foundedUser.Id)
	assert.Equal(t, domainUser.Name, foundedUser.Name)
	assert.Equal(t, domainUser.Email, foundedUser.Email)
	assert.NotNil(t, foundedUser.Password)
}
