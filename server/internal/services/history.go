package services

import (
	"context"

	"backend/internal"
	"backend/internal/models"
)

type HistoryRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user"`
}

func (cr *HistoryRequest) ToModel(generateNewID bool) *models.History {
	if generateNewID {
		cr.ID = generateUUID()
	}

	return &models.History{
		ID:     cr.ID,
		UserID: cr.UserID,
	}
}

func (cr *HistoryRequest) Create(ctx context.Context) (*models.History, error) {
	mr := models.HistoryRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(true)
	err := mr.CreateHistory(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (cr *HistoryRequest) Update(ctx context.Context) (*models.History, error) {
	mr := models.HistoryRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(false)
	err := mr.UpdateHistory(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (cr *HistoryRequest) Get(ctx context.Context) ([]*models.History, error) {
	mr := models.HistoryRepository{Db: internal.Database(ctx)}
	return mr.GetAllHistorys()
}

func (cr *HistoryRequest) Detail(ctx context.Context, id string) (*models.History, error) {
	mr := models.HistoryRepository{Db: internal.Database(ctx)}
	return mr.GetHistoryByID(id)
}
