package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"backend/internal/services"
	"backend/internal/views"
)

type categoryImpl struct {
	ctx  context.Context
	body services.CategoryRequest
}

func registerCategory(ctx context.Context, router *mux.Router) {
	impl := &categoryImpl{
		ctx: ctx,
	}
	s := router.PathPrefix("/category").Subrouter()
	s.HandleFunc("/", impl.create).Methods("POST")
	s.HandleFunc("/", impl.get).Methods("GET")
	s.HandleFunc("/{id}", impl.update).Methods("PUT")
	s.HandleFunc("/{id}", impl.detail).Methods("GET")
}

func (impl *categoryImpl) create(w http.ResponseWriter, r *http.Request) {
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

func (impl *categoryImpl) update(w http.ResponseWriter, r *http.Request) {
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

func (impl *categoryImpl) get(w http.ResponseWriter, r *http.Request) {
	mods, err := impl.body.Get(impl.ctx)

	if err == nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *categoryImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	model, err := impl.body.Detail(impl.ctx, categoryID)

	if err == nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}
