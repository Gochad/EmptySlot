package main

import (
	"emptyslot/cmd/server"
	"emptyslot/internal/database"
	"emptyslot/internal/models"
	"emptyslot/internal/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func main() {
	app := App{}
	app.DB = database.ConnectDb()
	models.Migration(app.DB)
	r := mux.NewRouter()
	routes.RegisterRoutes(r, app.DB)
	server.NewServer(r)
}
