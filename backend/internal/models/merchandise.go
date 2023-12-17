package models

import "gorm.io/gorm"

type Merchandise struct {
	gorm gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewMerchandise(id, name string) *Merchandise {
	return &Merchandise{
		ID:   id,
		Name: name,
	}
}
