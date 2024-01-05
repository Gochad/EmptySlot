package services

import (
	"context"

	"emptyslot/internal"
	"emptyslot/internal/models"
)

type ReservationRequest struct {
	ID             string             `json:"id"`
	MerchandiseReq MerchandiseRequest `json:"merchandise"`
	CustomerReq    CustomerRequest    `json:"customer"`
	Confirmed      bool               `json:"confirmed"`
	StartTime      string             `json:"starttime"`
	EndTime        string             `json:"endtime"`
	IsReserved     bool               `json:"isreserved"`
}

func (rr *ReservationRequest) ToModel(generateNewID bool) *models.Reservation {
	if generateNewID {
		rr.ID = generateUUID()
	}

	return &models.Reservation{
		ID:          rr.ID,
		Merchandise: *rr.MerchandiseReq.ToModel(generateNewID),
		Customer:    *rr.CustomerReq.ToModel(generateNewID),
		Confirmed:   rr.Confirmed,
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
