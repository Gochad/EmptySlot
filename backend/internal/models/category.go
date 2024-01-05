package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm         gorm.Model
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Merchandises []*Merchandise `gorm:"foreignKey:CategoryID" json:"merchandises"`
}

type CategoryRepository struct {
	Db *gorm.DB
}

func (r *CategoryRepository) CreateCategory(c *Category) error {
	if err := r.Db.Create(c).Error; err != nil {
		return fmt.Errorf("error creating category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(c *Category) error {
	if err := r.Db.Save(c).Error; err != nil {
		return fmt.Errorf("error updating category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) GetCategoryByID(id string) (*Category, error) {
	model := new(Category)
	if err := r.Db.Preload("Merchandises").First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error updating category: %v", err)
	}

	return model, nil
}

func (r *CategoryRepository) GetAllCategories() ([]*Category, error) {
	var list []*Category
	if err := r.Db.Preload("Merchandises").Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error updating category: %v", err)
	}

	for i, elem := range list {
		fmt.Println(i, "-ty element: ", elem.Merchandises)
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
