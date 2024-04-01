package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Merchandise struct {
	gorm.Model
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  *string `json:"category_id"`

	Price     int64  `json:"price"`
	Confirmed bool   `json:"confirmed"`
	StartTime string `json:"starttime"`
	EndTime   string `json:"endtime"`
}

type MerchandiseRepository struct {
	Db *gorm.DB
}

func (r *MerchandiseRepository) CreateMerchandise(m *Merchandise) error {
	if err := r.Db.Create(m).Error; err != nil {
		return fmt.Errorf("error creating merchandise: %v", err)
	}

	return nil
}

func (r *MerchandiseRepository) UpdateMerchandise(m *Merchandise) error {
	err := r.Db.Save(m).Error
	if err != nil {
		return fmt.Errorf("error updating merchandise: %v", err)
	}

	return nil
}

func (r *MerchandiseRepository) GetMerchandiseByID(id string) (*Merchandise, error) {
	model := new(Merchandise)
	if err := r.Db.First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error updating merchandise: %v", err)
	}

	return model, nil
}

func (r *MerchandiseRepository) GetAllMerchandise() ([]*Merchandise, error) {
	var list []*Merchandise
	if err := r.Db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error updating merchandise: %v", err)
	}
	return list, nil
}

func (r *MerchandiseRepository) DeleteMerchandise(id string) error {
	err := r.Db.Delete(&Merchandise{}, id).Error
	if err != nil {
		return fmt.Errorf("error updating merchandise: %v", err)
	}

	return nil
}
