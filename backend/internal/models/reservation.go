package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	IsReserved      bool     `json:"isreserved"`
	MerchandiseIDs  []string `json:"merchandises"`
	CalculatedPrice int64    `json:"price"`
}

type ReservationRepository struct {
	Db *gorm.DB
}

func (r *ReservationRepository) CreateReservation(m *Reservation) error {
	if err := r.Db.Create(m).Error; err != nil {
		return fmt.Errorf("error creating reservation: %v", err)
	}

	return nil
}

func (r *ReservationRepository) UpdateReservation(m *Reservation) error {
	err := r.Db.Save(m).Error
	if err != nil {
		return fmt.Errorf("error updating reservation: %v", err)
	}

	return nil
}

func (r *ReservationRepository) GetReservationByID(id string) (*Reservation, error) {
	model := new(Reservation)
	if err := r.Db.First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error getting reservation: %v", err)
	}

	return model, nil
}

func (r *ReservationRepository) GetAllReservations() ([]*Reservation, error) {
	var list []*Reservation
	if err := r.Db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error getting reservations: %v", err)
	}
	return list, nil
}

func (r *ReservationRepository) DeleteReservation(id string) error {
	err := r.Db.Delete(&Reservation{}, id).Error
	if err != nil {
		return fmt.Errorf("error deleting reservation: %v", err)
	}

	return nil
}

func (r *ReservationRepository) DeleteReservations() error {
	err := r.Db.Delete(&Reservation{}).Error
	if err != nil {
		return fmt.Errorf("error deleting reservations: %v", err)
	}

	return nil
}
