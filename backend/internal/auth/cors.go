package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func AddCors(r *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return c.Handler(r)
}
