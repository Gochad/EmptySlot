package services

import (
	"context"
	"emptyslot/internal"
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

func Create(ctx context.Context, mreq *MerchandiseRequest) *models.Merchandise {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	model := mreq.ToModel(true)
	mr.CreateMerchandise(model)
	return model
}

func Update(ctx context.Context, mreq *MerchandiseRequest) *models.Merchandise {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	model := mreq.ToModel(false)
	mr.UpdateMerchandise(model)
	return model
}

func Get(ctx context.Context) []*models.Merchandise {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	return mr.GetAllMerchandise()
}

func Detail(ctx context.Context, id string) *models.Merchandise {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	return mr.GetMerchandiseByID(id)
}
