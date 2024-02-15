package main

import (
	"net/http"
	"strings"

	"github.com/PailosNicolas/SimpleNotesInGoBackend/helpers"
	"github.com/PailosNicolas/SimpleNotesInGoBackend/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtAuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler jwtAuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		bearerToken = strings.ReplaceAll(bearerToken, "Bearer ", "")

		if bearerToken == "" {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		claims := &jwt.RegisteredClaims{}

		parseFunc := func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.jwtSecret), nil
		}

		jwtToken, err := jwt.ParseWithClaims(bearerToken, claims, parseFunc)

		if err != nil {
			helpers.RespondWithError(w, 401, "Unauthorized")
			return
		}

		issuer, err := jwtToken.Claims.GetIssuer()

		if err != nil {
			helpers.RespondWithError(w, 401, "Error geting issuer")
			return
		}

		if issuer != "access" {
			helpers.RespondWithError(w, 401, "Unauthorized")
			return
		}

		userIDstr, err := jwtToken.Claims.GetSubject()

		if err != nil {
			helpers.RespondWithError(w, 401, "Error geting id")
			return
		}

		userId, err := uuid.Parse(userIDstr)
		if err != nil {
			helpers.RespondWithError(w, 401, "Error parsing id")
			return
		}

		user, err := cfg.DB.GetUserById(r.Context(), userId)

		if err != nil {
			helpers.RespondWithError(w, 401, "Error getting user")
			return
		}

		handler(w, r, user)
	}
}
