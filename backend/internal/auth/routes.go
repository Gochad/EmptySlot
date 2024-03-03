package auth

import (
	"context"
	"encoding/json"
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
