package routes

import (
	"context"
	"emptyslot/internal/services"
	"emptyslot/internal/views"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	model := services.Create(impl.ctx, &body)
	views.SendJSONResponse(w, http.StatusOK, model)
}

func (impl *merchandiseImpl) update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchandiseID := vars["id"]
	fmt.Println(merchandiseID)

	var body services.MerchandiseRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	model := services.Update(impl.ctx, &body)
	views.SendJSONResponse(w, http.StatusOK, model)
}

func (impl *merchandiseImpl) get(w http.ResponseWriter, r *http.Request) {
	mods := services.Get(impl.ctx)
	views.SendJSONResponse(w, http.StatusOK, mods)
}

func (impl *merchandiseImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchandiseID := vars["id"]
	model := services.Detail(impl.ctx, merchandiseID)
	views.SendJSONResponse(w, http.StatusOK, model)
}
