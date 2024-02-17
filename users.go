package main

import (
	"database/sql"
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

	helpers.RespondWithJSON(w, http.StatusOK, newUser.GetDTO())

}

func (cfg *apiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type loginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
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

	user, err := cfg.DB.GetUserByUsername(r.Context(), params.Username)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Wrong username or password.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Wrong username or password.")
		return
	}

	accToken, RefToken, err := helpers.GetJWTTokenFromUser(&user, cfg.jwtSecret)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error generating tokens.")
		return
	}

	cfg.DB.UpdateTokens(r.Context(), database.UpdateTokensParams{
		Token: sql.NullString{
			String: accToken,
			Valid:  true,
		},
		RefreshToken: sql.NullString{
			String: RefToken,
			Valid:  true,
		},
		ID: user.ID,
	})

	helpers.RespondWithJSON(w, http.StatusOK, loginResponse{
		AccessToken:  accToken,
		RefreshToken: RefToken,
	})

}
