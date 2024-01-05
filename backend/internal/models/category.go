package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm         gorm.Model
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Merchandises []*Merchandise `json:"merchandises"`
}

type CategoryRepository struct {
	Db *gorm.DB
}

func (r *CategoryRepository) CreateCategory(m *Category) error {
	if err := r.Db.Create(m).Error; err != nil {
		return fmt.Errorf("error creating category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(m *Category) error {
	if err := r.Db.Save(m).Error; err != nil {
		return fmt.Errorf("error updating category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) GetCategoryByID(id string) (*Category, error) {
	model := new(Category)
	if err := r.Db.First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error updating category: %v", err)
	}

	return model, nil
}

func (r *CategoryRepository) GetAllCategories() ([]*Category, error) {
	var list []*Category
	if err := r.Db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error updating category: %v", err)
	}
	return list, nil
}

func (r *CategoryRepository) DeleteCategory(id string) error {
	err := r.Db.Delete(&Category{}, id).Error
	if err != nil {
		return fmt.Errorf("error updating category: %v", err)
	}

	return nil
}
