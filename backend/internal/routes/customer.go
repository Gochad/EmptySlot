package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/services"
	"backend/internal/views"
)

type historyImpl struct {
	ctx  context.Context
	body services.HistoryRequest
}

func registerHistory(ctx context.Context, router *mux.Router) {
	impl := &historyImpl{
		ctx: ctx,
	}
	s := router.PathPrefix("/history").Subrouter()
	s.HandleFunc("/", impl.create).Methods("POST")
	s.HandleFunc("/", impl.get).Methods("GET")
	s.HandleFunc("/{id}", impl.update).Methods("PUT")
	s.HandleFunc("/{id}", impl.detail).Methods("GET")
}

func (impl *historyImpl) create(w http.ResponseWriter, r *http.Request) {
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

func (impl *historyImpl) update(w http.ResponseWriter, r *http.Request) {
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

func (impl *historyImpl) get(w http.ResponseWriter, r *http.Request) {
	mods, err := impl.body.Get(impl.ctx)

	if err == nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *historyImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	historyID := vars["id"]

	model, err := impl.body.Detail(impl.ctx, historyID)

	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}
