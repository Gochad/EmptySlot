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

type merchandiseImpl struct {
	ctx context.Context
}

func registerMerchandise(ctx context.Context, router *mux.Router) {
	impl := &merchandiseImpl{
		ctx: ctx,
	}
	s := router.PathPrefix("/merchandises").Subrouter()
	s.HandleFunc("/", impl.create).Methods("POST")
	s.HandleFunc("/", impl.get).Methods("GET")
	s.HandleFunc("/{id}", impl.update).Methods("PUT")
	s.HandleFunc("/{id}", impl.detail).Methods("GET")
}

func (impl *merchandiseImpl) create(w http.ResponseWriter, r *http.Request) {
	var body services.MerchandiseRequest

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

func (impl *merchandiseImpl) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchandiseID := vars["id"]
	fmt.Println(merchandiseID)

	var body services.MerchandiseRequest

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

func (impl *merchandiseImpl) get(w http.ResponseWriter, r *http.Request) {
	var body services.MerchandiseRequest

	mods, err := body.Get(impl.ctx)

	if err == nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *merchandiseImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchandiseID := vars["id"]

	var body services.MerchandiseRequest
	model, err := body.Detail(impl.ctx, merchandiseID)

	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}
