package services

import (
	"context"

	"emptyslot/internal"
	"emptyslot/internal/models"
)

type CategoryRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (cr *CategoryRequest) ToModel(generateNewID bool) *models.Category {
	if generateNewID {
		cr.ID = generateUUID()
	}

	return &models.Category{
		ID:   cr.ID,
		Name: cr.Name,
	}
}

func (cr *CategoryRequest) Create(ctx context.Context) *models.Category {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(true)
	mr.CreateCategory(model)
	return model
}

func (cr *CategoryRequest) Update(ctx context.Context) *models.Category {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(false)
	mr.UpdateCategory(model)
	return model
}

func (cr *CategoryRequest) Get(ctx context.Context) []*models.Category {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	return mr.GetAllCategories()
}

func (cr *CategoryRequest) Detail(ctx context.Context, id string) *models.Category {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	return mr.GetCategoryByID(id)
}
