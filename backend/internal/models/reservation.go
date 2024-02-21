package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm         gorm.Model
	ID           string        `json:"id"`
	Merchandises []Merchandise `gorm:"foreignKey:ReservationID" json:"merchandises"`
	Customer     Customer      `gorm:"foreignKey:ReservationID" json:"customer"`
	Confirmed    bool          `json:"confirmed"`
	IsReserved   bool          `json:"isreserved"`
	StartTime    string        `json:"starttime"`
	EndTime      string        `json:"endtime"`
}

type ReservationRepository struct {
	Db *gorm.DB
}

func (r *ReservationRepository) CreateReservation(m *Reservation) error {
	var validMerchandises []Merchandise
	for _, merchandise := range m.Merchandises {
		if merchandise.ID != "" && merchandise.Price != 0 {
			validMerchandises = append(validMerchandises, merchandise)
		}
	}
	m.Merchandises = validMerchandises

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
	if err := r.Db.Preload("Merchandises").First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error updating reservation: %v", err)
	}

	return model, nil
}

func (r *ReservationRepository) GetAllReservations() ([]*Reservation, error) {
	var list []*Reservation
	if err := r.Db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error updating reservation: %v", err)
	}
	return list, nil
}

func (r *ReservationRepository) DeleteReservation(id string) error {
	err := r.Db.Delete(&Reservation{}, id).Error
	if err != nil {
		return fmt.Errorf("error updating reservation: %v", err)
	}

	return nil
}
