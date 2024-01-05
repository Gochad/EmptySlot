package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Customer struct {
	gorm gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
	User User   `json:"user"`
}

type CustomerRepository struct {
	Db *gorm.DB
}

func (r *CustomerRepository) CreateCustomer(m *Customer) error {
	if err := r.Db.Create(m).Error; err != nil {
		return fmt.Errorf("error creating customer: %v", err)
	}

	return nil
}

func (r *CustomerRepository) UpdateCustomer(m *Customer) error {
	if err := r.Db.Save(m).Error; err != nil {
		return fmt.Errorf("error updating customer: %v", err)
	}

	return nil
}

func (r *CustomerRepository) GetCustomerByID(id string) (*Customer, error) {
	model := new(Customer)
	if err := r.Db.First(model, id).Error; err != nil {
		return nil, fmt.Errorf("error updating customer: %v", err)
	}

	return model, nil
}

func (r *CustomerRepository) GetAllCustomers() ([]*Customer, error) {
	var list []*Customer
	if err := r.Db.Find(&list).Error; err != nil {
		return nil, fmt.Errorf("error updating customer: %v", err)
	}
	return list, nil
}

func (r *CustomerRepository) DeleteCustomer(id string) error {
	err := r.Db.Delete(&Customer{}, id).Error
	if err != nil {
		return fmt.Errorf("error updating customer: %v", err)
	}

	return nil
}
