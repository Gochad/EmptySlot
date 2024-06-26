package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Role int

const (
	Admin Role = iota
	Normal
)

type User struct {
	gorm.Model
	ID       string `json:"id" gorm:"unique_index"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Role     int    `json:"role"`

	ReservationID string `json:"reservation"`
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

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	model := new(User)
	if err := r.Db.Where("email = ?", email).First(&model).Error; err != nil {
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
