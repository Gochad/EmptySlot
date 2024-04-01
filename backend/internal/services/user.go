package services

import (
	"context"

	"backend/internal"
	"backend/internal/models"
)

type UserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`

	ReservationID string `json:"reservation"`
}

func (ur *UserRequest) ToModel(generateNewID bool) *models.User {
	if generateNewID {
		ur.ID = generateUUID()
	}
	return &models.User{
		ID:            ur.ID,
		Username:      ur.Username,
		Password:      ur.Password,
		Email:         ur.Email,
		Address:       ur.Address,
		Phone:         ur.Phone,
		ReservationID: ur.ReservationID,
	}
}

func (ur *UserRequest) Create(ctx context.Context) (*models.User, error) {
	mr := models.UserRepository{Db: internal.Database(ctx)}
	model := ur.ToModel(true)
	err := mr.CreateUser(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (ur *UserRequest) Update(ctx context.Context) (*models.User, error) {
	mr := models.UserRepository{Db: internal.Database(ctx)}
	model := ur.ToModel(false)
	err := mr.UpdateUser(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (ur *UserRequest) Get(ctx context.Context) ([]*models.User, error) {
	u := models.UserRepository{Db: internal.Database(ctx)}
	return u.GetAllUsers()
}

func (ur *UserRequest) Detail(ctx context.Context, email string) (*models.User, error) {
	u := models.UserRepository{Db: internal.Database(ctx)}
	return u.GetUserByEmail(email)
}
