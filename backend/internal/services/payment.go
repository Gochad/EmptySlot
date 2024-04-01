package services

import (
	"fmt"

	"backend/internal/models"
	"backend/stripegateway"
)

func makePaymentLink(reservation models.Reservation, redirectURL string) (string, error) {
	link, err := stripegateway.Pay(reservation.Name, reservation.Description, redirectURL, reservation.CalculatedPrice)
	if err != nil {
		return "", fmt.Errorf("error while creating payment link: err: %v", err)
	}

	return link, nil
}
