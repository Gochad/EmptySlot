package stripeGateway

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentlink"
)

type PaymentLink struct {
	URL    string
	Active bool
}

func generatePaymentLink(price *stripe.Price) (*PaymentLink, error) {
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			{
				Price:    stripe.String(price.ID),
				Quantity: stripe.Int64(1),
			},
		},
	}
	stripeLink, err := paymentlink.New(params)
	if err != nil {
		return nil, err
	}
	return &PaymentLink{URL: stripeLink.URL, Active: stripeLink.Active}, nil
}
