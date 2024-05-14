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
	r.HandleFunc("/register", impl.register).Methods("POST")

	r.HandleFunc("/google-sso", GoogleSignOn)
	r.HandleFunc("/callback", Callback)
}

func (impl *authImpl) login(w http.ResponseWriter, r *http.Request) {
	var creds UserCredentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var ureq services.UserRequest
	user, err := ureq.Detail(impl.ctx, creds.Email)
	if err != nil || !CheckPasswordHash(creds.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]any{"token": token, "reservation": user.ReservationID, "email": user.Email, "role": user.Role}); err != nil {
		fmt.Println("error during encoding token: ", err)
	}
}

func (impl *authImpl) register(w http.ResponseWriter, r *http.Request) {
	var newUser services.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usertype := r.URL.Query().Get("usertype")

	var rreq services.ReservationRequest
	reservation, _ := rreq.Create(impl.ctx)

	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newUser.Password = hashedPassword
	newUser.ReservationID = reservation.ID

	if usertype == "admin" {
		newUser.Role = 0 // admin
	} else {
		newUser.Role = 1 // standard
	}

	create, err := newUser.Create(impl.ctx)
	if err == nil {
		views.SendResponse(w, create)
	} else {
		views.SendErrorMsg(w, create)
	}
}
