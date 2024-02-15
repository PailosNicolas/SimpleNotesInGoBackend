package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PailosNicolas/SimpleNotesInGoBackend/helpers"
	"github.com/PailosNicolas/SimpleNotesInGoBackend/internal/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) HandlerCreateNewUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	if params.Username == "" || params.Password == "" {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Missing parameters.")
		return
	}

	encrypterPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating user.")
		return
	}

	newUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  params.Username,
		Password:  string(encrypterPass),
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating user.")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, newUser.GetUserDTO())

}
