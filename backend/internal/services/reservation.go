package services

import (
	"context"
	"fmt"

	"backend/internal"
	"backend/internal/models"
)

type ReservationRequest struct {
	ID              string               `json:"id"`
	MerchandisesReq []MerchandiseRequest `json:"merchandises"`
	CustomerReq     CustomerRequest      `json:"customer"`
	Confirmed       bool                 `json:"confirmed"`
	StartTime       string               `json:"starttime"`
	EndTime         string               `json:"endtime"`
	IsReserved      bool                 `json:"isreserved"`
}

func (rr *ReservationRequest) ToModel(generateNewID bool) *models.Reservation {
	if generateNewID {
		rr.ID = generateUUID()
	}

	merchandises := make([]models.Merchandise, len(rr.MerchandisesReq))
	for _, merch := range rr.MerchandisesReq {
		merchandises = append(merchandises, *merch.ToModel(generateNewID, ""))
	}

	return &models.Reservation{
		ID:           rr.ID,
		Merchandises: merchandises,
		Customer:     *rr.CustomerReq.ToModel(generateNewID),
		Confirmed:    rr.Confirmed,
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
