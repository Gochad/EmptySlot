package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

var mySigningKey = []byte("secret")

func isTokenValid(r *http.Request) (bool, error) {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, fmt.Errorf("invalid token")
}
