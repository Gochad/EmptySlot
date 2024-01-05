package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"emptyslot/internal/services"
	"emptyslot/internal/views"
)

type reservationImpl struct {
	ctx context.Context
}

func registerReservation(ctx context.Context, router *mux.Router) {
	impl := &reservationImpl{
		ctx: ctx,
	}
	s := router.PathPrefix("/reservations").Subrouter()
	s.HandleFunc("/", impl.create).Methods("POST")
	s.HandleFunc("/", impl.get).Methods("GET")
	s.HandleFunc("/{id}", impl.update).Methods("PUT")
	s.HandleFunc("/{id}", impl.detail).Methods("GET")
}

func (impl *reservationImpl) create(w http.ResponseWriter, r *http.Request) {
	var body services.ReservationRequest

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

func (impl *reservationImpl) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID := vars["id"]
	fmt.Println(reservationID)

	var body services.ReservationRequest

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

func (impl *reservationImpl) get(w http.ResponseWriter, r *http.Request) {
	var body services.ReservationRequest

	mods, err := body.Get(impl.ctx)

	if err != nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *reservationImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID := vars["id"]

	var body services.ReservationRequest
	model, err := body.Detail(impl.ctx, reservationID)

	if err != nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}
