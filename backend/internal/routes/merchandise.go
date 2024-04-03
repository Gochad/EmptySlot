package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/services"
	"backend/internal/views"
)

type merchandiseImpl struct {
	ctx  context.Context
	body services.MerchandiseRequest
}

func registerMerchandise(ctx context.Context, router *mux.Router) {
	impl := &merchandiseImpl{
		ctx:  ctx,
		body: services.MerchandiseRequest{},
	}
	s := router.PathPrefix("/merchandises").Subrouter()
	s.HandleFunc("/", impl.create).Methods("POST")
	s.HandleFunc("/", impl.get).Methods("GET")
	s.HandleFunc("/", impl.update).Methods("PUT")
	s.HandleFunc("/{id}", impl.detail).Methods("GET")

	s = router.PathPrefix("/merchandisesfromreservation").Subrouter()
	s.HandleFunc("/{id}", impl.getByReservation).Methods("GET")
}

func (impl *merchandiseImpl) create(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&impl.body); err != nil {
		views.SendErrorMsg(w, "Error decoding JSON")
		return
	}
	model, err := impl.body.Create(impl.ctx)
	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}

func (impl *merchandiseImpl) update(w http.ResponseWriter, r *http.Request) {
	log.Print("robie create")
	if err := json.NewDecoder(r.Body).Decode(&impl.body); err != nil {
		views.SendErrorMsg(w, "Error decoding JSON")
		return
	}
	model, err := impl.body.Update(impl.ctx)

	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}

func (impl *merchandiseImpl) get(w http.ResponseWriter, r *http.Request) {
	mods, err := impl.body.Get(impl.ctx)

	if err == nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *merchandiseImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchandiseID := vars["id"]

	model, err := impl.body.Detail(impl.ctx, merchandiseID)

	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}

func (impl *merchandiseImpl) getByReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID := vars["id"]

	mods, err := services.GetMerchByReservationID(impl.ctx, reservationID)

	if err == nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}
