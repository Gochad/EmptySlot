package services

import (
	"context"
	"fmt"

	"backend/internal"
	"backend/internal/models"
)

type ReservationRequest struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	IsReserved     bool            `json:"isreserved"`
	MerchandiseIDs []string        `json:"merchandises"`
	CustomerReq    CustomerRequest `json:"customer"`
}

func (rr *ReservationRequest) ToModel(generateNewID bool) *models.Reservation {
	if generateNewID {
		rr.ID = generateUUID()
	}

	return &models.Reservation{
		ID:          rr.ID,
		Name:        rr.Name,
		Description: rr.Description,
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

func (rr *ReservationRequest) Pay(ctx context.Context, id, redirectURL string) (string, error) {
	repo := models.ReservationRepository{Db: internal.Database(ctx)}
	reservation, err := repo.GetReservationByID(id)
	if err != nil {
		return "", fmt.Errorf("getting reservation from db error: %v", err)
	}

	if reservation == nil {
		return "", fmt.Errorf("reservation from db is nil")
	}

	mrepo := models.MerchandiseRepository{Db: internal.Database(ctx)}

	//merchandises := make([]*models.Merchandise, 0, len(reservation.MerchandiseIDs))
	for _, merch := range reservation.MerchandiseIDs {
		merchandise, err := mrepo.GetMerchandiseByID(merch)
		if err != nil {
			return "", fmt.Errorf("error getting merchandise: %v", err)
		}
		//merchandises[i] = merchandise
		reservation.CalculatedPrice += merchandise.Price
	}

	return makePaymentLink(*reservation, redirectURL)
}
