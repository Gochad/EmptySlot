package services

import (
	"fmt"

	"backend/internal/models"
	"backend/stripeGateway"
)

func makePaymentLink(reservation models.Reservation) (string, error) {
	price := int64(0)
	for _, merch := range reservation.Merchandises {
		price = price + merch.Price
	}
	link, err := stripeGateway.Pay(reservation.Merchandises[0].Name, "jakis opis", price)
	//link, err := stripeGateway.Pay(reservation.Merchandises[0].Name, reservation.Merchandises[0].Description, price)
	if err != nil {
		return "", fmt.Errorf("error while creating payment link: err: %v", err)
	}

	return link, nil
}
