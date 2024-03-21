package services

import (
	"fmt"

	"backend/internal/models"
	"backend/stripegateway"
)

func makePaymentLink(reservation models.Reservation) (string, error) {
	//TODO: pass description from frontend
	link, err := stripegateway.Pay(reservation.Name, "reservation.Description", reservation.Price)
	if err != nil {
		return "", fmt.Errorf("error while creating payment link: err: %v", err)
	}

	return link, nil
}
