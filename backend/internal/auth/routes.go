package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuth(_ context.Context, r *mux.Router) {
	r.HandleFunc("/login", login).Methods()
	r.HandleFunc("/dashboard", validateMiddleware(dashboard)).Methods()
}

func login(w http.ResponseWriter, r *http.Request) {
	token, err := CreateToken("1234")
	if err != nil {
		fmt.Println(w, "Error Creating Token")
		return
	}
	fmt.Println(w, token)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Super Secret Information")
}

func validateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := isTokenValid(r)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			_, err2 := w.Write([]byte("Unauthorized: " + err.Error()))
			if err2 != nil {
				return
			}
			return
		}
		next.ServeHTTP(w, r)
	}
}
