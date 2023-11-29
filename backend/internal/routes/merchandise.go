package routes

import (
	m "emptyslot/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnSingleMerchandise(w http.ResponseWriter, r *http.Request) {
	ms := []m.Merchandise{
		{ID: "1", Name: "F"},
		{ID: "2", Name: "F"},
	}
	key := r.URL.Query().Get("id")

	for _, m := range ms {
		if m.ID == key {
			json.NewEncoder(w).Encode(m)
		}
	}
	fmt.Fprintf(w, "Hello, this is a simple HTTP server written in Golang!")
}
