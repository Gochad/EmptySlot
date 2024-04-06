package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string
	jwt.StandardClaims
}

func TokenValidationMiddleware(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[7:]
		} else {
			http.Error(w, "Missing or malformed JWT", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return SecretKey, fmt.Errorf("coś sie zjebalo")
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
