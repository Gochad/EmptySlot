package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Merchandise struct {
	gorm gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MerchandiseRepository struct {
	Db *gorm.DB
}

func NewMerchandise(id, name string) *Merchandise {
	return &Merchandise{
		ID:   id,
		Name: name,
	}
}
func (r *MerchandiseRepository) CreateMerchandise(m *Merchandise) {
	err := r.Db.Create(m).Error
	if err != nil {
		fmt.Println("Error creating merchandise:", err)
	}
}

func (r *MerchandiseRepository) UpdateMerchandise(m *Merchandise) {
	err := r.Db.Save(m).Error
	if err != nil {
		fmt.Println("Error updating merchandise:", err)
	}
}
func (r *MerchandiseRepository) GetMerchandiseByID(id string) *Merchandise {
	model := new(Merchandise)
	if err := r.Db.First(model, id).Error; err != nil {
		fmt.Println("Error updating merchandise:", err)
		return nil
	}
	return model
}
func (r *MerchandiseRepository) GetAllMerchandise() []*Merchandise {
	var list []*Merchandise
	if err := r.Db.Find(&list).Error; err != nil {
		fmt.Println("Error updating merchandise:", err)
		return nil
	}
	return list
}
func (r *MerchandiseRepository) DeleteMerchandise(id string) {
	err := r.Db.Delete(&Merchandise{}, id).Error
	if err != nil {
		fmt.Println("Error updating merchandise:", err)
	}
}
