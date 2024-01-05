package services

import (
	"context"

	"backend/internal"
	"backend/internal/models"
)

type CustomerRequest struct {
	ID   string      `json:"id"`
	User UserRequest `json:"user"`
}

func (cr *CustomerRequest) ToModel(generateNewID bool) *models.Customer {
	if generateNewID {
		cr.ID = generateUUID()
	}

	return &models.Customer{
		ID:   cr.ID,
		User: *cr.User.ToModel(generateNewID),
	}
}

func (cr *CustomerRequest) Create(ctx context.Context) (*models.Customer, error) {
	mr := models.CustomerRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(true)
	err := mr.CreateCustomer(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (cr *CustomerRequest) Update(ctx context.Context) (*models.Customer, error) {
	mr := models.CustomerRepository{Db: internal.Database(ctx)}
	model := cr.ToModel(false)
	err := mr.UpdateCustomer(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (cr *CustomerRequest) Get(ctx context.Context) ([]*models.Customer, error) {
	mr := models.CustomerRepository{Db: internal.Database(ctx)}
	return mr.GetAllCustomers()
}

func (cr *CustomerRequest) Detail(ctx context.Context, id string) (*models.Customer, error) {
	mr := models.CustomerRepository{Db: internal.Database(ctx)}
	return mr.GetCustomerByID(id)
}
