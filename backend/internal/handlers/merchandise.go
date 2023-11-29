package handlers

import (
	m "emptyslot/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MerchandiseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func MerchandisesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	Merchandises := []m.Merchandise{
		*m.NewMerchandise("1", "goch"),
		*m.NewMerchandise("2", "tjzel"),
	}
	json.NewEncoder(w).Encode(Merchandises)
}
