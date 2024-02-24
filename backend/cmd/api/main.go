package main

import (
	"context"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"backend/cmd/server"
	"backend/internal/auth"
	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/routes"
	sg "backend/stripegateway"
)

func main() {
	db := database.ConnectDb()
	models.Migration(db)
	r := mux.NewRouter()
	ctx := context.WithValue(context.Background(), "DB", db)
	routes.RegisterRoutes(ctx, r)
	auth.RegisterAuth(ctx, r)
	sg.Setup(ctx)
	handler := auth.AddCors(r)
	server.NewServer(handler)
}
