package services

import (
	"context"

	"backend/internal"
	"backend/internal/models"
)

type MerchandiseRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	CategoryID  uint   `json:"category_id"`
}

func (mreq *MerchandiseRequest) ToModel(generateNewID bool, categoryId string) *models.Merchandise {
	if generateNewID {
		mreq.ID = generateUUID()
	}
	return &models.Merchandise{
		ID:          mreq.ID,
		Name:        mreq.Name,
		Description: mreq.Description,
		Price:       mreq.Price,
		CategoryID:  &categoryId,
	}
}

func (mreq *MerchandiseRequest) Create(ctx context.Context) (*models.Merchandise, error) {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	model := mreq.ToModel(true, "")
	err := mr.CreateMerchandise(model)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (mreq *MerchandiseRequest) Update(ctx context.Context) (*models.Merchandise, error) {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	model := mreq.ToModel(false, "")
	err := mr.UpdateMerchandise(model)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (mreq *MerchandiseRequest) Get(ctx context.Context) ([]*models.Merchandise, error) {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	return mr.GetAllMerchandise()
}

func (mreq *MerchandiseRequest) Detail(ctx context.Context, id string) (*models.Merchandise, error) {
	mr := models.MerchandiseRepository{Db: internal.Database(ctx)}
	return mr.GetMerchandiseByID(id)
}
