package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Category struct {
	gorm         gorm.Model
	ID           string `json:"id"`
	Name         string `json:"name"`
	Merchandises []*Merchandise
}

type CategoryRepository struct {
	Db *gorm.DB
}

func (r *CategoryRepository) CreateCategory(m *Category) {
	err := r.Db.Create(m).Error
	if err != nil {
		fmt.Println("Error creating category:", err)
	}
}

func (r *CategoryRepository) UpdateCategory(m *Category) {
	err := r.Db.Save(m).Error
	if err != nil {
		fmt.Println("Error updating category:", err)
	}
}

func (r *CategoryRepository) GetCategoryByID(id string) *Category {
	model := new(Category)
	if err := r.Db.First(model, id).Error; err != nil {
		fmt.Println("Error updating category:", err)
		return nil
	}
	return model
}

func (r *CategoryRepository) GetAllCategories() []*Category {
	var list []*Category
	if err := r.Db.Find(&list).Error; err != nil {
		fmt.Println("Error updating category:", err)
		return nil
	}
	return list
}

func (r *CategoryRepository) DeleteCategory(id string) {
	err := r.Db.Delete(&Category{}, id).Error
	if err != nil {
		fmt.Println("Error updating category:", err)
	}
}
