package routes

import (
	"emptyslot/internal/models"
	"emptyslot/internal/services"
	"emptyslot/internal/views"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type merchandiseImpl struct {
	db *gorm.DB
}

func registerMerchandise(router *mux.Router, db *gorm.DB) {
	impl := &merchandiseImpl{
		db: db,
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
	mr := models.MerchandiseRepository{Db: impl.db}
	model := body.ToModel()
	mr.CreateMerchandise(model)
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
	mr := models.MerchandiseRepository{Db: impl.db}
	model := body.ToModel()
	mr.UpdateMerchandise(model)
	views.SendJSONResponse(w, http.StatusOK, model)
}

func (impl *merchandiseImpl) get(w http.ResponseWriter, r *http.Request) {
	mr := models.MerchandiseRepository{Db: impl.db}
	mods := mr.GetAllMerchandise()
	views.SendJSONResponse(w, http.StatusOK, mods)
}

func (impl *merchandiseImpl) detail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchandiseID := vars["id"]

	mr := models.MerchandiseRepository{Db: impl.db}
	model := mr.GetMerchandiseByID(merchandiseID)
	views.SendJSONResponse(w, http.StatusOK, model)
}
