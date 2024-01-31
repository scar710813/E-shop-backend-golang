package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/goshop/internal/dto"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
)

type UserHandler struct {
	UserDatabase database.UserRespository
}

func NewUserHandler(database database.UserRespository) *UserHandler {
	return &UserHandler{
		UserDatabase: database,
	}
}

func (self *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest dto.CreateUserINput
	jsonDecoderUserError := json.NewDecoder(r.Body).Decode(&userRequest)
	if jsonDecoderUserError != nil {
		http.Error(w, jsonDecoderUserError.Error(), http.StatusBadRequest)
		return
	}

	user, createUserError := entity.NewUser(userRequest.Name, userRequest.Email, userRequest.Password)
	if createUserError != nil {
		http.Error(w, createUserError.Error(), http.StatusBadRequest)
		return
	}

	userCreated, createUserError := self.UserDatabase.Create(user)
	if createUserError != nil {
		http.Error(w, createUserError.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userCreated)
}
