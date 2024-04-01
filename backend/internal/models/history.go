package models

import (
	"fmt"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	User          User    `gorm:"constraint:OnDelete:CASCADE;" json:"user"`
	ReservationID *string `json:"reservation_id"`
}

type HistoryRepository struct {
	Db *gorm.DB
}

func (r *HistoryRepository) CreateHistory(m *History) error {
	if err := r.Db.Create(m).Error; err != nil {
		return fmt.Errorf("error creating history: %v", err)
	}

	return nil
}

func (r *HistoryRepository) UpdateHistory(m *History) error {
	if err := r.Db.Save(m).Error; err != nil {
		return fmt.Errorf("error updating history: %v", err)
	}

	return nil
}

func (r *HistoryRepository) GetHistoryByID(id string) (*History, error) {
	model := new(History)
	if err := r.Db.First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error updating history: %v", err)
	}

	return model, nil
}

func (r *HistoryRepository) GetAllHistorys() ([]*History, error) {
	var list []*History
	if err := r.Db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error updating history: %v", err)
	}
	return list, nil
}

func (r *HistoryRepository) DeleteHistory(id string) error {
	err := r.Db.Delete(&History{}, id).Error
	if err != nil {
		return fmt.Errorf("error updating history: %v", err)
	}

	return nil
}
