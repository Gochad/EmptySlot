package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/services"
	"backend/internal/views"
)

type authImpl struct {
	ctx context.Context
}

func RegisterAuth(ctx context.Context, r *mux.Router) {
	impl := &authImpl{
		ctx: ctx,
	}

	r.HandleFunc("/login", impl.login)
	r.HandleFunc("/register", impl.register)
	r.HandleFunc("/dashboard", impl.validateMiddleware(impl.dashboard))
}

func (impl *authImpl) login(w http.ResponseWriter, r *http.Request) {
	var creds UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := generateJWT(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (impl *authImpl) register(w http.ResponseWriter, r *http.Request) {
	var newUser services.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, _ := HashPassword(newUser.Password)
	newUser.Password = hashedPassword

	create, err := newUser.Create(impl.ctx)
	if err != nil {
		return
	}

	if err == nil {
		views.SendResponse(w, create)
	} else {
		views.SendErrorMsg(w, create)
	}
}

func (impl *authImpl) dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Super Secret Information")
}

func (impl *authImpl) validateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, err := isTokenValid(r)
		if err != nil || !ok {
			w.WriteHeader(http.StatusForbidden)
			_, err2 := w.Write([]byte("Unauthorized"))
			if err2 != nil {
				return
			}
			return
		}
		next.ServeHTTP(w, r)
	}
}
