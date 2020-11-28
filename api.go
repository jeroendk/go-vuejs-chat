package main

import (
	"encoding/json"
	"net/http"

	"github.com/jeroendk/chatApplication/auth"
	"github.com/jeroendk/chatApplication/repository"
)

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type API struct {
	UserRepository *repository.UserRepository
}

func (api *API) HandleLogin(w http.ResponseWriter, r *http.Request) {

	var user LoginUser

	// Try to decode the JSON request to a LoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find the user in the database by username
	dbUser := api.UserRepository.FindUserByUsername(user.Username)
	if dbUser == nil {
		returnErrorResponse(w)
		return
	}

	// Check if the passwords match
	ok, err := auth.ComparePassword(user.Password, dbUser.Password)

	if !ok || err != nil {
		returnErrorResponse(w)
		return
	}

	// Create a JWT
	token, err := auth.CreateJWTToken(dbUser)

	if err != nil {
		returnErrorResponse(w)
		return
	}

	w.Write([]byte(token))

}

func returnErrorResponse(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"error\"}"))
}
