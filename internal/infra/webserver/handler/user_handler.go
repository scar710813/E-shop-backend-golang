package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PaoloProdossimoLopes/goshop/internal/dto"
	"github.com/PaoloProdossimoLopes/goshop/internal/entity"
	"github.com/PaoloProdossimoLopes/goshop/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDatabase database.UserRespository
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDatabase database.UserRespository, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserDatabase: userDatabase,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
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

func (self *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var credentialsRequest dto.GetJwtInput
	jsonDecoderUserError := json.NewDecoder(r.Body).Decode(&credentialsRequest)
	if jsonDecoderUserError != nil {
		http.Error(w, jsonDecoderUserError.Error(), http.StatusBadRequest)
		return
	}

	findedUser, findUserError := self.UserDatabase.FindByEmail(credentialsRequest.Email)
	if findUserError != nil {
		http.Error(w, findUserError.Error(), http.StatusNotFound)
		return
	}

	if !findedUser.ValidatePassword(credentialsRequest.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	_, token, createTokenError := self.Jwt.Encode(map[string]interface{}{
		"sub":        findedUser.Id.String(),
		"expires_in": time.Now().Add(time.Second * time.Duration(self.JwtExpiresIn)).Unix(),
	})
	if createTokenError != nil {
		http.Error(w, createTokenError.Error(), http.StatusInternalServerError)
		return
	}

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}
