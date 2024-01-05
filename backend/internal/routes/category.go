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

type categoryImpl struct {
	ctx context.Context
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
	var body services.CategoryRequest

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

func (impl *categoryImpl) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]
	fmt.Println(categoryID)

	var body services.CategoryRequest

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

func (impl *categoryImpl) get(w http.ResponseWriter, r *http.Request) {
	var body services.CategoryRequest

	mods, err := body.Get(impl.ctx)

	if err != nil {
		views.SendResponse(w, mods)
	} else {
		views.SendErrorMsg(w, mods)
	}
}

func (impl *categoryImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var body services.CategoryRequest
	model, err := body.Detail(impl.ctx, categoryID)

	if err != nil {
		views.SendResponse(w, model)
	} else {
		views.SendErrorMsg(w, model)
	}
}
