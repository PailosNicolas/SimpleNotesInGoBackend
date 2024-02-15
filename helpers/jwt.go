package helpers

import (
	"time"

	"github.com/PailosNicolas/SimpleNotesInGoBackend/internal/database"
	"github.com/golang-jwt/jwt/v5"
)

func GetJWTTokenFromUser(user *database.User, jwtSecret string) (access string, refresh string, err error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		Issuer:    "access",
		Subject:   user.ID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	claimsRefresh := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 60)),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		Issuer:    "refresh",
		Subject:   user.ID.String(),
	}

	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", "", err
	}

	signedRefreshToken, err := tokenRefresh.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", "", err
	}

	return signedToken, signedRefreshToken, nil
}
