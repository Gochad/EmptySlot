package services

import (
	"context"
	"fmt"

	"backend/internal"
	"backend/internal/models"
)

type ReservationRequest struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       int64           `json:"price"`
	Confirmed   bool            `json:"confirmed"`
	StartTime   string          `json:"starttime"`
	EndTime     string          `json:"endtime"`
	IsReserved  bool            `json:"isreserved"`
	CustomerReq CustomerRequest `json:"customer"`
}

func (rr *ReservationRequest) ToModel(generateNewID bool) *models.Reservation {
	if generateNewID {
		rr.ID = generateUUID()
	}

	return &models.Reservation{
		ID:          rr.ID,
		Name:        rr.Name,
		Description: rr.Description,
		Price:       rr.Price,
		Confirmed:   rr.Confirmed,
		StartTime:   rr.StartTime,
		EndTime:     rr.EndTime,
		IsReserved:  rr.IsReserved,
		Customer:    *rr.CustomerReq.ToModel(generateNewID),
	}
}

func (rr *ReservationRequest) Create(ctx context.Context) (*models.Reservation, error) {
	mr := models.ReservationRepository{Db: internal.Database(ctx)}
	model := rr.ToModel(true)
	err := mr.CreateReservation(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (rr *ReservationRequest) Update(ctx context.Context) (*models.Reservation, error) {
	mr := models.ReservationRepository{Db: internal.Database(ctx)}
	model := rr.ToModel(false)
	err := mr.UpdateReservation(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (rr *ReservationRequest) Get(ctx context.Context) ([]*models.Reservation, error) {
	mr := models.ReservationRepository{Db: internal.Database(ctx)}
	return mr.GetAllReservations()
}

func (rr *ReservationRequest) Detail(ctx context.Context, id string) (*models.Reservation, error) {
	mr := models.ReservationRepository{Db: internal.Database(ctx)}
	return mr.GetReservationByID(id)
}

func (rr *ReservationRequest) DeleteOne(ctx context.Context, id string) error {
	mr := models.ReservationRepository{Db: internal.Database(ctx)}
	return mr.DeleteReservation(id)
}

func (rr *ReservationRequest) DeleteMany(ctx context.Context) error {
	mr := models.ReservationRepository{Db: internal.Database(ctx)}
	return mr.DeleteReservations()
}

func (rr *ReservationRequest) Pay(ctx context.Context, id string) (string, error) {
	mr := models.ReservationRepository{Db: internal.Database(ctx)}
	reservation, err := mr.GetReservationByID(id)
	if err != nil {
		return "", fmt.Errorf("getting reservation from db error: %v", err)
	}

	if reservation == nil {
		return "", fmt.Errorf("reservation from db is nil")
	}

	return makePaymentLink(*reservation)
}
