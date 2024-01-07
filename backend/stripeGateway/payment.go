package stripeGateway

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentlink"
)

// TODO: zroic jakas strukture ktora bedzie tworzyla price elementy w stripe api i potem będą tylko odczytywane z
// i przekazywane do stripe
func payForReservations(amount int64) (*stripe.PaymentLink, error) {
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			&stripe.PaymentLinkLineItemParams{
				Price:    stripe.String("price_1MoC3TLkdIwHu7ixcIbKelAC"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	return paymentlink.New(params)
}
