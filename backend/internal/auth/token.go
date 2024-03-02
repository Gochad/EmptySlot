package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var SecretKey = []byte("your_secret_key")

func generateJWT(user UserCredentials) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "EmptySlot",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}