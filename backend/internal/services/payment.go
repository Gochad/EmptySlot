package services

import (
	"backend/internal/models"
	"backend/stripeGateway"
)

func makePaymentLink(reservation models.Reservation) (string, error) {
	link, err := stripeGateway.Pay(reservation.Merchandise.Name, reservation.Merchandise.Description, reservation.Merchandise.Price)
	if err != nil {
		return "", err
	}

	return link, nil
}
