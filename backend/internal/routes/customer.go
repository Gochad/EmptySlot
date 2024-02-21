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

type customerImpl struct {
	ctx context.Context
}

func registerCustomer(ctx context.Context, router *mux.Router) {
	impl := &customerImpl{
		ctx: ctx,
	}
	s := router.PathPrefix("/customers").Subrouter()
	s.HandleFunc("/", impl.create).Methods("POST")
	s.HandleFunc("/", impl.get).Methods("GET")
	s.HandleFunc("/{id}", impl.update).Methods("PUT")
	s.HandleFunc("/{id}", impl.detail).Methods("GET")
}

func (impl *customerImpl) create(w http.ResponseWriter, r *http.Request) {
	var body services.CustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.SendErrorMsg(w, "Error decoding JSON")
		return
	}
	model, err := body.Create(impl.ctx)
	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}

func (impl *customerImpl) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]
	fmt.Println(customerID)

	var body services.CustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.SendErrorMsg(w, "Error decoding JSON")
		return
	}
	model, err := body.Update(impl.ctx)

	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}

func (impl *customerImpl) get(w http.ResponseWriter, r *http.Request) {
	var body services.CustomerRequest

	mods, err := body.Get(impl.ctx)

	if err == nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *customerImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]

	var body services.CustomerRequest
	model, err := body.Detail(impl.ctx, customerID)

	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}
