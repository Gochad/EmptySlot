package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func isTokenValid(r *http.Request) error {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user_id"], claims["authorized"])
	} else {
		return err
	}

	return nil
}
