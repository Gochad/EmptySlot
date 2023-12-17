package internal

import (
	h "emptyslot/internal/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router) {
	s := r.PathPrefix("/merchandises").Subrouter()

	s.HandleFunc("/{key}/", h.MerchandiseHandler)
	s.HandleFunc("/", h.MerchandisesHandler)
}
