package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	registerMerchandise(r, db)
}
