package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         string  `json:"id" gorm:"unique_index"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Address    string  `json:"address"`
	Phone      string  `json:"phone"`
	CustomerID *string `json:"customer_id"`
}

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) CreateUser(m *User) error {
	if err := r.Db.Create(m).Error; err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	return nil
}

func (r *UserRepository) UpdateUser(m *User) error {
	if err := r.Db.Save(m).Error; err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

func (r *UserRepository) GetUserByID(id string) (*User, error) {
	model := new(User)
	if err := r.Db.First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return model, nil
}

func (r *UserRepository) GetAllUsers() ([]*User, error) {
	var list []*User
	if err := r.Db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}
	return list, nil
}

func (r *UserRepository) DeleteUser(id string) error {
	err := r.Db.Delete(&User{}, id).Error
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}
