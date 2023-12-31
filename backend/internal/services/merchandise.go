package services

import "emptyslot/internal/models"

type MerchandiseRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (mr *MerchandiseRequest) ToModel() *models.Merchandise {
	return &models.Merchandise{
		ID:   mr.ID,
		Name: mr.Name,
	}
}
