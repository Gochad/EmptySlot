package services

import (
	"context"

	"backend/internal"
	"backend/internal/models"
)

type CategoryRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (cr *CategoryRequest) ToModel(generateNewID bool) *models.Category {
	if generateNewID {
		cr.ID = generateUUID()
	}

	return &models.Category{
		ID:    cr.ID,
		Name:  cr.Name,
		Color: cr.Color,
	}
}

func (cr *CategoryRequest) Create(ctx context.Context) (*models.Category, error) {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(true)
	err := mr.CreateCategory(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (cr *CategoryRequest) Update(ctx context.Context) (*models.Category, error) {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(false)
	err := mr.UpdateCategory(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (cr *CategoryRequest) Get(ctx context.Context) ([]*models.Category, error) {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	return mr.GetAllCategories()
}

func (cr *CategoryRequest) Detail(ctx context.Context, id string) (*models.Category, error) {
	mr := models.CategoryRepository{Db: internal.Database(ctx)}
	return mr.GetCategoryByID(id)
}
