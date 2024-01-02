package main

import (
	"context"
	"emptyslot/cmd/server"
	"emptyslot/internal/database"
	"emptyslot/internal/models"
	"emptyslot/internal/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	db := database.ConnectDb()
	models.Migration(db)
	r := mux.NewRouter()
	ctx := context.WithValue(context.Background(), "DB", db)
	routes.RegisterRoutes(ctx, r)

	server.NewServer(r)
}
