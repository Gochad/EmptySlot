package services

import (
	"emptyslot/internal/models"
	"github.com/google/uuid"
	"strconv"
)

type MerchandiseRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (mr *MerchandiseRequest) ToModel(generateNewID bool) *models.Merchandise {
	var id string
	if generateNewID {
		id = strconv.Itoa(int(uuid.New().ID()))
	} else {
		id = mr.ID
	}
	return &models.Merchandise{
		ID:   id,
		Name: mr.Name,
	}
}
