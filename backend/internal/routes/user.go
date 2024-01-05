package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/services"
	"backend/internal/views"
)

type userImpl struct {
	ctx context.Context
}

func registerUser(ctx context.Context, router *mux.Router) {
	impl := &userImpl{
		ctx: ctx,
	}
	s := router.PathPrefix("/users").Subrouter()
	s.HandleFunc("/", impl.create).Methods("POST")
	s.HandleFunc("/", impl.get).Methods("GET")
	s.HandleFunc("/{id}", impl.update).Methods("PUT")
	s.HandleFunc("/{id}", impl.detail).Methods("GET")
}

func (impl *userImpl) create(w http.ResponseWriter, r *http.Request) {
	var body services.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	model, err := body.Create(impl.ctx)
	if err != nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}

func (impl *userImpl) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Println(userID)

	var body services.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	model, err := body.Update(impl.ctx)

	if err != nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}

func (impl *userImpl) get(w http.ResponseWriter, r *http.Request) {
	var body services.UserRequest

	mods, err := body.Get(impl.ctx)

	if err != nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *userImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	var body services.UserRequest
	model, err := body.Detail(impl.ctx, userID)

	if err != nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}
